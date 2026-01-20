import * as React from 'react';

import {Head} from '../seo';

type ContentLayoutProps = {
    children: React.ReactNode;
    title: string;
};

export function ContentLayout({children, title}: ContentLayoutProps) {
    return (
        <>
            <Head title={title}/>
            <div className="py-4">
                <div className="mx-8 mb-4">
                    <h1 className="text-2xl font-semibold">{title}</h1>
                </div>
                <div className="mx-8">
                    {children}
                </div>
            </div>
        </>
    );
};