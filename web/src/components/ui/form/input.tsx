import * as React from "react";
import type {FieldError, UseFormRegisterReturn, FieldErrorsImpl, Merge} from "react-hook-form";
import {cn} from "@/utils/cn";

export type InputProps = React.InputHTMLAttributes<HTMLInputElement> & {
    className?: string;
    registration?: Partial<UseFormRegisterReturn>;
    onFilesChange?: (files: FileList | null) => void;
    error?: FieldError | Merge<FieldError, FieldErrorsImpl<never>>;
};

export const Input = React.forwardRef<HTMLInputElement, InputProps>(
    ({className, type = "text", error, registration, onFilesChange, onChange, name, ...props}, ref) => {
        const isFile = type === "file";

        const baseClass = isFile
            ? "file-input w-full"
            : "input w-full";

        const errorClass = isFile ? "file-input-error" : "input-error";

        return (
            <>
                <label htmlFor={name} className="label">{name}</label>
                <input
                    ref={ref}
                    id={name}
                    name={name}
                    type={type}
                    className={cn(baseClass, error ? errorClass : "", className)}
                    {...registration}
                    {...props}
                    onChange={(e) => {
                        registration?.onChange?.(e);
                        if (isFile) onFilesChange?.((e.target as HTMLInputElement).files);
                        onChange?.(e);
                    }}
                />
                {error && <p className="text-error text-sm mt-1">{error.message}</p>}
            </>
        );
    }
);

Input.displayName = "Input";
