'use client';
import { Sidebar as SidebarFlow, CustomFlowbiteTheme } from 'flowbite-react';
import { Flowbite } from 'flowbite-react';
import { usePathname } from 'next/navigation';
import {
  HiArrowSmRight,
  HiChartPie,
  HiOutlineMinusSm,
  HiOutlinePlusSm,
  HiShoppingBag,
  HiTable,
} from 'react-icons/hi';
import { twMerge } from 'tailwind-merge';
import useLoginState from '@/hooks/use-login';
import useSidebar from '@/hooks/use-sidebar';
import { Category } from '@/types';

interface SidebarProps {
  data?: Category[] | null;
}

const Sidebar: React.FC<SidebarProps> = ({ data }) => {
  const pathname = usePathname();
  const sidebar = useSidebar();
  const loginState = useLoginState();
  const categories = data?.map((category) => ({
    href: `/category/${category.id}`,
    label: category.name,
    active: pathname === `/category/${category.id}`,
  }));

  const customTheme: CustomFlowbiteTheme = {
    sidebar: {
      collapse: {
        button:
          'group flex w-full items-center rounded-lg p-2 text-base font-normal text-gray-900 transition duration-75 hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700',
        icon: {
          base: 'h-6 w-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white',
          open: {
            off: '',
            on: 'text-gray-900',
          },
        },
        label: {
          base: 'ml-3 flex-1 whitespace-nowrap text-left',
          icon: {
            base: 'h-6 w-6 transition ease-in-out delay-0',
            open: {
              on: 'rotate-180',
              off: '',
            },
          },
        },
        list: 'space-y-2 py-2',
      },
    },
  };
  return (
    <>
      {sidebar.isOpen && (
        <Flowbite theme={{ theme: customTheme }}>
          <SidebarFlow
            aria-label="Sidebar with multi-level dropdown example"
            className="z-10 absolute"
          >
            <SidebarFlow.Items>
              <SidebarFlow.ItemGroup>
                {loginState.isLogin && (
                  <SidebarFlow.Item href="/user" icon={HiChartPie}>
                    My page
                  </SidebarFlow.Item>
                )}
                <SidebarFlow.Item href="/" icon={HiShoppingBag}>
                  Products
                </SidebarFlow.Item>
                <SidebarFlow.Collapse
                  icon={HiShoppingBag}
                  label="Category"
                  renderChevronIcon={(theme, open) => {
                    const IconComponent = open ? HiOutlineMinusSm : HiOutlinePlusSm;

                    return (
                      <IconComponent
                        aria-hidden
                        className={twMerge(theme?.label?.icon?.open?.[open ? 'on' : 'off'])}
                      />
                    );
                  }}
                >
                  {categories ? (
                    categories.map((category) => (
                      <SidebarFlow.Item key={category.href} href={category.href}>
                        {category.label}
                      </SidebarFlow.Item>
                    ))
                  ) : (
                    <div>No categories available</div>
                  )}
                </SidebarFlow.Collapse>
                {!loginState.isLogin ? (
                  <>
                    <SidebarFlow.Item href="/user/login" icon={HiArrowSmRight}>
                      login
                    </SidebarFlow.Item>
                    <SidebarFlow.Item href="/user/signup" icon={HiTable}>
                      Sign up
                    </SidebarFlow.Item>
                  </>
                ) : (
                  <SidebarFlow.Item href="/user/logout" icon={HiTable}>
                    logout
                  </SidebarFlow.Item>
                )}
              </SidebarFlow.ItemGroup>
            </SidebarFlow.Items>
          </SidebarFlow>
        </Flowbite>
      )}
    </>
  );
};

export default Sidebar;
