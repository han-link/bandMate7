import { useRouter } from "expo-router";
import { Card, Image, Paragraph, styled } from "tamagui";

import { Performance } from "@/client";

const CardFrame = styled(Card, {
  name: "PerformanceCard",
  borderWidth: 1,
  borderColor: "$borderColor",
  width: "$15",
  height: "$15",
  overflow: "hidden",
  cursor: "pointer",
});

interface PerformanceCardProps {
  performance: Performance;
}

export function PerformanceCard({ performance: p }: PerformanceCardProps) {
  const router = useRouter();
  const fistPdf =
    p.resources?.filter((resource) => resource.type === "application/pdf") ??
    [];
  return (
    <CardFrame
      size="$3"
      onPress={() => router.navigate(`/resource/${fistPdf[0]?.id}`)}
    >
      <Card.Header p="$4">
        <Paragraph>{p.name}</Paragraph>
      </Card.Header>
      <Card.Background items="center" unstyled={true}>
        <Image
          objectFit="cover"
          width="100%"
          height="100%"
          src={p.cover?.url}
        />
      </Card.Background>
    </CardFrame>
  );
}
