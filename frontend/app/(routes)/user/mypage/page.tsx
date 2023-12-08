'use client';
import { getDynamicComponent } from '@/lib/dynamic-component';
import Container from '@/components/ui/container';

export const revalidate = 0;

const DynamicLazyMyPage = getDynamicComponent('views/mypage', <></>);

const MyPage = () => (
  <Container>
    <DynamicLazyMyPage />
  </Container>
);

export default MyPage;
