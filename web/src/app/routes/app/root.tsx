import { Outlet } from 'react-router';

import { DashboardLayout } from '@/components/layouts';


export function AppRoot() {
    return (
        <DashboardLayout>
            <Outlet />
        </DashboardLayout>
    );
}