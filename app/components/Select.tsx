import { Check, ChevronDown } from "@tamagui/lucide-icons-2";
import {
  Adapt,
  Paragraph,
  Select as TamaguiSelect,
  SelectProps as TamaguiSelectProps,
  Sheet,
} from "tamagui";

interface Value {
  name: string;
  value: string;
}

export type SelectValue = Lowercase<Value["value"]>;

interface SelectProps {
  val?: SelectValue;
  setVal: (val: SelectValue) => void;
  items: Value[];
  label: string;
  isLoading?: boolean;
}

export function Select({
  isLoading = false,
  val,
  setVal,
  items,
  label,
  ...props
}: TamaguiSelectProps<SelectValue> & SelectProps) {
  if (isLoading) return <Paragraph>Loading …</Paragraph>;
  return (
    <TamaguiSelect
      value={val}
      onValueChange={setVal}
      disablePreventBodyScroll
      {...props}
    >
      <TamaguiSelect.Trigger
        iconAfter={ChevronDown}
        borderRadius="$4"
        backgroundColor="$background"
      >
        <TamaguiSelect.Value placeholder="Something" />
      </TamaguiSelect.Trigger>

      <Adapt platform="touch">
        <Sheet
          modal
          dismissOnSnapToBottom
          transition="medium"
          snapPointsMode={"fit"}
        >
          <Sheet.Frame>
            <Sheet.ScrollView>
              <Adapt.Contents />
            </Sheet.ScrollView>
          </Sheet.Frame>
          <Sheet.Overlay backgroundColor="$shadow6" />
        </Sheet>
      </Adapt>

      <TamaguiSelect.Content>
        <TamaguiSelect.Viewport
          rounded="$4"
          borderWidth={1}
          borderColor="$borderColor"
        >
          <TamaguiSelect.Group>
            <TamaguiSelect.Label fontWeight="700">
              {label}
            </TamaguiSelect.Label>
            {items.map((item, i) => {
              return (
                <TamaguiSelect.Item
                  index={i}
                  key={item.name}
                  value={item.value}
                >
                  <TamaguiSelect.ItemText>{item.name}</TamaguiSelect.ItemText>
                  <TamaguiSelect.ItemIndicator marginLeft="auto">
                    <Check size={16} />
                  </TamaguiSelect.ItemIndicator>
                </TamaguiSelect.Item>
              );
            })}
          </TamaguiSelect.Group>
        </TamaguiSelect.Viewport>
      </TamaguiSelect.Content>
    </TamaguiSelect>
  );
}
