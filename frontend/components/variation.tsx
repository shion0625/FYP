interface VariationProps {
  name: string;
  values: string[];
  selectedValue: string | null;
  onSelect: (value: string) => void;
}

const Variation: React.FC<VariationProps> = ({
  name,
  values,
  selectedValue,
  onSelect,
}) => (
  <div>
    <h4>{name}</h4>
    <div className="flex flex-wrap">
      {values.map((value, i) => (
        <p
          key={i}
          className={`mr-2 p-2 rounded-full border ${
            selectedValue === value ? "bg-blue-500 text-white" : "bg-white"
          }`}
          onClick={() => onSelect(value)}
        >
          {value}
        </p>
      ))}
    </div>
  </div>
);

export default Variation;
