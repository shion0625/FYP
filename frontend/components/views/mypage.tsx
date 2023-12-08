'use client';
import useUserProfile from '@/hooks/use-user-profile';

const MyPageView = () => {
  const userProfile = useUserProfile();

  return <div className="space-y-10 pb-10"></div>;
};

export default MyPageView;
