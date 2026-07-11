import { useLocalSearchParams } from "expo-router";
import { View } from "tamagui";

import PdfView from "@/components/PdfView";

export default function PerformanceView() {
  const params = useLocalSearchParams<{ id: string }>();
  const host = process.env.EXPO_PUBLIC_API_URL ?? "http://localhost:8080";
  const source = `${host}/api/v1/resources/${params.id}`;
  return (
    <View flex={1} items="center">
      <PdfView uri={source} />
    </View>
  );
}
