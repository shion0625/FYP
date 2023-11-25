import { ProductVariationValue } from "@/types";

interface VariationProps {
  name: string;
  productVariationValues: ProductVariationValue[];
  selectedValue: ProductVariationValue | null;
  onSelect: (value: ProductVariationValue) => void;
}

const Variation: React.FC<VariationProps> = ({
  name,
  productVariationValues,
  selectedValue,
  onSelect,
}) => (
  <div>
    <h4>{name}</h4>
    <div className="flex flex-wrap">
      {productVariationValues.map((productVariationValue, i) => (
        <p
          key={i}
          className={`mr-2 p-2 rounded-full border ${
            selectedValue?.value === productVariationValue.value
              ? "bg-blue-500 text-white"
              : "bg-white"
          }`}
          onClick={() => onSelect(productVariationValue)}
        >
          {productVariationValue.value}
        </p>
      ))}
    </div>
  </div>
);

export default Variation;
