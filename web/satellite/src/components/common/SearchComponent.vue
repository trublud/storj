// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <input
        ref="input"
        @mouseenter="onMouseEnter"
        @mouseleave="onMouseLeave"
        @input="processSearchQuery"
        v-model="searchQuery"
        :placeholder="`Search ${placeHolder}`"
        :style="style"
        type="text"
        autocomplete="off">
</template>

<script lang="ts">
    import { Component, Prop, Vue } from 'vue-property-decorator';

    declare type searchCallback = (search: string) => Promise<void>;
    declare interface SearchStyle {
        width: string;
    }

    @Component
    export default class SearchComponent extends Vue {
        @Prop({default: ''})
        private readonly placeHolder: string;
        @Prop({default: () => { return ''; }})
        private readonly search: searchCallback;

        private inputWidth: string = '56px';
        private searchQuery: string = '';

        public $refs!: {
            input: HTMLElement;
        };

        public get style(): SearchStyle {
            return { width: this.inputWidth };
        }

        public get searchString(): string {
            return this.searchQuery;
        }

        public onMouseEnter(): void {
            this.inputWidth = '602px';

            this.$refs.input.focus();
        }

        public onMouseLeave(): void {
            if (!this.searchString) {
                this.inputWidth = '56px';
            }

            this.$refs.input.blur();
        }

        public clearSearch() {
            this.searchQuery = '';
            this.processSearchQuery();
            this.inputWidth = '56px';
        }

        private async processSearchQuery() {
            await this.search(this.searchString);
        }
    }
</script>

<style scoped lang="scss">
    input {
        position: absolute;
        right: 0;
        bottom: 0;
        padding: 0 38px 0 18px;
        border: 1px solid #F2F2F2;
        box-sizing: border-box;
        box-shadow: 0 4px 4px rgba(231, 232, 238, 0.6);
        outline: none;
        border-radius: 36px;
        height: 56px;
        font-family: 'font_regular';
        font-size: 16px;
        transition: all 0.4s ease-in-out;
        background-image: url('../../../static/images/team/searchIcon.svg');
        background-repeat: no-repeat;
        background-size: 22px 22px;
        background-position: top 16px right 16px;
    }

    ::-webkit-input-placeholder {
        color: #AFB7C1;
    }
</style>
