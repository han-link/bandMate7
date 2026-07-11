import { Tabs } from "expo-router";
import { SymbolView } from "expo-symbols";
import type { ComponentProps } from "react";
import {
  SizableText,
  View,
  createStyledContext,
  styled,
  useMedia,
  useTheme,
  withStaticProperties,
} from "tamagui";

const TabBarContext = createStyledContext<{
  layout: "sidebar" | "bottom";
  focused: boolean;
}>({
  layout: "bottom",
  focused: false,
});

const TabBarFrame = styled(View, {
  name: "TabBar",
  context: TabBarContext,

  variants: {
    layout: {
      sidebar: {
        width: "$9",
        px: "$2",
        py: "$6",
        gap: "$2",
        borderRightWidth: 1,
        borderColor: "$borderColor",
        items: "center",
      },
      bottom: {
        flexDirection: "row",
        pt: "$2",
        px: "$2",
        borderTopWidth: 1,
        borderColor: "$borderColor",
      },
    },
  } as const,
});

const TabBarItem = styled(View, {
  name: "TabBarItem",
  context: TabBarContext,
  items: "center",
  justify: "center",
  gap: "$2",
  cursor: "pointer",
  pressStyle: { opacity: 0.6 },

  variants: {
    layout: {
      sidebar: { rounded: "$2" },
      bottom: { flex: 1, py: "$2", rounded: "$3" },
    },
    focused: {
      true: { bg: "$highlight1" },
    },
    size: {
      "...size": (val, { tokens }) => ({
        width: tokens.size[val],
        height: tokens.size[val],
      }),
    },
  } as const,
});

const TabBarLabel = styled(SizableText, {
  name: "TabBarLabel",
  context: TabBarContext,
  size: "$1",

  variants: {
    focused: {
      true: { color: "$color", fontWeight: "600" },
      false: { color: "$color11", fontWeight: "400" },
    },
  } as const,
});

const TabBar = withStaticProperties(TabBarFrame, {
  Item: TabBarItem,
  Label: TabBarLabel,
});

type NavItem = {
  route: string;
  label: string;
  symbol: ComponentProps<typeof SymbolView>["name"];
};

const NAV_ITEMS: NavItem[] = [
  {
    route: "index",
    label: "Library",
    symbol: {
      ios: "music.note.list",
      android: "library_music",
      web: "library_music",
    },
  },
  {
    route: "setlists",
    label: "Setlists",
    symbol: { ios: "bookmark", android: "bookmark", web: "bookmark" },
  },
];

type TamaguiTabBarProps = Parameters<
  NonNullable<ComponentProps<typeof Tabs>["tabBar"]>
>[0];

export function TamaguiTabBar({ state, navigation }: TamaguiTabBarProps) {
  const media = useMedia();
  const theme = useTheme();
  const layout = media.md ? "sidebar" : "bottom";

  return (
    <TabBar layout={layout}>
      {NAV_ITEMS.map((item) => {
        const index = state.routes.findIndex((r) => r.name === item.route);
        if (index === -1) return null; // route not registered — skip it

        const route = state.routes[index];
        const isFocused = state.index === index;

        const onPress = () => {
          const event = navigation.emit({
            type: "tabPress",
            target: route.key,
            canPreventDefault: true,
          });

          if (!isFocused && !event.defaultPrevented) {
            navigation.navigate(route.name, route.params);
          }
        };

        return (
          <TabBar.Item
            key={item.route}
            focused={isFocused}
            size={layout === "sidebar" ? "$6" : undefined}
            onPress={onPress}
            role="button"
            aria-selected={isFocused}
          >
            <SymbolView
              name={item.symbol}
              size={26}
              tintColor={isFocused ? theme.color.val : theme.color11.val}
            />
            <TabBar.Label>{item.label}</TabBar.Label>
          </TabBar.Item>
        );
      })}
    </TabBar>
  );
}
