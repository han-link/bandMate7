import {useUser} from "@/lib/auth.ts";
import {ContentLayout} from "@/components/layouts/content-layout"

type EntryProps = {
    label: string;
    value: string;
};

function Entry({label, value}: EntryProps) {
    return (
        <div className="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt className="text-sm font-medium">{label}</dt>
            <dd className="mt-1 text-sm  sm:col-span-2 sm:mt-0">
                {value}
            </dd>
        </div>
    )
}


export function ProfileRoute() {
    const user = useUser();

    if (!user.data) return null;

    return (
        <ContentLayout title={"Profile"}>
            <div className="">
                <div className="px-4 py-5">
                    <dl className="sm:divide-y divide-gray-200">
                        <Entry label="First Name" value={user.data.firstName}/>
                        <Entry label="Last Name" value={user.data.lastName}/>
                        <Entry label="Email Address" value={user.data.email}/>
                        <Entry label="Role" value={user.data.role.name}/>
                    </dl>
                </div>
            </div>
        </ContentLayout>
    )
}