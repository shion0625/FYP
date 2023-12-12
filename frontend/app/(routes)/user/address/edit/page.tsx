'use client';
import Container from '@/components/ui/container';
import AddressView from '@/components/views/addressEdit';
export const revalidate = 0;

const AddressPage = () => (
  <Container>
    <AddressView />
  </Container>
);

export default AddressPage;
