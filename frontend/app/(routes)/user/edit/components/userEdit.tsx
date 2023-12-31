'use client';
import React, { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'react-hot-toast';
import { yupResolver } from '@hookform/resolvers/yup';
import { Label, TextInput, Button } from 'flowbite-react';
import { UseGetProfile, ProfileBody, ProfileSchema } from '@/actions/user/user-profile';
import { CardTitle, CardDescription, CardHeader, CardContent, Card } from '@/components/ui/card';

const AddressView = () => {
  const { getUserProfile, updateUserProfile } = UseGetProfile();

  const [userProfile, setUserProfile] = useState<ProfileBody>();

  useEffect(() => {
    getUserProfile().then((profile) => {
      setUserProfile(profile);
      reset(profile);
    });
  }, []);

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<ProfileBody>({
    mode: 'onBlur',
    defaultValues: { ...userProfile },
    resolver: yupResolver(ProfileSchema),
  });

  const onSubmit = async (data: ProfileBody) => {
    try {
      await updateUserProfile(data);
      toast.success('success to edit profile');
    } catch (error: unknown) {
      toast.error('failed to edit profile');
    }
  };

  return (
    <div className="flex justify-center p-6">
      <Card className="space-y-8">
        <CardHeader>
          <CardTitle className="text-3xl font-bold">User Information</CardTitle>
          <CardDescription className="text-gray-500 dark:text-gray-400">
            Please fill out the form below with your shipping details.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit(onSubmit)}>
            <Label htmlFor="userName" value="Username" />
            <span className="text-red-600"> *</span>
            <TextInput type="text" {...register('userName', { required: true })} />
            {errors.userName && (
              <Label
                className="text-sm block"
                htmlFor="userName"
                color="failure"
                value={errors.userName.message}
              />
            )}
            <Label htmlFor="firstName" value="First Name" />
            <span className="text-red-600"> *</span>
            <TextInput type="text" {...register('firstName', { required: true })} />
            {errors.firstName && (
              <Label
                className="text-sm block"
                htmlFor="firstName"
                color="failure"
                value={errors.firstName.message}
              />
            )}
            <Label htmlFor="lastName" value="Last Name" />
            <span className="text-red-600"> *</span>
            <TextInput type="text" {...register('lastName', { required: true })} />
            {errors.lastName && (
              <Label
                className="text-sm block"
                htmlFor="lastName"
                color="failure"
                value={errors.lastName.message}
              />
            )}
            <Label htmlFor="age" value="Age" />
            <span className="text-red-600"> *</span>
            <TextInput type="number" {...register('age', { required: true })} />
            {errors.age && (
              <Label
                className="text-sm block"
                htmlFor="age"
                color="failure"
                value={errors.age.message}
              />
            )}
            <Label htmlFor="email" value="Email" />
            <span className="text-red-600"> *</span>
            <TextInput type="email" {...register('email', { required: true })} />
            {errors.email && (
              <Label
                className="text-sm block"
                htmlFor="email"
                color="failure"
                value={errors.email.message}
              />
            )}
            <Label htmlFor="phone" value="Phone" />
            <span className="text-red-600"> *</span>
            <TextInput type="tel" {...register('phone', { required: true })} />
            {errors.phone && (
              <Label
                className="text-sm block"
                htmlFor="phone"
                color="failure"
                value={errors.phone.message}
              />
            )}
            <div className="flex justify-center mb-3">
              <Button type="submit">edit profile</Button>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
};

export default AddressView;
