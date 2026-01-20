import {createBrowserRouter, Navigate} from "react-router";
import {RouterProvider} from "react-router/dom";

import {paths} from "@/config/paths";
import {AppRoot} from "./routes/app/root";
import {NotFound} from "@/app/routes/not-found";

import {RouteErrorBoundary} from "@/app/routes/route-error-boundary";

export const router = createBrowserRouter([
    {
        path: paths.app.root.path,
        Component: AppRoot,
        ErrorBoundary: RouteErrorBoundary,
        HydrateFallback: () => <div>Loading…</div>,
        children: [
            {
                index: true,
                element: <Navigate to={paths.app.dashboard.path} replace/>,
            },
            {
                path: paths.app.dashboard.path,
                lazy: async () => {
                    const mod = await import("@/app/routes/app/dashboard");
                    return {Component: mod.DashboardRoute};
                },
            },
            {
                path: paths.app.performances.create.path,
                lazy: async () => {
                    const mod = await import("@/app/routes/app/performances/create-performance");
                    return {Component: mod.CreatePerformanceRoute};
                },
            },
            {
                path: paths.app.profile.path,
                lazy: async () => {
                    const mod = await import("@/app/routes/app/profile");
                    return {Component: mod.ProfileRoute};
                }
            }
        ],
    },

    {path: "*", Component: NotFound},
]);

export function AppRouter() {
    return <RouterProvider router={router}/>;
}
