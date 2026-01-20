import {PerformanceList} from "@/features/performances/components/performance-list.tsx";
import {ContentLayout} from "@/components/layouts/content-layout.tsx";

export function DashboardRoute() {
    return (
        <ContentLayout title={"Dashboard"}>
            <PerformanceList/>
        </ContentLayout>
    )
}