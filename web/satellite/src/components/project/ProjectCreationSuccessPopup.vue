// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div v-if="isPopupShown" class="project-creation-success-popup-container">
        <div  class="project-creation-success-popup" id='successfulProjectCreationPopup'>
            <div class="project-creation-success-popup__info-panel-container">
                <svg width="300" height="233" viewBox="0 0 300 233" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M175.712 232.904C239.885 232.904 291.878 180.748 291.878 116.411C291.878 52.0744 239.885 0 175.712 0C111.539 0 59.5459 52.1561 59.5459 116.493C59.5459 180.83 111.539 232.904 175.712 232.904Z" fill="#2683FF"/>
                    <path d="M239.721 179.603H113.173C107.614 179.603 103.036 175.025 103.036 169.466V68.9143C103.036 63.3553 107.614 58.7773 113.173 58.7773H239.639C245.198 58.7773 249.776 63.3553 249.776 68.9143V169.466C249.776 175.025 245.28 179.603 239.721 179.603Z" fill="#499FFC"/>
                    <path d="M233.428 168.897H121.186C118.161 168.897 115.627 166.363 115.627 163.338V77.8285C115.627 74.8038 118.161 72.2695 121.186 72.2695H233.428C236.453 72.2695 238.987 74.8038 238.987 77.8285V163.338C238.987 166.363 236.453 168.897 233.428 168.897Z" fill="white"/>
                    <path d="M162.06 161.538C161.406 161.538 160.834 160.966 160.834 160.312V80.361C160.834 79.707 161.406 79.1348 162.06 79.1348C162.714 79.1348 163.286 79.707 163.286 80.361V160.312C163.286 160.966 162.714 161.538 162.06 161.538Z" fill="black"/>
                    <path d="M191.326 133.416L171.788 119.11L191.326 103.251C191.898 102.76 191.98 101.943 191.49 101.37C190.999 100.798 190.182 100.716 189.609 101.207L168.845 118.129C168.518 118.374 168.354 118.783 168.354 119.192C168.354 119.6 168.6 120.009 168.927 120.254L189.855 135.542C190.1 135.705 190.345 135.787 190.672 135.787C191.081 135.787 191.49 135.623 191.735 135.215C192.062 134.642 191.898 133.825 191.326 133.416Z" fill="#2683FF"/>
                    <path d="M232.772 118.127L211.844 101.123C211.272 100.632 210.454 100.714 209.964 101.286C209.473 101.859 209.555 102.676 210.127 103.167L229.666 119.026L210.127 133.332C209.555 133.741 209.392 134.558 209.882 135.212C210.127 135.539 210.536 135.785 210.945 135.785C211.19 135.785 211.517 135.703 211.762 135.539L232.69 120.252C233.017 120.007 233.181 119.598 233.263 119.189C233.263 118.781 233.099 118.372 232.772 118.127Z" fill="#2683FF"/>
                    <path d="M204.487 97.3635C203.751 97.2 203.097 97.6905 202.934 98.4263L195.249 139.464C195.086 140.2 195.576 140.854 196.312 141.018C196.476 141.018 196.639 141.018 196.803 141.018C197.293 140.936 197.784 140.527 197.865 139.955L205.55 98.9168C205.713 98.2628 205.223 97.527 204.487 97.3635Z" fill="#2683FF"/>
                    <path d="M159.852 69.403H118.242C116.607 69.403 115.217 68.0132 115.217 66.3782C115.217 64.7433 116.607 63.3535 118.242 63.3535H159.852C161.487 63.3535 162.877 64.7433 162.877 66.3782C162.877 68.0132 161.569 69.403 159.852 69.403Z" fill="white"/>
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M298.253 150.664C299.806 150.664 300.951 151.89 300.951 153.444C300.215 174.535 291.304 193.583 277.325 207.562C263.346 221.541 244.299 230.451 223.207 231.187C221.654 231.269 220.428 230.043 220.428 228.49V204.537C220.428 203.065 221.572 201.921 223.044 201.839C236.124 201.185 247.896 195.544 256.643 186.879C265.39 178.132 270.949 166.36 271.603 153.28C271.685 151.809 272.829 150.664 274.301 150.664H298.253Z" fill="#DEDEFC"/>
                    <path d="M221.655 225.058C226.892 225.058 231.138 220.812 231.138 215.575C231.138 210.337 226.892 206.092 221.655 206.092C216.418 206.092 212.172 210.337 212.172 215.575C212.172 220.812 216.418 225.058 221.655 225.058Z" fill="#499FFC"/>
                    <path d="M67.3937 30.902C74.3424 30.902 80.0649 25.016 80.0649 17.8221C80.0649 10.6281 74.3424 4.74219 67.3937 4.74219C60.445 4.74219 54.7225 10.6281 54.7225 17.8221C54.6408 25.0978 60.3633 30.902 67.3937 30.902Z" fill="#499FFC"/>
                    <path d="M67.3943 23.5445C70.5008 23.5445 73.1168 21.0103 73.1168 17.8221C73.1168 14.7156 70.5826 12.0996 67.3943 12.0996C64.2879 12.0996 61.6719 14.6338 61.6719 17.8221C61.6719 21.0103 64.2061 23.5445 67.3943 23.5445Z" fill="white"/>
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M138.843 36.3769C141.786 36.3769 144.238 33.9244 144.238 30.9814C144.238 28.0384 141.786 25.5859 138.843 25.5859C135.9 25.5859 133.447 28.0384 133.447 30.9814C133.366 33.9244 135.818 36.3769 138.843 36.3769Z" fill="#6CB9FF"/>
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M102.383 207.398C105.326 207.398 107.778 204.946 107.778 202.003C107.778 199.06 105.326 196.607 102.383 196.607C99.4398 196.607 96.9873 199.06 96.9873 202.003C96.9873 205.028 99.358 207.398 102.383 207.398Z" fill="white"/>
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M117.916 207.398C120.859 207.398 123.311 204.946 123.311 202.003C123.311 199.06 120.859 196.607 117.916 196.607C114.973 196.607 112.521 199.06 112.521 202.003C112.439 205.028 114.891 207.398 117.916 207.398Z" fill="white"/>
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M133.366 207.398C136.309 207.398 138.762 204.946 138.762 202.003C138.762 199.06 136.309 196.607 133.366 196.607C130.423 196.607 127.971 199.06 127.971 202.003C127.971 205.028 130.423 207.398 133.366 207.398Z" fill="#6CB9FF"/>
                    <path d="M124.781 110.036H140.068C141.622 110.036 142.93 108.81 142.93 107.338C142.93 105.867 141.622 104.641 140.068 104.641H124.781C123.228 104.641 121.92 105.867 121.92 107.338C121.92 108.81 123.228 110.036 124.781 110.036Z" fill="#0F002D"/>
                    <path d="M124.699 99.2448H151.513C153.066 99.2448 154.374 97.855 154.374 96.22C154.374 94.5851 153.066 93.1953 151.513 93.1953H124.781C123.228 93.1953 121.92 94.5851 121.92 96.22C121.92 97.855 123.146 99.2448 124.699 99.2448Z" fill="#0F002D"/>
                    <path d="M126.09 125.405H127.888C129.441 125.405 130.749 124.015 130.749 122.38C130.749 120.663 129.523 119.355 127.888 119.355H126.09C124.537 119.355 123.229 120.745 123.229 122.38C123.31 124.097 124.537 125.405 126.09 125.405Z" fill="#2683FF"/>
                    <path d="M136.881 125.405H138.679C140.232 125.405 141.54 124.015 141.54 122.38C141.54 120.663 140.314 119.355 138.679 119.355H136.881C135.328 119.355 134.02 120.745 134.02 122.38C134.101 124.097 135.328 125.405 136.881 125.405Z" fill="#2683FF"/>
                    <path d="M116.361 136.03C115.952 136.03 115.625 136.03 115.298 136.111C114.644 136.193 113.99 135.784 113.745 135.212C106.061 121.233 91.1824 111.75 74.015 111.75C58.4826 111.75 44.7487 119.516 36.6555 131.37C36.2468 131.942 35.511 132.187 34.857 132.106C32.8133 131.697 30.6878 131.452 28.5623 131.452C12.5395 131.37 -0.376953 144.123 -0.376953 159.982C-0.376953 159.982 -0.376953 159.982 -0.376953 160.064C-0.376953 160.963 0.358791 161.699 1.33978 161.699H138.761C139.66 161.699 140.477 160.963 140.477 160.064V159.982C140.477 146.739 129.686 136.03 116.361 136.03Z" fill="#DEDEFC"/>
                </svg>
            </div>
            <div class="project-creation-success-popup__form-container">
                <h2 class="project-creation-success-popup__form-container__main-label-text">Congrats!</h2>
                <p>You just created your project. Next, we recommend you create your first API Key for this project. API Keys allow developers to manage their projects and build applications on top of the Storj network through our
                    <a href="https://github.com/storj/storj/wiki/Uplink-CLI" target="_blank">Uplink CLI.</a>
                </p>
                <div class="project-creation-success-popup__form-container__button-container">
                    <Button label="I will do it later" width="214px" height="50px" :onPress="onCloseClick" isWhite="true" />
                    <Button label="Create first API Key" width="214px" height="50px" :onPress="onCreateAPIKeyClick" />
                </div>
            </div>
            <div class="project-creation-success-popup__close-cross-container" @click="onCloseClick">
                <svg width="16" height="16" viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg" >
                    <path d="M15.7071 1.70711C16.0976 1.31658 16.0976 0.683417 15.7071 0.292893C15.3166 -0.0976311 14.6834 -0.0976311 14.2929 0.292893L15.7071 1.70711ZM0.292893 14.2929C-0.0976311 14.6834 -0.0976311 15.3166 0.292893 15.7071C0.683417 16.0976 1.31658 16.0976 1.70711 15.7071L0.292893 14.2929ZM1.70711 0.292893C1.31658 -0.0976311 0.683417 -0.0976311 0.292893 0.292893C-0.0976311 0.683417 -0.0976311 1.31658 0.292893 1.70711L1.70711 0.292893ZM14.2929 15.7071C14.6834 16.0976 15.3166 16.0976 15.7071 15.7071C16.0976 15.3166 16.0976 14.6834 15.7071 14.2929L14.2929 15.7071ZM14.2929 0.292893L0.292893 14.2929L1.70711 15.7071L15.7071 1.70711L14.2929 0.292893ZM0.292893 1.70711L14.2929 15.7071L15.7071 14.2929L1.70711 0.292893L0.292893 1.70711Z" fill="#384B65"/>
                </svg>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Vue } from 'vue-property-decorator';
    import Button from '@/components/common/Button.vue';
    import { APP_STATE_ACTIONS } from '@/utils/constants/actionNames';
    import { RouteConfig } from '@/router';

    @Component({
        components: {
            Button,
        }
    })
    export default class ProjectCreationSuccessPopup extends Vue {
        private onCloseClick(): void {
            this.$store.dispatch(APP_STATE_ACTIONS.TOGGLE_SUCCESSFUL_PROJECT_CREATION_POPUP);
        }

        public onCreateAPIKeyClick(): void {
            this.$router.push(RouteConfig.ApiKeys.path);
            this.onCloseClick();
        }

        public get isPopupShown(): boolean {
            return this.$store.state.appStateModule.appState.isSuccessfulProjectCreationPopupShown;
        }
    }
</script>

<style scoped lang="scss">
    p {
        font-family: 'font_medium';
        font-size: 16px;
        line-height: 21px;
        color: #354049;
        padding: 27px 0 0 0;
        margin: 0;
    }
    
    a {
        font-family: 'font_bold';
        color: #2683ff;
    }
    
    .project-creation-success-popup-container {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        background-color: rgba(134, 134, 148, 0.4);
        z-index: 1000;
        display: flex;
        justify-content: center;
        align-items: center;
    }
    
    .project-creation-success-popup {
        width: 100%;
        max-width: 845px;
        background-color: #FFFFFF;
        border-radius: 6px;
        display: flex;
        flex-direction: row;
        align-items: flex-start;
        position: relative;
        justify-content: center;
        padding: 80px 100px 80px 50px;
        
        &__info-panel-container {
            display: flex;
            flex-direction: column;
            justify-content: flex-start;
            align-items: center;
            margin-right: 100px;
            margin-top: 20px;
        }
        
        &__form-container {
            width: 100%;
            max-width: 440px;
            margin-top: 10px;
            
            &__main-label-text {
                font-family: 'font_bold';
                font-size: 32px;
                line-height: 39px;
                color: #384B65;
                margin: 0;
            }
            
            &__button-container {
                width: 100%;
                display: flex;
                flex-direction: row;
                justify-content: space-between;
                align-items: center;
                margin-top: 40px;
            }
        }
        
        &__close-cross-container {
            display: flex;
            justify-content: center;
            align-items: center;
            position: absolute;
            right: 30px;
            top: 40px;
            height: 24px;
            width: 24px;
            cursor: pointer;
            
            &:hover svg path {
                fill: #2683FF;
            }
        }
    }
    
    @media screen and (max-width: 720px) {
        .project-creation-success-popup {
        
            &__info-panel-container {
                display: none;
            }
            
            &__form-container {
                
                &__button-container {
                    width: 100%;
                }
            }
        }
    }
</style>
