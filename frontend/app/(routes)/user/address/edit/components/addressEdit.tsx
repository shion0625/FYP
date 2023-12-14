'use client';
import { useForm } from 'react-hook-form';
import { toast } from 'react-hot-toast';
import { yupResolver } from '@hookform/resolvers/yup';
import { Label, TextInput, Button } from 'flowbite-react';
import {
  UseUserAddresses,
  AddressBody,
  UpdateAddressBody,
  UpdateAddressSchema,
} from '@/actions/user/user-address';
import { CardTitle, CardDescription, CardHeader, CardContent, Card } from '@/components/ui/card';

const AddressView = () => {
  const { saveUserAddress } = UseUserAddresses();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<UpdateAddressBody>({
    mode: 'onBlur',
    defaultValues: {},
    resolver: yupResolver(UpdateAddressSchema),
  });

  const onSubmit = async (data: AddressBody) => {
    try {
      const response = await saveUserAddress(data);
      console.log('front', response);
      toast.success('success to add address');
    } catch (error: unknown) {
      toast.error('failed to add address');
    }
  };

  return (
    <div className="flex justify-center p-6">
      <Card className="space-y-8">
        <CardHeader>
          <CardTitle className="text-3xl font-bold">Shipping Information</CardTitle>
          <CardDescription className="text-gray-500 dark:text-gray-400">
            Please fill out the form below with your shipping details.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit(onSubmit)}>
            <Label htmlFor="name" value="Address Name" />
            <span className="text-red-600"> *</span>
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
            <span className="text-red-600"> *</span>
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
            <span className="text-red-600"> *</span>
            <TextInput type="text" {...register('city', { required: true })} />
            {errors.city && (
              <Label
                className="text-sm block"
                htmlFor="city"
                color="failure"
                value={errors.city.message}
              />
            )}
            <div className="grid grid-cols-2 gap-4">
              <div>
                <Label htmlFor="area" value="Area/State" />
                <span className="text-red-600"> *</span>
                <TextInput type="text" {...register('area', { required: true })} />
                {errors.area && (
                  <Label
                    className="text-sm block"
                    htmlFor="area"
                    color="failure"
                    value={errors.area.message}
                  />
                )}
              </div>
              <div>
                <Label htmlFor="pincode" value="Pincode" />
                <span className="text-red-600"> *</span>
                <TextInput type="text" {...register('pincode', { required: true })} />
                {errors.pincode && (
                  <Label
                    className="text-sm block"
                    htmlFor="pincode"
                    color="failure"
                    value={errors.pincode.message}
                  />
                )}
              </div>
            </div>
            <Label htmlFor="countryName" value="Country" />
            <span className="text-red-600"> *</span>
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
            <span className="text-red-600"> *</span>
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
            <span className="text-red-600"> *</span>
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
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
};

export default AddressView;
