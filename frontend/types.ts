export interface User {
  googleImage: string;
  firstName: string;
  lastName: string;
  age: number;
  email: string;
  userName: string;
  phone: string;
  blockStatus: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface Address {
  id: number;
  name: string;
  phoneNumber: string;
  house: string;
  area: string;
  landMark: string;
  city: string;
  pincode: number;
  countryName: string;
  isDefault?: boolean;
}

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

export interface ProductItemInfo {
  productItemId: number;
  variationValues: ProductVariationValue[];
  count: number;
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
  error?: unknown;
  data: T | null;
}
export interface TokenResponse {
  accessToken: string;
  userId: string;
}

export interface PaymentMethod {
  id: number;
  number: string;
  cardCompany: string;
}
