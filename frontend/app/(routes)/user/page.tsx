'use client';
import { getDynamicComponent } from '@/lib/dynamic-component';
import Container from '@/components/ui/container';

export const revalidate = 0;

const DynamicLazyUser = getDynamicComponent('views/user', <></>);

const MyUser = () => (
  <div className="bg-white">
    <Container>
      <DynamicLazyUser />
    </Container>
  </div>
);

export default MyUser;
