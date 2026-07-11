import * as DocumentPicker from "expo-document-picker";
import { Button, Paragraph } from "tamagui";

type FilePickerProps = {
  value?: DocumentPicker.DocumentPickerAsset;
  onPick: (asset: DocumentPicker.DocumentPickerAsset) => void;
};

export function FilePicker({ value, onPick }: FilePickerProps) {
  const pickFile = async () => {
    try {
      const result = await DocumentPicker.getDocumentAsync();

      if (!result.canceled) {
        onPick(result.assets[0]);
      } else {
        console.warn("Document selection cancelled.");
      }
    } catch (error) {
      console.error("Error picking documents:", error);
    }
  };
  return (
    <>
      <Button onPress={pickFile}>Pick a File</Button>
      <Paragraph>{value?.name}</Paragraph>
    </>
  );
}
