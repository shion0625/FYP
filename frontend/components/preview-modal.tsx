"use client";

import usePreviewModal from "@/hooks/use-preview-modal";
import ProductItemDetail from "@/components/product-item-detail";
import Modal from "@/components/ui/modal";
import Gallery from "@/components/gallery";

const PreviewModal = () => {
  const previewModal = usePreviewModal();
  const productItem = usePreviewModal((state) => state.data);

  if (!productItem) {
    return null;
  }

  return (
    <Modal open={previewModal.isOpen} onClose={previewModal.onClose}>
      <div className="grid grid-cols-3 gap-4 ">
        <div className="col-span-1">
          <Gallery id={productItem.sku} urls={productItem.images} />
        </div>
        <div className="col-span-2 pr-6">
          <ProductItemDetail data={productItem} />
        </div>
      </div>
    </Modal>
  );
};

export default PreviewModal;
