import * as yup from 'yup';

export const signUpSchema = yup.object().shape({
  userName: yup
    .string()
    .required('Username is required')
    .matches(/^[a-zA-Z0-9]*$/, 'Only alphanumeric characters are allowed for username')
    .min(3, 'Username must be at least 3 characters')
    .max(15, "Username can't be longer than 15 characters"),
  firstName: yup
    .string()
    .required('First name is required')
    .matches(/^[a-zA-Z]*$/, 'Only alphabets are allowed for first name')
    .min(2, 'First name must be at least 2 characters')
    .max(50, "First name can't be longer than 50 characters"),
  lastName: yup
    .string()
    .required('Last name is required')
    .matches(/^[a-zA-Z]*$/, 'Only alphabets are allowed for last name')
    .min(1, 'Last name must be at least 1 character')
    .max(50, "Last name can't be longer than 50 characters"),
  age: yup
    .number()
    .required('Age is required')
    .min(0, "Age can't be less than 0")
    .max(120, "Age can't be more than 120")
    .transform((value, originalValue) => {
      return originalValue === '' ? undefined : value;
    }),
  email: yup.string().required('Email is required').email('Must be a valid email address'),
  phone: yup
    .string()
    .required('Phone number is required')
    .matches(/^\+[1-9]\d{1,14}$/, 'Must be a valid E.164 format for phone number'),
  password: yup
    .string()
    .required('Password is required')
    .min(5, 'Password must be at least 5 characters')
    .max(30, "Password can't be longer than 30 characters"),
  confirmPassword: yup
    .string()
    .required('Confirm password is required')
    .oneOf([yup.ref('password')], 'Passwords must match'),
  agree: yup
    .boolean()
    .required('terms and conditions is required')
    .oneOf([true], 'You must agree to the terms and conditions'),
});

export const loginSchema = yup.object().shape({
  email: yup.string().required('Email is required').email('Must be a valid email address'),
  password: yup
    .string()
    .required('Password is required')
    .min(5, 'Password must be at least 5 characters')
    .max(30, "Password can't be longer than 30 characters"),
});

export const creditCardSchema = yup.object().shape({
  number: yup
    .string()
    .required('Card number is required')
    .matches(/^[0-9]{13,16}$/, 'Card number must be 13 to 16 digits'),
  name: yup
    .string()
    .required('Name is required')
    .matches(/^[a-zA-Z ]*$/, 'Only alphabets are allowed for name')
    .min(2, 'Name must be at least 2 characters')
    .max(50, "Name can't be longer than 50 characters"),
  expiry: yup
    .string()
    .required('Expiry date is required')
    .matches(
      /^(0[1-9]|1[0-2])\/([0-9]{4}|[0-9]{2})$/,
      'Must be a valid MM/YY format'
    ),
  cvc: yup
    .string()
    .required('CVC is required')
    .matches(/^[0-9]{3,4}$/, 'Must be a valid CVC number'),
});

export const addressSchema = yup.object().shape({
  name: yup
    .string()
    .required('Name is required')
    .min(2, 'Name must be at least 2 characters')
    .max(100, "Name can't be longer than 100 characters"),
  phoneNumber: yup
    .string()
    .required('Phone number is required')
    .matches(/^\+[1-9]\d{1,14}$/, 'Must be a valid E.164 format for phone number'),
  house: yup.string().required('House is required'),
  area: yup.string().required('area is required'),
  landMark: yup.string().required('Landmark is required'),
  city: yup.string().required('city is required'),
  pincode: yup.string().required('Pincode is required'),
  countryName: yup.string().required('Country name is required'),
});
