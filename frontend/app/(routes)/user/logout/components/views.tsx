'use client';

import { useEffect, memo } from 'react';
import { toast } from 'react-hot-toast';
import { useRouter } from 'next/navigation';
import { useLogout } from '@/actions/user/logout';
import useLoginState from '@/hooks/use-login';

const LogoutView = () => {
  const router = useRouter();
  const { logout } = useLogout();
  const loginState = useLoginState();

  useEffect(() => {
    try {
      logout();
      loginState.onLogout();
      toast.success('successful logout');
    } catch (error: unknown) {
      toast.error('failed to logout');
    }
    router.push('/');
  }, []);

  return <></>;
};

export default memo(LogoutView);
