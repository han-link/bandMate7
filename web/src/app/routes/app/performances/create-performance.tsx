import {ContentLayout} from "@/components/layouts/content-layout.tsx";
import {CreatePerformance} from "@/features/performances/components/create-performance.tsx";
import {paths} from "@/config/paths.ts";

export function CreatePerformanceRoute() {
    return (
        <ContentLayout title={"Create Performance"}>
            <CreatePerformance onSuccessTo={paths.app.dashboard.href()}/>
        </ContentLayout>
    );
}