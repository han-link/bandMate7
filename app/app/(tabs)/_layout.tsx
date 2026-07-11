import { Tabs } from "expo-router";
import { getTokenValue, useMedia, useTheme } from "tamagui";

import { TamaguiTabBar } from "@/components/TamaguiTabBar";

export default function TabLayout() {
  const media = useMedia();
  const theme = useTheme();

  return (
    <Tabs
      tabBar={(props) => <TamaguiTabBar {...props} />}
      screenOptions={{
        tabBarPosition: media.md ? "left" : "bottom",
        headerShown: false,
        sceneStyle: {
          backgroundColor: theme.background.val,
          padding: getTokenValue("$6", "space"),
        },
      }}
    ></Tabs>
  );
}
