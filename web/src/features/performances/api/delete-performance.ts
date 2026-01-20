import {useMutation, useQueryClient} from '@tanstack/react-query';

import {api} from '@/lib/api-client';
import type {MutationConfig} from '@/lib/react-query';

import {getPerformancesQueryOptions} from './get-performances';

type DeletePerformanceParams = {
    performanceId: string;
}

function deletePerformance({performanceId}: DeletePerformanceParams) {
    return api.delete(`/performances/${performanceId}`);
}

type UseDeletePerformanceOptions = {
    mutationConfig?: MutationConfig<typeof deletePerformance>;
};

export function useDeletePerformance({mutationConfig}: UseDeletePerformanceOptions = {}) {
    const queryClient = useQueryClient();

    const {onSuccess, ...restConfig} = mutationConfig || {};

    return useMutation({
        onSuccess: (...args) => {
            queryClient.invalidateQueries({
                queryKey: getPerformancesQueryOptions().queryKey,
            });
            onSuccess?.(...args);
        }, ...restConfig, mutationFn: deletePerformance,
    });
};