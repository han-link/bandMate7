import { defaultConfig } from "@tamagui/config/v5";
import { animationsReactNative } from "@tamagui/config/v5-rn";
import { createTamagui } from "tamagui";

export const config = createTamagui({
  ...defaultConfig,
  animations: animationsReactNative,
});

type AppConfig = typeof config;

declare module "tamagui" {
  // eslint-disable-next-line @typescript-eslint/no-empty-object-type
  interface TamaguiCustomConfig extends AppConfig {}
}
