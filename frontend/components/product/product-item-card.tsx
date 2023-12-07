// Frontend/components/product-item-card.tsx
import { MouseEventHandler } from 'react';

import Gallery from '@/components/gallery';
import Currency from '@/components/ui/currency';

import usePreviewModal from '@/hooks/use-preview-modal';
import { ProductItem } from '@/types';

interface ProductItemCardProps {
  data: ProductItem;
}

const ProductItemCard: React.FC<ProductItemCardProps> = ({ data }) => {
  const previewModal = usePreviewModal();
  const onPreview: MouseEventHandler<HTMLElement> = (event) => {
    event.stopPropagation();

    previewModal.onOpen(data);
  };

  return (
    <div
      className="bg-white group cursor-pointer rounded-xl border p-3 space-y-4"
      onClick={onPreview}
    >
      <div className="aspect-square rounded-xl bg-gray-100 relative">
        <div className="aspect-square rounded-xl bg-gray-100 relative mb-4">
          <Gallery id={data.sku} urls={data.images} />
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
