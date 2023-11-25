// frontend/components/product-item-card.tsx
import { ProductItem } from "@/types";
import Currency from "@/components/ui/currency";
import Gallery from "@/components/gallery";
import usePreviewModal from "@/hooks/use-preview-modal";
import useCart from "@/hooks/use-cart";
import IconButton from "@/components/ui/icon-button";
import { Expand, ShoppingCart } from "lucide-react";
import { MouseEventHandler } from "react";

interface ProductItemCardProps {
  data: ProductItem;
}

const ProductItemCard: React.FC<ProductItemCardProps> = ({ data }) => {
  const cart = useCart();
  const previewModal = usePreviewModal();
  const onPreview: MouseEventHandler<HTMLButtonElement> = (event) => {
    event.stopPropagation();

    previewModal.onOpen(data);
  };

  const onAddToCart: MouseEventHandler<HTMLButtonElement> = (event) => {
    event.stopPropagation();

    cart.addItem(data);
  };
  return (
    <div className="bg-white group cursor-pointer rounded-xl border p-3 space-y-4">
      <div className="aspect-square rounded-xl bg-gray-100 relative">
        <div className="aspect-square rounded-xl bg-gray-100 relative mb-4">
          <Gallery id={data.sku} urls={data.images} />
        </div>
        <div className="opacity-0 group-hover:opacity-100 transition absolute w-full px-6 bottom-5">
          <div className="flex gap-x-6 justify-center">
            <IconButton
              onClick={onPreview}
              icon={<Expand size={20} className="text-gray-600" />}
            />
            <IconButton
              onClick={onAddToCart}
              icon={<ShoppingCart size={20} className="text-gray-600" />}
            />
          </div>
        </div>
      </div>
      {/* Name */}
      <div>
        <h2 className="text-lg text-stone-800">{data.itemName}</h2>
      </div>
      {/* Price */}
        <Currency value={data?.price} discountPrice={data.discountPrice} />
      {/* Description */}
      <div>
        <p className="text-sm text-gray-500">stock: {data.qtyInStock}</p>
      </div>
    </div>
  );
};

export default ProductItemCard;
