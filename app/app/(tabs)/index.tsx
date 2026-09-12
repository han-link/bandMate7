import { useRouter } from "expo-router";
import { H1, Button, XStack } from "tamagui";

import { PerformancesGrid } from "@/features/performances/PerformancesGrid";
import { PerformancesList } from "@/features/performances/PerformancesList";
import { Select, SelectValue } from "@/components/Select";
import { ViewMode, ViewToggle } from "@/components/ViewToggle";
import { ArrowDownWideNarrow, ArrowUpWideNarrow } from "@tamagui/lucide-icons-2";
import { useState } from "react";

export default function TabOneScreen() {
  const router = useRouter();
  const sortOptions = [{name: "Titel", value: "titel"}, {name: "Created", value: "created_at"}]
  const [mode, setMode] = useState<ViewMode>("grid")
  const [desc, setDesc] = useState<boolean>(false)
  const [orderBy, setOrderBy] = useState<SelectValue>("titel")
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
        <Select val={orderBy} setVal={setOrderBy} items={sortOptions} />
          <Button onPress={() => setDesc(!desc)}>
            { desc ? <ArrowDownWideNarrow/> : <ArrowUpWideNarrow /> }
          </Button>
      </XStack>
      {mode === 'grid' && <PerformancesGrid orderBy={orderBy} desc={desc} />}
      {mode === 'list' && <PerformancesList orderBy={orderBy} desc={desc} />}
    </>
  );
}
