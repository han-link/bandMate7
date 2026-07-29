import { ToggleGroup, XGroup } from 'tamagui'
import { LayoutGrid, List } from '@tamagui/lucide-icons-2'

export type ViewMode = "grid" | "list";

interface ViewToggleProps {
    value: ViewMode
    onChange: (mode: ViewMode) => void
}

export function ViewToggle({ value, onChange }: ViewToggleProps) {
    return (
        <ToggleGroup type="single" disableDeactivation={true} value={value} onValueChange={(v) => v && onChange(v as ViewMode)}>
            <XGroup>
                <XGroup.Item>
                    <ToggleGroup.Item value="grid" borderRadius="$4" activeStyle={{ bg: '$color5' }}>
                        <LayoutGrid />
                    </ToggleGroup.Item>
                </XGroup.Item>
                <XGroup.Item>
                    <ToggleGroup.Item value="list" borderRadius="$4" activeStyle={{ bg: '$color5' }}>
                        <List />
                    </ToggleGroup.Item>
                </XGroup.Item>
            </XGroup>
        </ToggleGroup>
    )
}