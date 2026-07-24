import { useLocalSearchParams } from "expo-router";
import { View, Text } from "tamagui";

import PdfView from "@/components/PdfView";
import { useQuery } from "@tanstack/react-query";
import { getResourcesByIdMetaOptions } from "@/client/@tanstack/react-query.gen";

export default function PerformanceView() {
  const params = useLocalSearchParams<{ id: string }>();
  const resourceQuery = useQuery(getResourcesByIdMetaOptions({
    path: {
      id: params.id
    }
  }))

  if (resourceQuery.error) {
    return <Text>Error querying resource</Text>
  }

  if (!resourceQuery.data) {
    return <Text>Resource not found</Text>
  }

  return (
    <View flex={1} items="center">
      <PdfView resource={resourceQuery.data} />
    </View>
  );
}
