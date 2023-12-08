'use client';
import Container from '@/components/ui/container';
import MyPageView from '@/components/views/mypage';

export const revalidate = 0;

const MyPage = () => (
  <Container>
    <MyPageView />
  </Container>
);

export default MyPage;
