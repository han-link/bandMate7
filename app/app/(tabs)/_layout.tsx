// tamagui-ignore
import {
  TabSlot,
  type UseTabsWithTriggersOptions,
  useTabsWithTriggers,
} from "expo-router/ui";
import { View } from "tamagui";

import { TamaguiTabBar } from "@/components/TamaguiTabBar";

const TRIGGERS: UseTabsWithTriggersOptions["triggers"] = [
  { type: "internal", name: "index", href: "/" },
  { type: "internal", name: "setlists", href: "/setlists" },
  { type: "internal", name: "performance/create", href: "/performance/create" },
];

export default function TabLayout() {
  const { NavigationContent } = useTabsWithTriggers({ triggers: TRIGGERS });

  return (
    <NavigationContent>
      <View
        flex={1}
        flexDirection="column-reverse"
        $md={{ flexDirection: "row" }}
      >
        <TamaguiTabBar />
        <View flex={1} overflow="hidden" bg="$background" p="$6">
          <TabSlot />
        </View>
      </View>
    </NavigationContent>
  );
}
