import { Link, Stack } from "expo-router";
import { Paragraph, YStack } from "tamagui";

export default function NotFoundScreen() {
  return (
    <>
      <Stack.Screen options={{ title: "Oops!" }} />
      <YStack flex={1} items="center" justify="center" p="$5" gap="$4">
        <Paragraph size="$6" fontWeight="bold">
          This screen doesn&apos;t exist.
        </Paragraph>

        <Link href="/">
          <Paragraph>Go to home screen!</Paragraph>
        </Link>
      </YStack>
    </>
  );
}
