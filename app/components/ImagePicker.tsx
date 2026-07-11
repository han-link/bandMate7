import * as DocumentPicker from "expo-document-picker";
import { Button, Image } from "tamagui";

type ImagePickerProps = {
  value?: DocumentPicker.DocumentPickerAsset;
  onPick: (asset: DocumentPicker.DocumentPickerAsset) => void;
};

export function ImagePicker({ value, onPick }: ImagePickerProps) {
  const pickImage = async () => {
    try {
      const result = await DocumentPicker.getDocumentAsync({ type: "image/*" });

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
      <Button onPress={pickImage}>Pick an Image</Button>
      {value && (
        <Image
          src={value.uri}
          width="$10"
          height="$10"
          objectFit="cover"
          rounded="$2"
        />
      )}
    </>
  );
}
