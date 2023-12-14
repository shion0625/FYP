'use client';
import Container from '@/components/ui/container';
import AddressView from '@/app/(routes)/user/address/edit/components/addressEdit';
export const revalidate = 0;

const AddressPage = () => (
  <Container>
    <AddressView />
  </Container>
);

export default AddressPage;
