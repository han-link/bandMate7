import { useRouter } from "expo-router";
import { H1, Button, XStack } from "tamagui";

import { PerformancesGrid } from "@/features/performances/PerformancesGrid";
import { PerformancesList } from "@/features/performances/PerformancesList";
import { Select } from "@/components/Select";
import { ViewMode, ViewToggle } from "@/components/ViewToggle";
import { useState } from "react";

export default function TabOneScreen() {
  const router = useRouter();
  const [mode, setMode] = useState<ViewMode>("grid")
  return (
    <>
      <H1>Performances</H1>
      <XStack gap="$2" justify="flex-end" mb={"$4"}>
        <Button
          size="$3"
          theme="accent"
          onPress={() => router.navigate("/(tabs)/performance/create")}
        >
          Create
        </Button>
      </XStack>
      <XStack gap={"$2"} justify={"flex-end"} mb={"$4"}>
        <ViewToggle value={mode} onChange={setMode}/>
        <Select setVal={()=> {}} items={[]} />
      </XStack>
      {mode === 'grid' && <PerformancesGrid />}
      {mode === 'list' && <PerformancesList />}
    </>
  );
}
