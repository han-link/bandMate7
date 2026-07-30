import { useQuery } from "@tanstack/react-query";
import { Paragraph, XStack } from "tamagui";

import { getPerformancesOptions } from "@/client/@tanstack/react-query.gen";

import { PerformanceCard } from "./PerformanceCard";
import { Pagination } from "@/app/types/pagination";

interface PerformancesGridProps extends Pagination{}

export function PerformancesGrid(props: PerformancesGridProps) {
  const performancesQuery = useQuery(getPerformancesOptions({
    query: {
      desc: props.desc,
      orderBy: props.orderBy
    }
  }));

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
