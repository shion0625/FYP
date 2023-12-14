'use client';
import AddressView from '@/app/(routes)/user/address/edit/components/addressEdit';
import Container from '@/components/ui/container';
export const revalidate = 0;

const AddressPage = () => (
  <Container>
    <AddressView />
  </Container>
);

export default AddressPage;
