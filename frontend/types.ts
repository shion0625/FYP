export interface Billboard {
  id: string;
  label: string;
  imageUrl: string;
}

export interface Category {
  id: string;
  name: string;
}

export interface Product {
  id: number;
  name: string;
  description: string;
  categoryId: string;
  brandId: string;
  price: number;
  discountPrice: number;
  image: string;
  createdAt: string;
  updatedAt: string;
}

export interface ProductItem {
  id: number;
  name: string;
  productId: number;
  itemName: string;
  price: number;
  discountPrice: number;
  sku: string;
  qtyInStock: number;
  categoryName: string;
  mainCategoryName: string;
  brandId: number;
  brandName: string;
  variationValues: ProductVariationValue[];
  images: string[];
}

export interface ProductVariationValue {
  variationId: number;
  name: string;
  variationOptionId: number;
  value: string;
}

export interface Response<T> {
  status: boolean;
  message: string;
  error?: any;
  data: T | null;
}
