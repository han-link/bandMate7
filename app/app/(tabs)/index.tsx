import { useRouter } from "expo-router";
import { H1, Button, XStack } from "tamagui";

import { PerformancesList } from "@/features/performances/PerformancesList";

export default function TabOneScreen() {
  const router = useRouter();

  return (
    <>
      <H1>Performances</H1>
      <XStack gap="$2" justify="flex-end">
        <Button
          size="$3"
          theme="accent"
          onPress={() => router.navigate("/(tabs)/performance/create")}
        >
          Create
        </Button>
      </XStack>
      <PerformancesList />
    </>
  );
}
