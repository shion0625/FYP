import dynamic from 'next/dynamic';

import LoadingSkeleton from '@/components/ui/loading-skeleton';

export function getDynamicComponent<P = object>(c: string, children: React.ReactNode) {
  return dynamic<P>(() => import(`@/components/${c}`), {
    ssr: false,
    loading: () => <LoadingSkeleton>{children}</LoadingSkeleton>,
  });
}
