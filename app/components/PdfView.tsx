import Pdf from "react-native-pdf";
import { ScrollView, useWindowDimensions } from "tamagui";

export default function PdfView({ uri }: { uri: string }) {
  const { width, height } = useWindowDimensions();
  const source = {
    uri: uri,
    cache: true,
  };
  return (
    <ScrollView style={{ flex: 1 }} contentInsetAdjustmentBehavior="automatic">
      <Pdf source={source} style={{ flex: 1, width, height }} />
    </ScrollView>
  );
}
