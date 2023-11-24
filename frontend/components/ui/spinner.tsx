"use client";

import { Spinner as SpinnerFlow } from "flowbite-react";

const Spinner = () => {
  return (
    <div className="flex flex-wrap items-center gap-2">
      <SpinnerFlow aria-label="Extra large spinner example" size="xl" />
    </div>
  );
};

export default Spinner;
