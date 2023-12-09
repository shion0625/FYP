'use client';
import { getDynamicComponent } from '@/lib/dynamic-component';
import Container from '@/components/ui/container';

export const revalidate = 0;

const DynamicLazyMyPage = getDynamicComponent('views/mypage', <></>);

const MyPage = () => (
  <div className="bg-white">
    <Container>
      <DynamicLazyMyPage />
    </Container>
  </div>
);

export default MyPage;
