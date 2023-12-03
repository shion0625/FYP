"use client";
import { Label, TextInput, Button, Checkbox } from "flowbite-react";
import {
  HiUser,
  HiLockClosed,
  HiMail,
  HiPhone,
  HiIdentification,
} from "react-icons/hi";
import Link from "next/link";
import { UseSignIn } from "@/actions/user";
import { signInSchema } from "@/schema/user";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { toast } from "react-hot-toast";

const SignInView = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    mode: "onBlur", // バリデーションチェックのトリガー（フォーカスを外した時）
    defaultValues: {
      userName: "",
      firstName: "",
      lastName: "",
      age: 0,
      email: "",
      phone: "",
      password: "",
      confirmPassword: "",
      agree: false,
    },
    resolver: yupResolver(signInSchema),
  });

  const { signIn } = UseSignIn();

  const onSubmit = async (data: any) => {
    try {
      const response = await signIn({
        ...data,
        age: parseInt(data.age),
      });
      console.log(response);
    } catch (error: any) {
      toast.error(error.response.data.error);
    }
  };

  return (
    <div className="flex justify-center p-6">
      <form onSubmit={handleSubmit(onSubmit)} className="max-w-md">
        <div className="grid grid-cols-2 gap-4 mb-3">
          <div className="block col-span-1">
            <Label htmlFor="firstName" value="First Name" />
            {errors.firstName && (
              <Label
                className="text-sm block"
                htmlFor="firstName"
                color="failure"
                value={errors.firstName.message}
              />
            )}
            <TextInput
              id="firstName"
              type="text"
              icon={HiUser}
              placeholder="Kwame"
              color={errors.firstName ? "failure" : undefined}
              {...register("firstName", { required: true })}
            />
          </div>
          <div className=" block col-span-1">
            <Label htmlFor="lastName" value="Last Name" />
            {errors.lastName && (
              <Label
                className="text-sm block"
                htmlFor="lastName"
                color="failure"
                value={errors.lastName.message}
              />
            )}
            <TextInput
              id="lastName"
              type="text"
              icon={HiUser}
              placeholder="Nkrumah"
              color={errors.lastName ? "failure" : undefined}
              {...register("lastName", { required: true })}
            />
          </div>
        </div>
        <div className="grid grid-cols-4 gap-4 mb-3">
          <div className="block col-span-3">
            <Label htmlFor="userName" value="Username" />
            {errors.userName && (
              <Label
                className="text-sm block"
                htmlFor="userName"
                color="failure"
                value={errors.userName.message}
              />
            )}
            <TextInput
              id="userName"
              type="text"
              icon={HiUser}
              placeholder="Username"
              color={errors.userName ? "failure" : undefined}
              {...register("userName", { required: true })}
            />
          </div>
          <div className="block col-span-1">
            <Label htmlFor="age" value="Age" />
            {errors.age && (
              <Label
                className="text-sm block"
                htmlFor="age"
                color="failure"
                value={errors.age.message}
              />
            )}
            <TextInput
              id="age"
              type="number"
              icon={HiIdentification}
              placeholder="20"
              color={errors.age ? "failure" : undefined}
              {...register("age", { required: true })}
            />
          </div>
        </div>
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
            color={errors.email ? "failure" : undefined}
            {...register("email", { required: true })}
          />
        </div>
        <div className="block mb-3">
          <Label htmlFor="phone" value="Phone" />
          {errors.phone && (
            <Label
              className="text-sm block"
              htmlFor="phone"
              color="failure"
              value={errors.phone.message}
            />
          )}
          <TextInput
            id="phone"
            type="tel"
            icon={HiPhone}
            placeholder="07012345678"
            color={errors.phone ? "failure" : undefined}
            {...register("phone", { required: true })}
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
            color={errors.password ? "failure" : undefined}
            {...register("password", { required: true })}
          />
        </div>
        <div className="block mb-3">
          <Label htmlFor="confirmPassword" value="Confirm Password" />
          {errors.confirmPassword && (
            <Label
              className="text-sm block"
              htmlFor="confirmPassword"
              color="failure"
              value={errors.confirmPassword.message}
            />
          )}
          <TextInput
            id="confirmPassword"
            type="password"
            icon={HiLockClosed}
            placeholder="confirm Password"
            color={errors.confirmPassword ? "failure" : undefined}
            {...register("confirmPassword", { required: true })}
          />
        </div>
        <div className="block mb-3">
          {errors.agree && (
            <Label
              className="text-sm block"
              htmlFor="agree"
              color="failure"
              value={errors.agree.message}
            />
          )}
          <div className="flex items-center gap-2 mb-3">
            <Checkbox id="agree" {...register("agree", { required: true })} />
            <Label htmlFor="agree" className="flex">
              I agree with the&nbsp;
              <Link
                href="#"
                className="text-cyan-600 hover:underline dark:text-cyan-500"
              >
                terms and conditions
              </Link>
            </Label>
          </div>
        </div>
        <div className="flex justify-center mb-3">
          <Button type="submit">Register new account</Button>
        </div>
      </form>
    </div>
  );
};

export default SignInView;
