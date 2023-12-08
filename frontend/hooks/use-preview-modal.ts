import { ProductItem } from '@/types';
import { create } from 'zustand';

interface PreviewModalStore {
  isOpen: boolean;
  data?: ProductItem;
  updateItem: (data: ProductItem) => void;
  onOpen: (data: ProductItem) => void;
  onClose: () => void;
}

const usePreviewModal = create<PreviewModalStore>((set) => ({
  isOpen: false,
  data: undefined,
  updateItem: (data: ProductItem) => set({ data }),
  onOpen: (data: ProductItem) => set({ isOpen: true, data }),
  onClose: () => set({ isOpen: false }),
}));

export default usePreviewModal;
