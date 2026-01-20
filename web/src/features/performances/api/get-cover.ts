import { queryOptions, useQuery } from "@tanstack/react-query";

import { api } from "@/lib/api-client";
import type { QueryConfig } from "@/lib/react-query";

// ToDo: Fix ts error
async function getCoverBlob(resourceId: string): Promise<Blob> {
    const blob = await api.get(`/resources/${resourceId}`, { responseType: "blob" });
    return blob as Blob;
}


export function getCoverQueryOptions(resourceId?: string) {
    return queryOptions({
        queryKey: ["cover", resourceId],
        queryFn: () => getCoverBlob(resourceId!),
        enabled: !!resourceId,
        staleTime: 5 * 60 * 1000,
        gcTime: 30 * 60 * 1000,
    });
}

type UseCoverOptions = {
    resourceId?: string;
    queryConfig?: QueryConfig<typeof getCoverQueryOptions>;
};

export const useCover = ({ resourceId, queryConfig }: UseCoverOptions) => {
    return useQuery({
        ...getCoverQueryOptions(resourceId),
        ...queryConfig,
    });
};
