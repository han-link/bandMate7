import {usePerformances} from "../api/get-performances.ts";

import {PerformanceCard} from "./performance-card.tsx"

export function PerformanceList() {
    const performancesQuery = usePerformances({});

    if (performancesQuery.isLoading) {
        return (
            <div className="flex h-48 w-full items-center justify-center">
                <p>Loading</p> //ToDo: Replace with spinner
            </div>
        );
    }

    const performances = performancesQuery.data?.data;

    if (!performances) return null;

    return (
        <div className="flex gap-4 flex-wrap">
            {performances?.map((performance) => (
                <PerformanceCard
                    key={performance.id}
                    performance={performance}
                />
            ))}
        </div>
    )
}