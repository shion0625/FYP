'use client';

import { toast } from 'react-hot-toast';
import { useRouter } from 'next/navigation';
import { useLogout } from '@/actions/user/logout';

const LogoutView = () => {
  const router = useRouter();
  const { logout } = useLogout();
  try {
    logout();
    toast.success('successful logout');
  } catch (error: unknown) {
    toast.error('failed to logout');
  }
  router.push('/');
  return <></>;
};

export default LogoutView;
