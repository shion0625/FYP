import dynamic from "next/dynamic";
import LoadingSkeleton from "@/components/ui/loading-skeleton";

export function getDynamicComponent<P = {}>(c: string) {
  return dynamic<P>(() => import(`@/components/${c}`), {
    ssr: false,
    loading: () => <LoadingSkeleton />,
  });
}
