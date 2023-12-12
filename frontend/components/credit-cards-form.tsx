import React from 'react';
import Cards from 'react-credit-cards-2';
import 'react-credit-cards-2/dist/es/styles-compiled.css';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { Label, TextInput } from 'flowbite-react';
import { creditCardSchema } from '@/schema/user';

const CreditCardsForm = () => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm({
    mode: 'onSubmit',
    defaultValues: {
      number: '',
      name: '',
      expiry: '',
      cvc: '',
    },
    resolver: yupResolver(creditCardSchema),
  });
  const onSubmit = (data: unknown) => console.log(data);
  const watchAllFields = watch();

  return (
    <div className="flex justify-center">
      <Cards
        number={watchAllFields.number}
        expiry={watchAllFields.expiry}
        cvc={watchAllFields.cvc}
        name={watchAllFields.name}
      />
      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col items-center">
        <TextInput
          type="number"
          placeholder="Card Number"
          {...register('number', { required: true })}
          className="form-control"
        />
        {errors.number && (
          <Label
            className="text-sm block"
            htmlFor="number"
            color="failure"
            value={errors.number.message}
          />
        )}
        <TextInput
          type="text"
          className="form-control"
          placeholder="Name"
          {...register('name', { required: true })}
        />
        {errors.name && (
          <Label
            className="text-sm block"
            htmlFor="name"
            color="failure"
            value={errors.name.message}
          />
        )}
        <TextInput
          type="tel"
          className="form-control"
          placeholder="MMYY"
          pattern="\d\d\d\d"
          {...register('expiry', { required: true })}
        />
        {errors.expiry && (
          <Label
            className="text-sm block"
            htmlFor="expiry"
            color="failure"
            value={errors.expiry.message}
          />
        )}
        <TextInput
          type="tel"
          className="form-control"
          placeholder="CVC"
          pattern="\d{3,4}"
          {...register('cvc', { required: true })}
        />
        {errors.cvc && (
          <Label
            className="text-sm block"
            htmlFor="cvc"
            color="failure"
            value={errors.cvc.message}
          />
        )}
        <TextInput type="submit" className="btn btn-primary" />
      </form>
    </div>
  );
};

export default CreditCardsForm;
