import { useEffect, useState } from "react";
import { File, Paths } from "expo-file-system";
import Pdf from "react-native-pdf";
import { Paragraph, View, useWindowDimensions } from "tamagui";
import { Resource } from "@/client";

export interface PdfViewProps {
  resource: Resource;
}

export default function PdfView(props: PdfViewProps) {
  const { width, height } = useWindowDimensions();
  const [localUri, setLocalUri] = useState<string>();
  const [error, setError] = useState<string>();

  useEffect(() => {
    let cancelled = false;

    (async () => {
      try {
        const file = new File(Paths.cache, `pdf-${props.resource.id}.pdf`);
        if (!file.exists) {
          await File.downloadFileAsync(props.resource.url ?? "", file);
        }
        if (!cancelled) setLocalUri(file.uri);
      } catch (e) {
        if (!cancelled) setError(String(e));
      }
    })();

    return () => { cancelled = true; };
  }, [props]);

  if (error) return <Paragraph>Couldn’t load PDF: {error}</Paragraph>;
  if (!localUri) return <Paragraph>Loading …</Paragraph>;

  return (
    <View style={{ flex: 1 }}>
      <Pdf source={{ uri: localUri }} enablePaging={true} style={{ flex: 1, width, height }} />
    </View>
  );
}