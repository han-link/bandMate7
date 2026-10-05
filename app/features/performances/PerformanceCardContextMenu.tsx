import { ContextMenu } from '@tamagui/context-menu'
import { Performance } from "@/client";
import { ReactNode } from 'react';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import {
    deletePerformancesByIdMutation,
    getPerformancesByIdQueryKey,
    getPerformancesQueryKey,
} from '@/client/@tanstack/react-query.gen';

interface PerformanceCardContextMenuProps {
    performance: Performance;
    children: ReactNode | ReactNode[];
}

export function PerformanceCardContextMenu({ performance: p, children }: PerformanceCardContextMenuProps) {
    const queryClient = useQueryClient()
    const { mutate: deletePerformance } = useMutation({
        ...deletePerformancesByIdMutation(),
        onSuccess: (_data, { path }) => {
            queryClient.removeQueries({ queryKey: getPerformancesByIdQueryKey({ path }) });
            queryClient.invalidateQueries({ queryKey: getPerformancesQueryKey() });
        },
    });

    const handleDelete = () => {
        if (!p.id) return;
        deletePerformance({ path: { id: p.id } });
    };
    return (
        <ContextMenu>
            <ContextMenu.Trigger>
                {children}
            </ContextMenu.Trigger>
            <ContextMenu.Portal zIndex={100}>
                <ContextMenu.Content
                    p="$1.5"
                    minW={180}
                    borderWidth={1}
                    borderColor="$borderColor"
                    transformOrigin="left top"
                    enterStyle={{ scale: 0.9, opacity: 0, y: -5 }}
                    exitStyle={{ scale: 0.95, opacity: 0, y: -3 }}
                    elevation="$3"
                    transition="100ms"
                >
                    <ContextMenu.Item
                        onSelect={handleDelete}
                        destructive>
                        <ContextMenu.ItemTitle color="red">Delete {p.titel}</ContextMenu.ItemTitle>
                    </ContextMenu.Item>
                    <ContextMenu.Item
                        disabled={true}>
                        <ContextMenu.ItemTitle>Edit</ContextMenu.ItemTitle>
                    </ContextMenu.Item>
                </ContextMenu.Content>
            </ContextMenu.Portal>
        </ContextMenu >
    )
}
