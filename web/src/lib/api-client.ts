import Axios from 'axios';

import { useNotifications } from '@/components/ui/notifications';
import { env } from '@/config/env';

export const api = Axios.create({
    baseURL: env.API_URL,
});
api.interceptors.response.use(
    (response) => {
        return response.data;
    },
    (error) => {
        const message = error.response?.data?.message || error.message;
        useNotifications.getState().addNotification({
            type: 'error',
            title: 'Error',
            message,
        });

        return Promise.reject(error);
    },
);