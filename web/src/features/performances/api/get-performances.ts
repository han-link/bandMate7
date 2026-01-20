import {queryOptions, useQuery} from '@tanstack/react-query';

import {api} from '@/lib/api-client';
import type {QueryConfig} from '@/lib/react-query';
import type {Performance} from '@/types/api';

function getPerformances(): Promise<{ data: Performance[] }> {
    return api.get(`/performances`);
}

export function getPerformancesQueryOptions() {
    return queryOptions({
        queryKey: ['performances'],
        queryFn: () => getPerformances(),
    });
}

type UsePerformancesOptions = {
    queryConfig?: QueryConfig<typeof getPerformancesQueryOptions>;
};

export const usePerformances = ({queryConfig}: UsePerformancesOptions) => {
    return useQuery({
        ...getPerformancesQueryOptions(),
        ...queryConfig,
    });
};