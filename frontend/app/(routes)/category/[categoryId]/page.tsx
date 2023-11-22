import Container from "@/components/ui/container";
import getProducts from "@/actions/product/get-products";
import ProductList from "@/components/product-list";

export const revalidate = 0;

interface CategoryIdPageProps {
  params: {
    categoryId: string;
  };
}

const CategoryIdPage = async ({ params }: CategoryIdPageProps) => {
  const products = await getProducts({
    categoryId: params.categoryId,
  });

  return (
    <Container>
      <div className="space-y-10 pb-10">
        <div className="flex flex-col gap-y-8 px-4 sm:px-6 lg:px-8">
          <ProductList title="Featured Products" items={products} />
        </div>
      </div>
    </Container>
  );
};

export default CategoryIdPage;
