import { useEffect, useMemo } from "react";
import { MusicalNoteIcon } from "@heroicons/react/24/solid";
import { useCover } from "../api/get-cover";

type Resource = { id: string };

export function Cover(props: { cover?: Resource }) {
    const resourceId = props.cover?.id;

    const { data: blob, isLoading, isError } = useCover({ resourceId });

    const objectUrl = useMemo(() => {
        if (!blob) return null;
        return URL.createObjectURL(blob);
    }, [blob]);

    useEffect(() => {
        if (!objectUrl) return;
        return () => URL.revokeObjectURL(objectUrl);
    }, [objectUrl]);

    if (!props.cover) {
        return (
            <div className="size-60 inline-flex justify-center items-center">
                <MusicalNoteIcon className="size-16" />
            </div>
        );
    }

    if (isLoading) return <div className="skeleton size-60" />;
    if (isError || !objectUrl) return <div className="skeleton size-60" />;

    return (
        <div className="size-60 relative overflow-hidden rounded">
            <img className="size-60 object-cover" src={objectUrl} alt="Album cover" />
        </div>
    );
}
