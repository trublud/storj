// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <!-- To close popup we need to use method onCloseClick -->
    <div class="account-dropdown-choice-container" id="accountDropdown">
        <div class="account-dropdown-overflow-container">
            <!-- TODO: add selection logic onclick -->
            <div class="account-dropdown-item-container settings" v-on:click="onAccountSettingsClick" >
                <div class="account-dropdown-item-container__image-container">
                    <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M20 8.72549C20 9.60784 20 10.4902 20 11.2745C19.9024 11.5686 19.7073 11.6667 19.4146 11.7647C18.7317 11.9608 18.1463 12.0588 17.4634 12.2549C17.3659 12.2549 17.2683 12.3529 17.2683 12.451C17.1707 12.7451 16.9756 13.0392 16.878 13.4314C16.878 13.5294 16.878 13.6275 16.878 13.7255C17.1707 14.2157 17.4634 14.7059 17.7561 15.2941C17.9512 15.6863 17.9512 15.8824 17.6585 16.1765C17.1707 16.6667 16.6829 17.1569 16.1951 17.6471C15.9024 17.9412 15.6098 17.9412 15.3171 17.7451C14.8293 17.451 14.3415 17.1569 13.8537 16.8627C13.7561 16.7647 13.6585 16.7647 13.561 16.8627C13.2683 17.0588 12.878 17.1569 12.5854 17.2549C12.4878 17.2549 12.3902 17.3529 12.3902 17.451C12.1951 18.1373 12.0976 18.7255 11.9024 19.4118C11.8049 19.7059 11.7073 19.902 11.4146 20C10.5366 20 9.65854 20 8.87805 20C8.58537 19.902 8.48781 19.7059 8.39024 19.4118C8.19512 18.7255 8.09756 18.1373 7.90244 17.451C7.90244 17.3529 7.80488 17.2549 7.70732 17.2549C7.41463 17.1569 7.02439 16.9608 6.73171 16.8627C6.63415 16.7647 6.53659 16.8627 6.43902 16.8627C5.95122 17.1569 5.46341 17.451 4.97561 17.7451C4.58537 17.9412 4.39024 17.9412 4.09756 17.6471C3.60976 17.1569 3.12195 16.6667 2.63415 16.1765C2.34146 15.8824 2.34146 15.5882 2.53659 15.2941C2.82927 14.8039 3.12195 14.3137 3.41463 13.8235C3.5122 13.7255 3.5122 13.6275 3.41463 13.5294C3.21951 13.2353 3.12195 12.8431 3.02439 12.549C3.02439 12.451 2.92683 12.3529 2.82927 12.3529C2.14634 12.1569 1.46341 12.0588 0.780488 11.8627C0.195122 11.7647 0 11.5686 0 11.1765C0 10.4902 0 9.70588 0 9.01961C0 8.62745 0.195122 8.43137 0.585366 8.33333C1.26829 8.13725 1.95122 7.94118 2.53659 7.7451C2.63415 7.7451 2.73171 7.64706 2.73171 7.54902C2.92683 7.35294 3.02439 7.05882 3.12195 6.76471C3.12195 6.66667 3.12195 6.56863 3.12195 6.47059C2.82927 5.88235 2.43902 5.39216 2.14634 4.80392C1.95122 4.41176 1.95122 4.21569 2.2439 3.92157C2.73171 3.43137 3.21951 2.94118 3.70732 2.45098C4 2.15686 4.29268 2.15686 4.58537 2.35294C5.17073 2.64706 5.65854 3.03922 6.2439 3.33333C6.34146 3.33333 6.43902 3.33333 6.53659 3.33333C6.82927 3.13725 7.21951 2.94118 7.5122 2.84314C7.60976 2.84314 7.70732 2.7451 7.70732 2.64706C7.90244 1.96078 8 1.27451 8.19512 0.588235C8.39024 0.196078 8.58537 0 8.97561 0C9.65854 0 10.439 0 11.122 0C11.5122 0 11.7073 0.196078 11.8049 0.588235C12 1.27451 12.0976 1.96078 12.2927 2.64706C12.2927 2.7451 12.3902 2.84314 12.4878 2.84314C12.7805 2.94118 13.1707 3.13725 13.4634 3.23529C13.561 3.33333 13.6585 3.23529 13.7561 3.23529C14.2439 2.94118 14.7317 2.64706 15.2195 2.35294C15.6098 2.15686 15.8049 2.15686 16.1951 2.45098C16.6829 2.94118 17.1707 3.43137 17.6585 3.92157C17.9512 4.21569 18.0488 4.5098 17.7561 4.90196C17.4634 5.39216 17.1707 5.88235 16.878 6.47059C16.878 6.56863 16.878 6.66667 16.878 6.76471C16.9756 7.05882 17.1707 7.45098 17.2683 7.7451C17.2683 7.84314 17.3659 7.94118 17.4634 7.94118C18.1463 8.13726 18.7317 8.23529 19.4146 8.43137C19.7073 8.33333 19.9024 8.52941 20 8.72549ZM10.0488 5.98039C7.80488 5.98039 5.95122 7.84314 5.95122 10.098C5.95122 12.3529 7.80488 14.2157 10.0488 14.2157C12.2927 14.2157 14.1463 12.3529 14.1463 10.098C14.1463 7.7451 12.2927 5.98039 10.0488 5.98039Z" fill="#354049"/>
                    </svg>
                </div>
                <h2>Account Settings</h2>
            </div>
            <div class="account-dropdown-item-container logout" v-on:click="onLogoutClick">
                <div class="account-dropdown-item-container__image-container">
                    <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path fill-rule="evenodd" clip-rule="evenodd" d="M9.53516 1C4.28906 1 0 5.25293 0 10.5C0 15.7471 4.28906 20 9.53516 20C14.3887 20 18.4229 16.3604 19 11.6611H10.2783L11.8682 13.2441C12.1562 13.5322 12.248 14.0156 12.1436 14.4277C12.0947 14.6211 12.0029 14.7988 11.8682 14.9336C11.4443 15.3555 10.5967 15.3555 10.1729 14.9336L5.72266 10.5L10.1729 6.06641C10.5967 5.64453 11.4443 5.64453 11.8682 6.06641C12.291 6.48926 12.291 7.33301 11.8682 7.75586L10.2783 9.33887H19C18.4229 4.63965 14.3887 1 9.53516 1Z" fill="#354049"/>
                    </svg>
                </div>
                <h2>Log Out</h2>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator';
    import { AuthToken } from '@/utils/authToken';
    import { RouteConfig } from '@/router';
    import {
        APP_STATE_ACTIONS,
        PM_ACTIONS,
        API_KEYS_ACTIONS,
        NOTIFICATION_ACTIONS,
    } from '@/utils/constants/actionNames';
    import { USER_ACTIONS } from '@/store/modules/users';
    import { BUCKET_ACTIONS } from '@/store/modules/buckets';
    import { PROJECTS_ACTIONS } from '@/store/modules/projects';

    @Component
    export default class AccountDropdown extends Vue {
        public onCloseClick(): void {
            this.$store.dispatch(APP_STATE_ACTIONS.TOGGLE_ACCOUNT);
        }

        public onAccountSettingsClick(): void {
            this.$router.push(RouteConfig.Account.with(RouteConfig.Profile).path);
            this.$store.dispatch(APP_STATE_ACTIONS.TOGGLE_ACCOUNT);
        }

        public onLogoutClick(): void {
            AuthToken.remove();

            this.$router.push(RouteConfig.Login.path);
            this.$store.dispatch(PM_ACTIONS.CLEAR);
            this.$store.dispatch(PROJECTS_ACTIONS.CLEAR);
            this.$store.dispatch(USER_ACTIONS.CLEAR);
            this.$store.dispatch(API_KEYS_ACTIONS.CLEAR);
            this.$store.dispatch(NOTIFICATION_ACTIONS.CLEAR);
            this.$store.dispatch(BUCKET_ACTIONS.CLEAR);
        }
    }
</script>

<style scoped lang="scss">
    .account-dropdown-choice-container {
        position: absolute;
        top: 9vh;
        right: 0;
        border-radius: 4px;
        padding: 10px 0px 10px 0px;
        box-shadow: 0px 4px rgba(231, 232, 238, 0.6);
        background-color: #FFFFFF;
        z-index: 1120;
    }
    .account-dropdown-overflow-container {
        position: relative;
        width: 210px;
        height: auto;
        background-color: #FFFFFF;
    }
    .account-dropdown-item-container {
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: flex-start;
        padding-left: 20px;
        padding-right: 20px;

        h2{
            font-family: 'font_regular';
            margin-left: 20px;
            font-size: 14px;
            line-height: 20px;
            color: #354049;
        }

        &:hover {
            background-color: #F2F2F6;

            path {
                fill: #2683FF !important;
            }
        }

        &__image-container {
            width: 20px;

            svg {
                object-fit: cover;
            }
        }
    }

    a {
        text-decoration: none;
        outline: none;
    }
</style>
