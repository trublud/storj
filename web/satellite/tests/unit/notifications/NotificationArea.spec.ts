// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

import { shallowMount, mount } from '@vue/test-utils';
import NotificationArea from '@/components/notifications/NotificationArea.vue';
import { NOTIFICATION_TYPES } from '@/utils/constants/notification';
import { DelayedNotification } from '@/types/DelayedNotification';

describe('NotificationArea.vue', () => {

    it('renders correctly', () => {
        const wrapper = shallowMount(NotificationArea, {
            computed: {
                notifications: () => []
            }
        });

        expect(wrapper).toMatchSnapshot();
    });

    it('renders correctly with notification', () => {
        const testMessage = 'testMessage';
        const notifications = [new DelayedNotification(
            jest.fn(),
            NOTIFICATION_TYPES.SUCCESS,
            testMessage
        ), new DelayedNotification(
            jest.fn(),
            NOTIFICATION_TYPES.ERROR,
            testMessage
        ), new DelayedNotification(
            jest.fn(),
            NOTIFICATION_TYPES.WARNING,
            testMessage
        ), new DelayedNotification(
            jest.fn(),
            NOTIFICATION_TYPES.NOTIFICATION,
            testMessage
        )];

        const wrapper = mount(NotificationArea, {
            computed: {
                notifications: () => notifications
            }
        });

        expect(wrapper).toMatchSnapshot();
    });
});
