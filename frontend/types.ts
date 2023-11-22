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
  id: string;
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

export interface Image {
  id: string;
  url: string;
}

export interface ProductItem {
  id: number;
  name: string;
  productId: number;
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
