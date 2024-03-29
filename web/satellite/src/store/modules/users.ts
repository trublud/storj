// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import { UpdatedUser, User, UsersApi } from '@/types/users';
import { StoreModule } from '@/store';

export const USER_ACTIONS = {
    UPDATE: 'updateUser',
    GET: 'getUser',
    CLEAR: 'clearUser',
};

export const USER_MUTATIONS = {
    SET_USER: 'setUser',
    UPDATE_USER: 'updateUser',
    CLEAR: 'clearUser',
};

const {
    GET,
    UPDATE,
} = USER_ACTIONS;

const {
    SET_USER,
    UPDATE_USER,
    CLEAR,
} = USER_MUTATIONS;

/**
 * creates users module with all dependencies
 *
 * @param api - users api
 */
export function makeUsersModule(api: UsersApi): StoreModule<User> {
    return {
        state: new User(),

        mutations: {
            [SET_USER](state: User, user: User): void {
                state.id = user.id;
                state.email = user.email;
                state.shortName = user.shortName;
                state.fullName = user.fullName;
                state.partnerId = user.partnerId;
            },

            [CLEAR](state: User): void {
                state.id = '';
                state.email = '';
                state.shortName = '';
                state.fullName = '';
                state.partnerId = '';
            },

            [UPDATE_USER](state: User, user: UpdatedUser): void {
                state.fullName = user.fullName;
                state.shortName = user.shortName;
            },
        },

        actions: {
            [UPDATE]: async function ({commit}: any, userInfo: UpdatedUser): Promise<void> {
                await api.update(userInfo);

                commit(UPDATE_USER, userInfo);
            },
            [GET]: async function ({commit}: any): Promise<User> {
                let user = await api.get();

                commit(SET_USER, user);

                return user;
            },
            [CLEAR]: function({commit}: any) {
                commit(CLEAR);
            },
        },

        getters: {
            user: (state: User): User => state,
            userName: (state: User): string => state.getFullName(),
        },
    };
}
