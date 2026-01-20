const join = (...parts: string[]) =>
    "/" +
    parts
        .filter(Boolean)
        .join("/")
        .replace(/\/+/g, "/")
        .replace(/^\/|\/$/g, "");

export const paths = {
    app: {
        root: {
            path: "/",
            href: () => "/",
        },

        dashboard: {
            path: "dashboard",
            href: () => join("dashboard"),
        },

        performances: {
            list: {
                path: "performances",
                href: () => join("performances"),
            },
            create: {
                path: "performances/create",
                href: () => join("performances", "create"),
            },
        },

        profile: {
            path: "profile",
            href: () => join("profile"),
        },
    }
} as const;
