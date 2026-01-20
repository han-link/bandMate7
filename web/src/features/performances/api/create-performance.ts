import {useMutation, useQueryClient} from '@tanstack/react-query';
import {z} from 'zod';

import {api} from '@/lib/api-client';
import type {MutationConfig} from '@/lib/react-query';
import type {Performance} from '@/types/api';
import {getPerformancesQueryOptions} from "@/features/performances/api/get-performances.ts";

export const createPerformanceInputSchema = z.object({
    name: z.string().min(1, 'Required'),
    bpm: z.number().int().gt(0).nullish(),
    cover: z.custom<File | null>().nullish(),
    file: z.custom<File | null>().nullish(),
});

function createPerformance({data}: { data: FormData }): Promise<Performance> {
    return api.post(`/performances`, data);
}

type UseCreateDiscussionOptions = {
    mutationConfig?: MutationConfig<typeof createPerformance>;
};

export function useCreatePerformance({mutationConfig}: UseCreateDiscussionOptions = {}) {
    const queryClient = useQueryClient();

    const {onSuccess, ...restConfig} = mutationConfig || {};

    return useMutation({
        onSuccess: (...args) => {
            queryClient.invalidateQueries({
                queryKey: getPerformancesQueryOptions().queryKey,
            });
            onSuccess?.(...args);
        },
        ...restConfig,
        mutationFn: createPerformance,
    });
};