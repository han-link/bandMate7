import { ChevronLeft, ChevronRight } from "@tamagui/lucide-icons-2";
import { useCallback, useEffect, useState } from "react";
import { useWindowDimensions } from "react-native";
import "react-pdf/dist/Page/TextLayer.css";
import "react-pdf/dist/Page/AnnotationLayer.css";
import { Button, SizableText, Theme, View, XStack, styled } from "tamagui";
import { PdfViewProps } from "./PdfView";

type ReactPdf = typeof import("react-pdf");

const TapZone = styled(View, {
  name: "PdfTapZone",
  position: "absolute",
  t: 0,
  b: 0,
  width: "33%",
  variants: {
    side: {
      left: { l: 0 },
      right: { r: 0 },
    },
  } as const,
});

const PagePill = styled(SizableText, {
  name: "PagePill",
  color: "$color",
  bg: "$shadow5",
  px: "$2.5",
  py: "$1.5",
  rounded: "$4",
  overflow: "hidden",
});

export default function PdfView(props: PdfViewProps) {
  const [reactPdf, setReactPdf] = useState<ReactPdf>();
  const [pageNumber, setPageNumber] = useState(1);
  const [numPages, setNumPages] = useState<number>();
  const { height } = useWindowDimensions();

  useEffect(() => {
    let active = true;
    import("react-pdf").then((mod) => {
      if (!active) return;
      mod.pdfjs.GlobalWorkerOptions.workerSrc = `//unpkg.com/pdfjs-dist@${mod.pdfjs.version}/build/pdf.worker.min.mjs`;
      setReactPdf(mod);
    });
    return () => {
      active = false;
    };
  }, []);

  const goToPage = useCallback(
    (delta: number) =>
      setPageNumber((p) => Math.min(Math.max(p + delta, 1), numPages ?? p)),
    [numPages],
  );

  const STEP = 2;

  useEffect(() => {
    function onKeyDown(e: KeyboardEvent) {
      if (e.key === "ArrowRight" || e.key === "ArrowDown") goToPage(STEP);
      else if (e.key === "ArrowLeft" || e.key === "ArrowUp") goToPage(-STEP);
    }
    window.addEventListener("keydown", onKeyDown);
    return () => window.removeEventListener("keydown", onKeyDown);
  }, [goToPage]);

  function onDocumentLoadSuccess({ numPages }: { numPages: number }): void {
    setNumPages(numPages);
  }

  if (!props.resource.url || !reactPdf) return null;

  const { Document, Page } = reactPdf;
  const hasSecondPage = !!numPages && pageNumber + 1 <= numPages;
  const canPrev = pageNumber > 1;
  const canNext = !!numPages && pageNumber + STEP <= numPages;

  return (
    <View flex={1}>
      <Document file={props.resource.url} onLoadSuccess={onDocumentLoadSuccess}>
        <XStack gap="$2">
          <Page pageNumber={pageNumber} height={height} />
          {hasSecondPage && (
            <Page pageNumber={pageNumber + 1} height={height} />
          )}
        </XStack>
      </Document>

      <TapZone
        side="left"
        onPress={() => goToPage(-STEP)}
        disabled={!canPrev}
      />
      <TapZone
        side="right"
        onPress={() => goToPage(STEP)}
        disabled={!canNext}
      />

      {/* Pinned dark so the scrim and text stay readable over any page. */}
      <Theme name="dark">
        <XStack position="absolute" t="$3.5" r="$3.5" items="center" gap="$2">
          <Button
            onPress={() => goToPage(-STEP)}
            disabled={!canPrev}
            bg="$shadow5"
          >
            <ChevronLeft />
          </Button>
          <PagePill>
            {hasSecondPage ? `${pageNumber}–${pageNumber + 1}` : pageNumber}
            {numPages ? ` / ${numPages}` : ""}
          </PagePill>
          <Button
            onPress={() => goToPage(STEP)}
            disabled={!canNext}
            bg="$shadow5"
          >
            <ChevronRight />
          </Button>
        </XStack>
      </Theme>
    </View>
  );
}
