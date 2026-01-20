import {PlusIcon, RectangleStackIcon} from "@heroicons/react/24/solid";
import {NavLink, Link} from "react-router";
import * as React from "react";
import {type ElementType, useState} from "react";

import {SidebarIcon} from "@/components/ui/icons";
import {paths} from "@/config/paths.ts";
import {cn} from '@/utils/cn';

interface NavItem {
    to: string;
    name: string;
    Icon: ElementType;
    onClick?: () => void;
}

function NavButton({to, name, Icon, onClick}: NavItem) {
    return (
        <NavLink to={to}
                 className={({isActive}) =>
                     cn(
                         "is-drawer-close:tooltip is-drawer-close:tooltip-right",
                         isActive && "bg-primary"
                     )
                 }
                 data-tip={name}
                 onClick={onClick}
        >
            <Icon className="my-1.5 inline-block size-4"/>
            <span className="is-drawer-close:hidden">{name}</span>
        </NavLink>
    )
}

export function DashboardLayout({children}: { children: React.ReactNode }) {
    const [open, setOpen] = useState(false);
    const navigation: NavItem[] = [
        {name: 'Dashboard', to: paths.app.dashboard.href(), Icon: RectangleStackIcon},
    ];

    return (
        <>
            <div className="drawer lg:drawer-open">
                <input
                    id="my-drawer-4"
                    type="checkbox"
                    className="drawer-toggle"
                    checked={open}
                    onChange={(e) => setOpen(e.target.checked)}
                />
                <div className="drawer-content">
                    <nav className="navbar w-full bg-base-300 flex">
                        <label htmlFor="my-drawer-4" aria-label="open sidebar" className="btn btn-square btn-ghost">
                            <SidebarIcon/>
                        </label>
                        <div className="px-4">BandMate7</div>
                        <Link to={paths.app.profile.href()} className="avatar ml-auto px-4">
                            <div className="w-8 rounded">
                                <img
                                    src="https://img.daisyui.com/images/profile/demo/superperson@192.webp"
                                    alt="Tailwind-CSS-Avatar-component"
                                />
                            </div>
                        </Link>
                    </nav>
                    <div className="p-4">
                        {children}
                    </div>
                </div>

                <div className="drawer-side is-drawer-close:overflow-visible">
                    <label htmlFor="my-drawer-4" aria-label="close sidebar" className="drawer-overlay"></label>
                    <div
                        className="flex min-h-full flex-col items-start bg-base-200 is-drawer-close:w-14 is-drawer-open:w-64">
                        <ul className="menu w-full grow">
                            <li>
                                <NavButton to={paths.app.performances.create.href()} name={"Create"} Icon={PlusIcon} onClick={() => setOpen(false)} />
                            </li>
                            <div className="divider my-0.5"></div>
                            {navigation.map((item: NavItem) => (
                                <li key={item.name}>
                                    <NavButton to={item.to} name={item.name} Icon={item.Icon} onClick={() => setOpen(false)}/>
                                </li>
                            ))}
                        </ul>
                    </div>
                </div>
            </div>
        </>
    )
}