// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div class="team-header-container">
	    <h1>Project Members</h1>
	    <div class="team-header-container__wrapper">
            <HeaderComponent ref="headerComponent" placeHolder="Team Members" :search="processSearchQuery">
                <div class="header-default-state" v-if="headerState === 0">
                    <span>The only project role currently available is Admin, which gives <b>full access</b> to the project.</span>
                    <Button class="button" label="+Add" width="122px" height="48px" :onPress="onAddUsersClick"/>
                </div>
                <div class="header-selected-members" v-if="headerState === 1 && !isDeleteClicked">
                    <Button class="button deletion" label="Delete" width="122px" height="48px" :onPress="onFirstDeleteClick"/>
                    <Button class="button" label="Cancel" width="122px" height="48px" isWhite="true" :onPress="onClearSelection"/>
                </div>
                <div class="header-after-delete-click" v-if="headerState === 1 && isDeleteClicked">
                    <span>Are you sure you want to delete {{selectedProjectMembers}} {{userCountTitle}}</span>
                    <div class="header-after-delete-click__button-area">
                        <Button class="button deletion" label="Delete" width="122px" height="48px" :onPress="onDelete"/>
                        <Button class="button" label="Cancel" width="122px" height="48px" isWhite="true" :onPress="onClearSelection"/>
                    </div>
                </div>
            </HeaderComponent>
	    </div>
    </div>
</template>

<script lang="ts">
    import { Component, Prop, Vue } from 'vue-property-decorator';

    import { APP_STATE_ACTIONS, NOTIFICATION_ACTIONS, PM_ACTIONS } from '@/utils/constants/actionNames';
    import Button from '@/components/common/Button.vue';
    import HeaderComponent from '@/components/common/HeaderComponent.vue';
    import { ProjectMember } from '@/types/projectMembers';
    import { RequestResponse } from '@/types/response';

    declare interface ClearSearch {
        clearSearch: () => void;
    }

    @Component({
        components: {
            Button,
            HeaderComponent,
        }
    })
    export default class HeaderArea extends Vue {
        @Prop({default: 0})
        private readonly headerState: number;
        @Prop({default: 0})
        private readonly selectedProjectMembers: number;

        private FIRST_PAGE = 1;

        private isDeleteClicked: boolean = false;

        public $refs!: {
            headerComponent: HeaderComponent & ClearSearch
        };

        public get userCountTitle(): string {
            if (this.selectedProjectMembers === 1) {
                return 'user';
            }

            return 'users';
        }

        public onAddUsersClick(): void {
            this.$store.dispatch(APP_STATE_ACTIONS.TOGGLE_TEAM_MEMBERS);
        }

        public onFirstDeleteClick(): void {
            this.isDeleteClicked = true;
        }

        public onClearSelection(): void {
            this.$store.dispatch(PM_ACTIONS.CLEAR_SELECTION);
            this.isDeleteClicked = false;

            this.$refs.headerComponent.clearSearch();
        }

        public async onDelete(): Promise<void> {
            const projectMemberEmails = this.$store.getters.selectedProjectMembers.map((member: ProjectMember) => {
                return member.user.email;
            });

            try {
                await this.$store.dispatch(PM_ACTIONS.DELETE, projectMemberEmails);
            } catch (err) {
                this.$store.dispatch(NOTIFICATION_ACTIONS.ERROR, `Error while deleting users from projectMembers. ${err.message}`);

                return;
            }

            this.$store.dispatch(NOTIFICATION_ACTIONS.SUCCESS, 'Members was successfully removed from project');
            this.isDeleteClicked = false;

            this.$refs.headerComponent.clearSearch();
        }

        public async processSearchQuery(search: string): Promise<void> {
            this.$store.dispatch(PM_ACTIONS.SET_SEARCH_QUERY, search);
            try {
                await this.$store.dispatch(PM_ACTIONS.FETCH, this.FIRST_PAGE);
            } catch (err) {
                this.$store.dispatch(NOTIFICATION_ACTIONS.ERROR, `Unable to fetch project members. ${err.message}`);
            }
        }
    }
</script>

<style scoped lang="scss">
	h1 {
		font-family: 'font_bold';
		font-size: 32px;
		line-height: 39px;
		margin: 0;
	}

    .header-default-state,
    .header-after-delete-click {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        height: 85px;

        &__button-area {
            display: flex;

            .deletion {
                margin-top: 2px;
            }
        }
    }

    .header-selected-members {
        display: flex;
        align-items: flex-end;
        height: 85px;
        justify-content: center;
    }

    .button {
        margin-right: 12px;
    }

    span {
        font-family: 'font_regular';
        font-size: 14px;
        line-height: 28px;
    }

    .team-header-container__wrapper {
        margin-bottom: 4px;
        display: flex;
        align-items: center;
        justify-content: flex-start;
    }

    .container.deletion {
        background-color: #FF4F4D;

        &.label {
            color: #FFFFFF;
        }

        &:hover {
            background-color: #DE3E3D;
            box-shadow: none;
        }
    }
</style>
