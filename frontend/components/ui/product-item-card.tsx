// frontend/components/product-item-card.tsx
import { ProductItem } from "@/types";
import Currency from "@/components/ui/currency";
import Gallery from "@/components/gallery";

interface ProductItemCardProps {
  data: ProductItem;
}

const ProductItemCard: React.FC<ProductItemCardProps> = ({ data }) => {
  return (
    <div>
      <div className="aspect-square rounded-xl bg-gray-100 relative mb-4">
        <Gallery id={data.sku} urls={data.images} />
      </div>
      {/* Price */}
      <div className="text-gray-500 ml-2">
        <Currency value={data?.price} discountPrice={data.discountPrice} />
      </div>
      {/* Description */}
      <div>
        <p className="text-sm text-gray-500">在庫: {data.qtyInStock}</p>
      </div>
    </div>
  );
};

export default ProductItemCard;
