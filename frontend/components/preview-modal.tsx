'use client';

import { MouseEventHandler } from 'react';
import React, { useState, useMemo, useEffect } from 'react';

import { Button } from 'flowbite-react';
import { ShoppingCart } from 'lucide-react';

import Gallery from '@/components/gallery';
import ProductItemDetail from '@/components/product/product-item-detail';
import Modal from '@/components/ui/modal';

import useCart from '@/hooks/use-cart';
import usePreviewModal from '@/hooks/use-preview-modal';
import { ProductVariationValue } from '@/types';

const PreviewModal = () => {
  const previewModal = usePreviewModal();
  const productItem = usePreviewModal((state) => state.data);
  const cart = useCart();

  // Nameの配列（重複なし）
  const names = useMemo(
    () =>
      productItem && productItem.variationValues
        ? Array.from(new Set(productItem.variationValues.map((item) => item.name)))
        : [],
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [productItem?.variationValues]
  );

  // Valueのmap配列(keyがname)
  const valuesMap = useMemo(
    () =>
      productItem?.variationValues?.reduce(
        (acc: { [key: string]: ProductVariationValue[] }, curr) => {
          if (!acc[curr.name]) {
            acc[curr.name] = [];
          }
          acc[curr.name].push(curr);
          return acc;
        },
        {}
      ) || {},
    [productItem?.variationValues]
  );
  // 選択された値を追跡するためのstate
  const [selectedValues, setSelectedValues] = useState<{
    [key: string]: ProductVariationValue | null;
  }>({});

  // UseEffectを使用して、valuesMapが更新されたときにselectedValuesを更新します
  useEffect(() => {
    const initialSelectedValues = Object.keys(valuesMap).reduce(
      (acc, key) => {
        acc[key] = valuesMap[key][0] || null;
        return acc;
      },
      {} as { [key: string]: ProductVariationValue | null }
    );

    setSelectedValues(initialSelectedValues);
  }, [valuesMap]);

  // SelectedValuesを配列に変換
  const selectedValuesArray = useMemo(
    () => Object.values(selectedValues).filter(Boolean) as ProductVariationValue[],
    [selectedValues]
  );

  const onAddToCart: MouseEventHandler<HTMLButtonElement> = (event) => {
    event.stopPropagation();
    if (selectedValuesArray && productItem) {
      cart.addItem({
        ...productItem,
        variationValues: selectedValuesArray,
      });
    }
  };

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
          <ProductItemDetail
            data={productItem}
            names={names}
            variationsMap={valuesMap}
            selectedValues={selectedValues}
            setSelectedValues={setSelectedValues}
          />
        </div>
      </div>
      <Button color="dark" className="flex items-center gap-x-2 mx-auto mt-2" onClick={onAddToCart}>
        Add To Cart
        <ShoppingCart />
      </Button>
    </Modal>
  );
};

export default PreviewModal;
