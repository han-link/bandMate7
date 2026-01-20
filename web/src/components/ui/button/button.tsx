import {cn} from '@/utils/cn';
import * as React from "react";

type ButtonProps = {
    type?: "submit" | "reset" | "button" | undefined;
    className?: string;
    isLoading?: boolean;
    children: React.ReactNode;
}

export function Button(props: ButtonProps) {
    return (
        <button className={cn("btn", props.className)} type={props.type}>
            {props.isLoading && <span className="loading loading-spinner"></span>}
            {props.children}
        </button>
    )
}