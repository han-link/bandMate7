import { useQuery } from "@tanstack/react-query";
import { api } from "./api-client";

type Role = { id: string; name: string };
type User = { email: string; firstName: string; lastName: string; role: Role };

const DEV_BASE_USER = {
    email: "dev@example.com",
    firstName: "Dev",
    lastName: "User",
};

const getRoles = (): Promise<{
    data: Role[];
}> => {
    return api.get(`/userRoles`);
};

export const useUser = () =>
    useQuery({
        queryKey: ["dev-user", "userRoles"],
        queryFn: async () => {
            const response = await getRoles();
            const firstRole = response.data[0] ?? { id: "no-role", name: "no-role" };

            const user: User = { ...DEV_BASE_USER, role: firstRole };
            return user;
        },
        staleTime: Infinity,
        retry: false,
    });
