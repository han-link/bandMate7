import { useQuery } from "@tanstack/react-query";
import { Paragraph, XStack } from "tamagui";

import { getPerformancesOptions } from "@/client/@tanstack/react-query.gen";

import { PerformanceCard } from "./PerformanceCard";

export function PerformancesGrid() {
  const performancesQuery = useQuery(getPerformancesOptions());

  if (performancesQuery.isLoading) return <Paragraph>Loading …</Paragraph>;

  if (performancesQuery.data?.length === 0)
    return <Paragraph>No Performances</Paragraph>;

  if (performancesQuery.isError || performancesQuery.data === undefined)
    return <Paragraph>Query Error</Paragraph>;

  return (
    <XStack flexWrap="wrap" gap="$4" justify="flex-start">
      {performancesQuery.data.map((performance, index) => (
        <PerformanceCard key={index} performance={performance} />
      ))}
    </XStack>
  );
}
