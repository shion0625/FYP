"use client";

import { useEffect, useState } from "react";

const formatter = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
});

interface CurrencyProps {
  value?: string | number;
  discountPrice?: string | number;
}

const Currency: React.FC<CurrencyProps> = ({ value, discountPrice }) => {
  return (
    <div>
      {!discountPrice ? (
        <div className="font-semibold text-stone-700">
          {formatter.format(Number(value))}
        </div>
      ) : (
        <div style={{ position: "relative" }}>
          <div
            className="font-semibold"
            style={{
              position: "absolute",
              top: -15,
              right: -10,
              color: "gray",
              textDecoration: "line-through",
              fontSize: "80%",
            }}
          >
            {formatter.format(Number(value))}
          </div>
          <div className="font-semibold text-stone-700">
            {formatter.format(Number(discountPrice))}
          </div>
        </div>
      )}
    </div>
  );
};

export default Currency;
