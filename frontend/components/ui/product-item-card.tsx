// frontend/components/product-item-card.tsx
import { ProductItem } from "@/types";
import Currency from "@/components/ui/currency";
import NextImage from "next/image";

interface ProductItemCardProps {
  data: ProductItem;
}

const ProductItemCard: React.FC<ProductItemCardProps> = ({ data }) => {
  return (
    <div className="border p-4 rounded-md">
      <h2 className="text-xl font-bold">{data.name}</h2>
      <p className="text-gray-500">{data.sku}</p>
      <p className="mt-2">
        <Currency value={data.price} />
        {data.discountPrice !== data.price && (
          <>
            <span className="line-through text-gray-500 ml-2">
              <Currency value={data.discountPrice} />
            </span>
          </>
        )}
      </p>
      <p className="mt-2">Stock: {data.qtyInStock}</p>
      <p className="mt-2">Category: {data.categoryName}</p>
      <p className="mt-2">Brand: {data.brandName}</p>
      {data.images.length > 0 && (
        <NextImage src={data.images[0]} alt={data.name} className="mt-4" />
      )}
    </div>
  );
};

export default ProductItemCard;
