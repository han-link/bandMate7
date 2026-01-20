import type {Performance} from "@/types/api";
// TODO: Remove heroicons??
import {EllipsisHorizontalIcon} from '@heroicons/react/24/solid'
import {useNotifications} from "@/components/ui/notifications";
import {useDeletePerformance} from "../api/delete-performance.ts";
import {Cover} from "./performance-cover.tsx";

interface PerformanceCardProps {
    performance: Performance;
}

export function PerformanceCard(props: PerformanceCardProps) {
    const {addNotification} = useNotifications();
    const deletePerformanceMutation = useDeletePerformance({
        mutationConfig: {
            onSuccess: () => {
                addNotification({
                    type: 'success',
                    title: `Deleted Performance ${props.performance.name}`,
                });
            },
        },
    });

    return (
        <div className="card bg-black w-60 shadow-sm">
            <figure>
                <Cover cover={props.performance.cover}/>
            </figure>
            <div className="card-body">
                <h2 className="card-title text-white">
                    {props.performance.name}
                </h2>
                <h3>&ndash; Robbie Wiliams</h3>
                <div className="dropdown dropdown-start">
                    <button tabIndex={0} className="btn btn-xs btn-ghost w-fit">
                        <EllipsisHorizontalIcon className="size-6"/>
                    </button>
                    <ul tabIndex={-1} className="dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm">
                        <li>
                            <button
                                disabled={deletePerformanceMutation.isPending}
                                onClick={() =>
                                    deletePerformanceMutation.mutate({performanceId: props.performance.id})
                                }
                            >
                                Delete
                            </button>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    )
}