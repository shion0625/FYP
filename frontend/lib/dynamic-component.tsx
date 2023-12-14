import dynamic from 'next/dynamic';
import LoadingSkeleton from '@/components/ui/loading-skeleton';

export function getDynamicComponent<P = object>(c: string, children: React.ReactNode) {
  return dynamic<P>(() => import(`@/app/(routes)/${c}/components/views`), {
    ssr: false,
    loading: () => <LoadingSkeleton>{children}</LoadingSkeleton>,
  });
}
