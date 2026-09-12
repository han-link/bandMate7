import { useQuery } from "@tanstack/react-query";

import { getUserRolesOptions } from "@/client/@tanstack/react-query.gen";
import { Select, SelectValue } from "@/components/Select";

interface UserRoleSelectProps {
  userRoleId?: SelectValue;
  setUserRoleId: (val: SelectValue) => void;
}

export function UserRoleSelect({
  userRoleId,
  setUserRoleId,
}: UserRoleSelectProps) {
  const userRolesQuery = useQuery(getUserRolesOptions());
  const items =
    userRolesQuery.data?.map(({ id, titel }) => ({
      name: titel || "",
      value: id || "",
    })) ?? [];
  return (
    <Select
      val={userRoleId}
      setVal={setUserRoleId}
      items={items}
      label={"User Roles"}
      isLoading={userRolesQuery.isLoading}
    />
  );
}
