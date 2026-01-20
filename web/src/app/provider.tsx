import {QueryClient, QueryClientProvider} from '@tanstack/react-query';
import * as React from "react";

import {Notifications} from '@/components/ui/notifications';
import {queryConfig} from '@/lib/react-query';

export function AppProvider({children}: { children: React.ReactNode }) {
    const [queryClient] = React.useState(
        () =>
            new QueryClient({
                defaultOptions: queryConfig,
            }),
    );
    return (
        <QueryClientProvider client={queryClient}>
            <Notifications/>
            <>{children}</>
        </QueryClientProvider>
    )
}