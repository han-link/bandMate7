type HeadProps = {
    title?: string;
};

export function Head(props: HeadProps = {}) {
    const title = props.title ? `${props.title} | BandMate7` : 'BandMate7'
    return (
        <>
            <meta charSet="utf-8"/>
            <title>{title}</title>
        </>
    );
}