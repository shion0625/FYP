'use client';

import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { Label, TextInput, Button } from 'flowbite-react';
import { addressSchema } from '@/schema/user';

const AddressView = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    mode: 'onBlur',
    defaultValues: {
      name: '',
      area: '',
      city: '',
      countryName: '',
      house: '',
      landMark: '',
      phoneNumber: '',
      pincode: '',
    },
    resolver: yupResolver(addressSchema),
  });

  const onSubmit = (data: unknown) => console.log(data);

  return (
    <div className="flex justify-center p-6">
      <form onSubmit={handleSubmit(onSubmit)}>
        <Label htmlFor="name" value="Address Name" />
        <TextInput type="text" {...register('name', { required: true })} />
        {errors.name && (
          <Label
            className="text-sm block"
            htmlFor="name"
            color="failure"
            value={errors.name.message}
          />
        )}
        <Label htmlFor="house" value="Address" />
        <TextInput type="text" {...register('house', { required: true })} />
        {errors.house && (
          <Label
            className="text-sm block"
            htmlFor="house"
            color="failure"
            value={errors.house.message}
          />
        )}
        <Label htmlFor="city" value="City" />
        <TextInput type="text" {...register('city', { required: true })} />
        {errors.city && (
          <Label
            className="text-sm block"
            htmlFor="city"
            color="failure"
            value={errors.city.message}
          />
        )}
        <Label htmlFor="area" value="Area/State" />
        <TextInput type="text" {...register('area', { required: true })} />
        {errors.area && (
          <Label
            className="text-sm block"
            htmlFor="area"
            color="failure"
            value={errors.area.message}
          />
        )}
        <Label htmlFor="pincode" value="Pincode" />
        <TextInput type="text" {...register('pincode', { required: true })} />
        {errors.pincode && (
          <Label
            className="text-sm block"
            htmlFor="pincode"
            color="failure"
            value={errors.pincode.message}
          />
        )}
        <Label htmlFor="countryName" value="Country" />
        <TextInput type="text" {...register('countryName', { required: true })} />
        {errors.countryName && (
          <Label
            className="text-sm block"
            htmlFor="countryName"
            color="failure"
            value={errors.countryName.message}
          />
        )}
        <Label htmlFor="landMark" value="Landmark" />
        <TextInput type="text" {...register('landMark', { required: true })} />
        {errors.landMark && (
          <Label
            className="text-sm block"
            htmlFor="landMark"
            color="failure"
            value={errors.landMark.message}
          />
        )}
        <Label htmlFor="phoneNumber" value="Phone Number" />
        <TextInput type="tel" {...register('phoneNumber', { required: true })} />
        {errors.phoneNumber && (
          <Label
            className="text-sm block"
            htmlFor="phoneNumber"
            color="failure"
            value={errors.phoneNumber.message}
          />
        )}
        <div className="flex justify-center mb-3">
          <Button type="submit">add address</Button>
        </div>{' '}
      </form>
    </div>
  );
};

export default AddressView;
