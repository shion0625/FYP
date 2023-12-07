'use client';
import { useForm } from 'react-hook-form';
import { toast } from 'react-hot-toast';

import { yupResolver } from '@hookform/resolvers/yup';
import { Label, TextInput, Button } from 'flowbite-react';
import { HiLockClosed, HiMail } from 'react-icons/hi';

import { UseLogin, LoginBody } from '@/actions/user';

import { loginSchema } from '@/schema/user';

const LoginView = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginBody>({
    mode: 'onBlur',
    defaultValues: {
      email: '',
      password: '',
    },
    resolver: yupResolver(loginSchema),
  });

  const { login } = UseLogin();

  const onSubmit = async (data: LoginBody) => {
    try {
      const response = await login({
        ...data,
      });
      toast.success(response.message);
    } catch (error: unknown) {
      toast.error('failed to login');
    }
  };

  return (
    <div className="flex justify-center">
      <form onSubmit={handleSubmit(onSubmit)} className="max-w-md w-2/4">
        <div className="block mb-3">
          <Label htmlFor="email" value="Email" />
          {errors.email && (
            <Label
              className="text-sm block"
              htmlFor="email"
              color="failure"
              value={errors.email.message}
            />
          )}
          <TextInput
            id="email"
            type="email"
            icon={HiMail}
            placeholder="name@flowbite.com"
            color={errors.email ? 'failure' : undefined}
            {...register('email', { required: true })}
          />
        </div>
        <div className="block mb-3">
          <Label htmlFor="password" value="Password" />
          {errors.password && (
            <Label
              className="text-sm block"
              htmlFor="password"
              color="failure"
              value={errors.password.message}
            />
          )}
          <TextInput
            id="password"
            type="password"
            icon={HiLockClosed}
            placeholder="password"
            color={errors.password ? 'failure' : undefined}
            {...register('password', { required: true })}
          />
        </div>

        <div className="flex justify-center mb-3">
          <Button type="submit">Login</Button>
        </div>
      </form>
    </div>
  );
};

export default LoginView;
