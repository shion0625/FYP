'use client';
import UserEditView from '@/app/(routes)/user/edit/components/userEdit';
import Container from '@/components/ui/container';
export const revalidate = 0;

const AddressPage = () => (
  <Container>
    <UserEditView />
  </Container>
);

export default AddressPage;
