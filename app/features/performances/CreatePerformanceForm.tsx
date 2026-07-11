import { useMutation, useQueryClient } from "@tanstack/react-query";
import { DocumentPickerAsset } from "expo-document-picker";
import { useRouter } from "expo-router";
import { useState } from "react";
import { Platform } from "react-native";
import { Form, Input, Label, Button, Paragraph } from "tamagui";
import * as z from "zod";

import {
  postPerformancesMutation,
  getPerformancesQueryKey,
} from "@/client/@tanstack/react-query.gen";
import { zPostPerformancesBody } from "@/client/zod.gen";
import { FilePicker } from "@/components/FilePicker";
import { ImagePicker } from "@/components/ImagePicker";
import { SelectValue } from "@/components/Select";

import { UserRoleSelect } from "./UserRoleSelect";

const createPerformanceSchema = zPostPerformancesBody
  .omit({ cover: true, file: true })
  .extend({
    name: z.string().trim().min(1, "Give the performance a name"),
    bpm: z
      .number({ message: "BPM must be a number" })
      .int("BPM must be a whole number")
      .optional(),
    userRoleId: z.string().uuid("Select a role").optional(),
  });

type FormErrors = Partial<
  Record<keyof z.infer<typeof createPerformanceSchema>, string[]>
>;

const FieldError = ({ messages }: { messages?: string[] }) =>
  messages?.length ? <Paragraph>{messages[0]}</Paragraph> : null;
const multipartBodySerializer = (body: unknown) => {
  const data = new FormData();
  Object.entries(body as Record<string, unknown>).forEach(([key, value]) => {
    if (value === undefined || value === null) return;
    data.append(key, value as string | Blob);
  });
  return data;
};

const toUploadPart = (asset: DocumentPickerAsset) =>
  Platform.OS === "web"
    ? asset.file
    : ({
        uri: asset.uri,
        name: asset.name,
        type: asset.mimeType ?? "application/octet-stream",
      } as unknown as Blob);

export function CreatePerformanceForm() {
  const [name, setName] = useState("");
  const [bpm, setBpm] = useState("");
  const [cover, setCover] = useState<DocumentPickerAsset>();
  const [file, setFile] = useState<DocumentPickerAsset>();
  const [userRoleId, setUserRoleId] = useState<SelectValue>("");
  const [errors, setErrors] = useState<FormErrors>({});
  const queryClient = useQueryClient();
  const router = useRouter();

  const createPerformance = useMutation(postPerformancesMutation());

  const clearError = (field: keyof FormErrors) =>
    setErrors((prev) => (prev[field] ? { ...prev, [field]: undefined } : prev));

  const handleSubmit = () => {
    const parsed = createPerformanceSchema.safeParse({
      name,
      bpm: bpm ? Number(bpm) : undefined,
      userRoleId: userRoleId || undefined,
    });
    if (!parsed.success) {
      setErrors(parsed.error.flatten().fieldErrors);
      return;
    }
    setErrors({});
    createPerformance.mutate(
      {
        body: {
          ...parsed.data,
          cover: cover && toUploadPart(cover),
          file: file && toUploadPart(file),
        },
        bodySerializer: multipartBodySerializer,
      },
      {
        onSuccess: () => {
          setName("");
          setBpm("");
          setCover(undefined);
          setFile(undefined);
          setUserRoleId("");
          queryClient.invalidateQueries({
            queryKey: getPerformancesQueryKey(),
          });
          router.navigate("/");
        },
      },
    );
  };
  return (
    <Form width="100%" gap="$1.5" onSubmit={handleSubmit}>
      <Label htmlFor="name">Name</Label>
      <Input
        id="name"
        placeholder="Narcotic"
        value={name}
        onChangeText={(text) => {
          setName(text);
          clearError("name");
        }}
      />
      <FieldError messages={errors.name} />
      <Label htmlFor="bpm">Bpm</Label>
      <Input
        id="bpm"
        placeholder="22"
        value={bpm}
        keyboardType="numeric"
        onChangeText={(text) => {
          setBpm(text);
          clearError("bpm");
        }}
      />
      <FieldError messages={errors.bpm} />
      <ImagePicker value={cover} onPick={setCover} />
      <FilePicker value={file} onPick={setFile} />
      <UserRoleSelect
        userRoleId={userRoleId}
        setUserRoleId={(val) => {
          setUserRoleId(val);
          clearError("userRoleId");
        }}
      />
      <FieldError messages={errors.userRoleId} />
      <Form.Trigger asChild>
        <Button>{createPerformance.isPending ? "Creating…" : "Create"}</Button>
      </Form.Trigger>
      {createPerformance.isSuccess && (
        <Paragraph>Performance created</Paragraph>
      )}
      {createPerformance.isError && (
        <Paragraph>
          {createPerformance.error.message}{" "}
          {createPerformance.failureReason?.response?.data.error}
        </Paragraph>
      )}
    </Form>
  );
}
