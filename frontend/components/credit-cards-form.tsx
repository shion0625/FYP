import React from 'react';
import Cards from 'react-credit-cards-2';
import { useForm, Controller } from 'react-hook-form';
import { toast } from 'react-hot-toast';
import 'react-credit-cards-2/dist/es/styles-compiled.css';
import { yupResolver } from '@hookform/resolvers/yup';
import { Label, TextInput } from 'flowbite-react';
import { HiCreditCard, HiUser, HiCalendar, HiLockClosed } from 'react-icons/hi';
import {
  UsePaymentMethod,
  PaymentMethodBody,
  CreditCardSchema,
} from '@/actions/user/payment-method';

interface CreditCardsFormProps {
  setIsSubmitted: React.Dispatch<React.SetStateAction<boolean>>;
}

const CreditCardsForm: React.FC<CreditCardsFormProps> = ({ setIsSubmitted }) => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
    control,
  } = useForm<PaymentMethodBody>({
    mode: 'onChange',
    defaultValues: {
      number: '',
      name: '',
      expiry: '',
      cvc: '',
    },
    resolver: yupResolver(CreditCardSchema),
  });
  const { savePaymentMethod } = UsePaymentMethod();

  const formatExpiryDate = (value: string) => {
    return value
      .replace(/\W/gi, '')
      .replace(/(.{2})/, '$1/')
      .slice(0, 5);
  };

  const onSubmit = async (data: PaymentMethodBody) => {
    try {
      await savePaymentMethod(data);
      toast.success('success to add paymentMethod');
      setIsSubmitted(true);
    } catch (error: unknown) {
      toast.error('failed to add paymentMethod');
    }
  };
  const watchAllFields = watch();

  return (
    <div className="grid grid-cols-1 gap-4">
      <Cards
        number={watchAllFields.number}
        expiry={watchAllFields.expiry}
        cvc={watchAllFields.cvc}
        name={watchAllFields.name}
      />
      <form onSubmit={handleSubmit(onSubmit)} className="mx-4">
        <TextInput
          type="number"
          className="form-control mt-2"
          placeholder="Card Number"
          icon={HiCreditCard}
          {...register('number', {
            required: true,
            maxLength: 16,
          })}
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
          className="form-control mt-2"
          placeholder="Name"
          icon={HiUser}
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
        <div className="grid grid-cols-2 gap-4">
          <div>
            <Controller
              name="expiry"
              control={control}
              defaultValue=""
              rules={{ required: true, maxLength: 5 }}
              render={({ field }) => (
                <TextInput
                  type="tel"
                  className="form-control mt-2"
                  placeholder="MM/YY"
                  icon={HiCalendar}
                  value={formatExpiryDate(field.value)}
                  onChange={(e) => field.onChange(e.target.value)}
                />
              )}
            />
            {errors.expiry && (
              <Label
                className="text-sm block"
                htmlFor="expiry"
                color="failure"
                value={errors.expiry.message}
              />
            )}
          </div>
          <div>
            <TextInput
              type="tel"
              className="form-control mt-2"
              placeholder="CVC"
              pattern="\d{3,4}"
              icon={HiLockClosed}
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
          </div>
        </div>
        <TextInput type="submit" className="btn btn-primary mt-2" />
      </form>
    </div>
  );
};

export default CreditCardsForm;
