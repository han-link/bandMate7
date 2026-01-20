import {useNotifications} from "@/components/ui/notifications";
import {Button} from "@/components/ui/button"

import {createPerformanceInputSchema, useCreatePerformance} from "../api/create-performance";
import {Form} from "@/components/ui/form/form.tsx";
import {Input} from "@/components/ui/form/input.tsx";
import {useNavigate} from "react-router-dom";

type CreatePerformanceProps = {
    onSuccessTo: string;
}

export function CreatePerformance(props: CreatePerformanceProps) {
    const {addNotification} = useNotifications();
    const navigate = useNavigate();
    const createPerformanceMutation = useCreatePerformance({
        mutationConfig: {
            onSuccess: (data) => {
                addNotification({
                    type: "success",
                    title: `Created Performance ${data.name}`
                })
                navigate(props.onSuccessTo)
            }
        }
    })

    return (
        <Form
            id={"create-performance"}
            onSubmit={(values) => {
                const fd = new FormData();
                fd.append("name", values.name);

                if (values.cover instanceof FileList) {
                    fd.append("cover", values.cover[0]);
                }

                if (values.file instanceof FileList) {
                    fd.append("file", values.file[0]);
                }

                createPerformanceMutation.mutate({data: fd});
            }}
            schema={createPerformanceInputSchema}
        >
            {({register, formState}) => (
                <>
                    <Input
                        name={"Name"}
                        error={formState.errors.name}
                        registration={register('name')}
                    />
                    <Input
                        name={"Bpm"}
                        type={"number"}
                        error={formState.errors.bpm}
                        registration={register('bpm', {
                            setValueAs: (v) => v === "" ? null : +v
                        })}
                    />
                    <Input
                        name={"Cover"}
                        type="file"
                        accept="image/*"
                        error={formState.errors.cover}
                        registration={register("cover")}
                    />
                    <Input
                        name={"File"}
                        type="file"
                        error={formState.errors.file}
                        registration={register("file")}
                    />
                    <Button type={"submit"} className={"btn-primary"} isLoading={createPerformanceMutation.isPending}>
                        Create
                    </Button>
                </>
            )}
        </Form>
    )
}