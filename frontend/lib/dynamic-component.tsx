import dynamic from "next/dynamic";
import LoadingSkeleton from "@/components/ui/loading-skeleton";

export function getDynamicComponent<P = {}>(
  c: string,
  count?: number,
  type?: "productCard" | "productCardItem"
) {
  return dynamic<P>(() => import(`@/components/${c}`), {
    ssr: false,
    loading: () => <LoadingSkeleton count={count} type={type} />,
  });
}
