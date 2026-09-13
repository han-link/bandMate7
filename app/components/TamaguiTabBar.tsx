import { useTabTrigger } from "expo-router/ui";
import { SymbolView } from "expo-symbols";
import type { ComponentProps } from "react";
import {
  SizableText,
  View,
  createStyledContext,
  styled,
  useTheme,
  withStaticProperties,
} from "tamagui";

const TabBarContext = createStyledContext<{
  focused: boolean;
}>({
  focused: false,
});

const TabBarFrame = styled(View, {
  name: "TabBar",
  context: TabBarContext,
  flexDirection: "row",
  pt: "$2",
  px: "$2",
  borderTopWidth: 1,
  borderColor: "$borderColor",

  $md: {
    flexDirection: "column",
    width: "$9",
    py: "$6",
    gap: "$2",
    borderTopWidth: 0,
    borderRightWidth: 1,
    items: "center",
  },
});

const TabBarItem = styled(View, {
  name: "TabBarItem",
  context: TabBarContext,
  items: "center",
  justify: "center",
  gap: "$2",
  cursor: "pointer",
  pressStyle: { opacity: 0.6 },
  grow: 1,
  flexBasis: 0,
  py: "$2",
  rounded: "$3",

  $md: {
    grow: 0,
    flexBasis: "auto",
    py: 0,
    width: "$6",
    height: "$6",
    rounded: "$2",
  },

  variants: {
    focused: {
      true: { bg: "$highlight1" },
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

function TabBarTrigger({ item }: { item: NavItem }) {
  const theme = useTheme();
  const { trigger, triggerProps } = useTabTrigger({ name: item.route });
  if (!trigger) return null; // route not registered - skip it

  const { isFocused, onPress, onLongPress } = triggerProps;

  return (
    <TabBar.Item
      focused={isFocused}
      onPress={onPress}
      onLongPress={onLongPress}
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
}

export function TamaguiTabBar() {
  return (
    <TabBar>
      {NAV_ITEMS.map((item) => (
        <TabBarTrigger key={item.route} item={item} />
      ))}
    </TabBar>
  );
}
