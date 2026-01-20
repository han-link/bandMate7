import * as React from "react";
import {z, ZodType} from "zod";
import {FormProvider, type SubmitHandler, useForm, type UseFormProps, type UseFormReturn} from "react-hook-form"
import {zodResolver} from "@hookform/resolvers/zod";

type FormProps<TSchema extends ZodType<any, any, any>> = {
    id?: string;
    schema: TSchema;
    options?: UseFormProps<z.input<TSchema>, any, z.output<TSchema>>;
    onSubmit: SubmitHandler<z.output<TSchema>>;
    className?: string;
    children: (
        methods: UseFormReturn<z.input<TSchema>, any, z.output<TSchema>>
    ) => React.ReactNode;
};

export function Form<TSchema extends ZodType<any, any, any>>(
    {
        id,
        schema,
        options,
        onSubmit,
        className = "space-y-4",
        children,
    }: FormProps<TSchema>) {
    const methods = useForm<z.input<TSchema>, any, z.output<TSchema>>({
        ...options,
        resolver: zodResolver(schema),
    });

    return (
        <FormProvider {...methods}>
            <form id={id} className={className} onSubmit={methods.handleSubmit(onSubmit)}>
                {children(methods)}
            </form>
        </FormProvider>
    );
}
