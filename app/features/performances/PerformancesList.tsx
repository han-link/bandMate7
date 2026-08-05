import { getPerformancesOptions } from "@/client/@tanstack/react-query.gen"
import { Performance, Resource } from "@/client";
import { Table } from "@/components/tableParts"
import { useQuery } from "@tanstack/react-query"
import { createColumnHelper, useReactTable, getCoreRowModel, flexRender } from "@tanstack/react-table"
import { useRef } from "react"
import { Avatar, Text, Paragraph } from "tamagui"
import { useRouter } from "expo-router";
import { Pagination } from "@/app/types/pagination";

type Item = {
    coverUrl: string
    name: string
    createdAt: string
    resourceId?: string
}

const columnHelper = createColumnHelper<Item>()

const dateFmt = new Intl.DateTimeFormat(undefined, {
    year: 'numeric', month: 'short', day: 'numeric',
})

const columns = [
    columnHelper.accessor('coverUrl', {
        id: 'cover',
        header: () => 'Cover',
        cell: (info) => (
            <Avatar size="$5" rounded="$3">
                <Avatar.Image alt="Cover" src={info.getValue()} />
                <Avatar.Fallback bg="$color6" />
            </Avatar>
        ),
    }),
    columnHelper.accessor('name', {
        header: () => 'Name',
        cell: (info) => <Text>{info.getValue()}</Text>,
    }),
    columnHelper.accessor('createdAt', {
        header: () => 'Created',
        cell: (info) => <Text>{dateFmt.format(new Date(info.getValue()))}</Text>,
    }),
]

interface PerformancesListProps extends Pagination {}

export function PerformancesList(props: PerformancesListProps) {
    const router = useRouter();
    const performancesQuery = useQuery(getPerformancesOptions({
        query: {
            desc: props.desc,
            orderBy: props.orderBy
        }
    }));

    if (performancesQuery.isLoading) return <Paragraph>Loading …</Paragraph>;

    if (performancesQuery.data?.length === 0)
        return <Paragraph>No Performances</Paragraph>;

    if (performancesQuery.isError || performancesQuery.data === undefined)
        return <Paragraph>Query Error</Paragraph>;

    const fistPdfId = (p: Performance): Resource | undefined =>
        p.resources?.filter((resource) => resource.type === "application/pdf")[0];

    const data: Item[] = performancesQuery.data.map((p) => ({ coverUrl: p.cover?.url || "", name: p.titel || "", createdAt: p.createdAt || "", resourceId: fistPdfId(p)?.id }))

    const table = useReactTable({
        data: data,
        columns,
        getCoreRowModel: getCoreRowModel()
    })

    const headerGroups = table.getHeaderGroups()
    const tableRows = table.getRowModel().rows
    const footerGroups = table.getFooterGroups()

    const allRowsLength = tableRows.length + headerGroups.length + footerGroups.length
    const rowCounter = useRef(-1)
    rowCounter.current = -1



    return (
        <Table
            alignCells={{ x: 'center', y: 'center' }}
            alignHeaderCells={{ x: 'center', y: 'center' }}
            cellWidth={"$18"}
            cellHeight={"$8"}
            borderWidth={0}
            p={"$2"}
            maxW={"100%"}
            maxH={600}
            gap={"$5"}
        >
            <Table.Head>
                {headerGroups.map((headerGroup) => {
                    rowCounter.current++
                    return (
                        <Table.Row
                            rowLocation={
                                rowCounter.current === 0
                                    ? 'first'
                                    : rowCounter.current === allRowsLength - 1
                                        ? 'last'
                                        : 'middle'
                            }
                            key={headerGroup.id}
                            justify={'flex-start'}
                        >
                            {headerGroup.headers.map((header) => (
                                <Table.HeaderCell
                                    cellLocation={
                                        header.id === 'fullName'
                                            ? 'first'
                                            : header.id === 'role'
                                                ? 'last'
                                                : 'middle'
                                    }
                                    key={header.id}
                                    borderWidth={0}
                                    justify="flex-start"
                                    {...(header.column.id === 'user_base'
                                        ? {
                                            flexShrink: 1,
                                        }
                                        : {
                                            flexShrink: 3,
                                        })}
                                >
                                    <Text>
                                        {header.isPlaceholder
                                            ? null
                                            : flexRender(header.column.columnDef.header, header.getContext())}
                                    </Text>
                                </Table.HeaderCell>
                            ))}
                        </Table.Row>
                    )
                })}
            </Table.Head>
            <Table.Body>
                {tableRows.map((row) => {
                    rowCounter.current++
                    return (
                        <Table.Row
                            onPress={() => router.navigate(`/resource/${row.original.resourceId}`)}
                            hoverStyle={{
                                bg: '$color2',
                            }}
                            rowLocation={
                                rowCounter.current === 0
                                    ? 'first'
                                    : rowCounter.current === allRowsLength - 1
                                        ? 'last'
                                        : 'middle'
                            }
                            key={row.id}
                        >
                            {row.getVisibleCells().map((cell) => (
                                <Table.Cell
                                    cellLocation={
                                        cell.column.id === 'fullName'
                                            ? 'first'
                                            : cell.column.id === 'role'
                                                ? 'last'
                                                : 'middle'
                                    }
                                    key={cell.id}
                                    borderWidth={0}
                                    justify="flex-start"
                                    {...(cell.column.id === 'user_base'
                                        ? {
                                            flexShrink: 1,
                                        }
                                        : {
                                            flexShrink: 3,
                                        })}
                                >
                                    {cell.column.id === 'user_base' ? (
                                        flexRender(cell.column.columnDef.cell, cell.getContext())
                                    ) : (
                                        <Text>
                                            {flexRender(cell.column.columnDef.cell, cell.getContext())}
                                        </Text>
                                    )}
                                </Table.Cell>
                            ))}
                        </Table.Row>
                    )
                })}
            </Table.Body>
        </Table>
    )
}