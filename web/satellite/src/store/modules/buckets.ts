// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import { Bucket, BucketCursor, BucketPage, BucketsApi } from '@/types/buckets';
import { StoreModule } from '@/store';

export const BUCKET_ACTIONS = {
    FETCH: 'setBuckets',
    SET_SEARCH: 'setBucketSearch',
    CLEAR: 'clearBuckets'
};

export const BUCKET_MUTATIONS = {
    SET: 'setBuckets',
    SET_SEARCH: 'setBucketSearch',
    SET_PAGE: 'setBucketPage',
    CLEAR: 'clearBuckets',
};

const {
    FETCH
} = BUCKET_ACTIONS;
const {
    SET,
    SET_PAGE,
    SET_SEARCH,
    CLEAR,
} = BUCKET_MUTATIONS;
const bucketPageLimit = 8;
const firstPage = 1;

class BucketsState {
    public cursor: BucketCursor = { limit: bucketPageLimit, search: '', page: firstPage };
    public page: BucketPage = { buckets: new Array<Bucket>(), currentPage: 1, pageCount: 1, offset: 0, limit: bucketPageLimit, search: '', totalCount: 0 };
}

/**
 * creates buckets module with all dependencies
 *
 * @param api - buckets api
 */
export function makeBucketsModule(api: BucketsApi): StoreModule<BucketsState> {
    return {
        state: new BucketsState(),

        mutations: {
            [SET](state: BucketsState, page: BucketPage) {
                state.page = page;
            },
            [SET_PAGE](state: BucketsState, page: number) {
                state.cursor.page = page;
            },
            [SET_SEARCH](state: BucketsState, search: string) {
                state.cursor.search = search;
            },
            [CLEAR](state: BucketsState) {
                state.cursor = new BucketCursor('', bucketPageLimit, firstPage);
                state.page = new BucketPage([], '', bucketPageLimit, 0, 1, 1, 0);
            }
        },
        actions: {
            [FETCH]: async function({commit, rootGetters, state}: any, page: number): Promise<BucketPage> {
                const projectID = rootGetters.selectedProject.id;
                const before = new Date();
                state.cursor.page = page;

                commit(SET_PAGE, page);

                let result = await api.get(projectID, before, state.cursor);

                commit(SET, result);

                return result;
            },
            [BUCKET_ACTIONS.SET_SEARCH]: function({commit}, search: string) {
                commit(SET_SEARCH, search);
            },
            [BUCKET_ACTIONS.CLEAR]: function({commit}) {
                commit(CLEAR);
            }
        },
        getters: {
            page: (state: BucketsState): BucketPage => state.page,
            cursor: (state: BucketsState): BucketCursor => state.cursor,
        }
    };
}
