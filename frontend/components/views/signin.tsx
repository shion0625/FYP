"use client";
import { Label, TextInput, Button, Checkbox } from "flowbite-react";
import {
  HiUser,
  HiLockClosed,
  HiMail,
  HiPhone,
  HiIdentification,
} from "react-icons/hi";
import { useRef } from "react";
import Link from "next/link";
import { UseSignIn } from "@/actions/user";

const SignInView = () => {
  const signInRef = useRef({
    userName: "",
    firstName: "",
    lastName: "",
    age: "",
    email: "",
    phone: "",
    password: "",
    confirmPassword: "",
  });

  const { signIn } = UseSignIn();

  const handleSignIn = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const values = signInRef.current;
    const response = await signIn({
      ...values,
      age: parseInt(values.age),
    });
    console.log(response);
  };

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const name = event.target.id as keyof typeof signInRef.current;
    signInRef.current[name] = event.target.value;
  };

  return (
    <div className="flex justify-center p-6">
      <form onSubmit={handleSignIn} className="max-w-md">
        <div className="grid grid-cols-2 gap-4 mb-3">
          <div className="block col-span-1">
            <Label htmlFor="firstName" value="First Name" />
            <TextInput
              id="firstName"
              type="text"
              icon={HiUser}
              placeholder="Kwame"
              required
              onChange={handleChange}
            />
          </div>
          <div className=" block col-span-1">
            <Label htmlFor="lastName" value="Last Name" />
            <TextInput
              id="lastName"
              type="text"
              icon={HiUser}
              placeholder="Nkrumah"
              required
              onChange={handleChange}
            />
          </div>
        </div>
        <div className="grid grid-cols-4 gap-4 mb-3">
          <div className="block col-span-3">
            <Label htmlFor="userName" value="Username" />
            <TextInput
              id="userName"
              type="text"
              icon={HiUser}
              placeholder="Username"
              required
              onChange={handleChange}
            />
          </div>
          <div className="block col-span-1">
            <Label htmlFor="age" value="Age" />
            <TextInput
              id="age"
              type="number"
              icon={HiIdentification}
              placeholder="20"
              required
              onChange={handleChange}
            />
          </div>
        </div>
        <div className="block mb-3">
          <Label htmlFor="email" value="Email" />
          <TextInput
            id="email"
            type="email"
            icon={HiMail}
            placeholder="name@flowbite.com"
            required
            onChange={handleChange}
          />
        </div>
        <div className="block mb-3">
          <Label htmlFor="phone" value="Phone" />
          <TextInput
            id="phone"
            type="tel"
            icon={HiPhone}
            placeholder="07012345678"
            required
            onChange={handleChange}
          />
        </div>
        <div className="block mb-3">
          <Label htmlFor="password" value="Password" />
          <TextInput
            id="password"
            type="password"
            icon={HiLockClosed}
            placeholder="password"
            required
            onChange={handleChange}
          />
        </div>
        <div className="block mb-3">
          <Label htmlFor="confirmPassword" value="Confirm Password" />
          <TextInput
            id="confirmPassword"
            type="password"
            icon={HiLockClosed}
            placeholder="confirm Password"
            required
            onChange={handleChange}
          />
        </div>
        <div className="flex items-center gap-2 mb-3">
          <Checkbox id="agree" />
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
        <div className="flex justify-center mb-3">
          <Button type="submit">Register new account</Button>
        </div>
      </form>
    </div>
  );
};

export default SignInView;
