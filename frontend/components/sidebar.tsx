"use client";
import type { CustomFlowbiteTheme } from "flowbite-react";
import { Sidebar as SidebarFlow } from "flowbite-react";
import {
  HiArrowSmRight,
  HiChartPie,
  HiInbox,
  HiOutlineMinusSm,
  HiOutlinePlusSm,
  HiShoppingBag,
  HiTable,
  HiUser,
} from "react-icons/hi";
import { twMerge } from "tailwind-merge";
import { Flowbite } from "flowbite-react";
import { Category } from "@/types";
import { usePathname } from "next/navigation";
import useSidebar from "@/hooks/use-sidebar";

interface SidebarProps {
  data?: Category[] | null;
}

const Sidebar: React.FC<SidebarProps> = ({ data }) => {
  const pathname = usePathname();
  const sidebar = useSidebar();
  if (!data || data.length === 0) {
    return <div>No categories available</div>;
  }

  const categories = data.map((category) => ({
    href: `/category/${category.id}`,
    label: category.name,
    active: pathname === `/category/${category.id}`,
  }));

  const customTheme: CustomFlowbiteTheme = {
    sidebar: {
      collapse: {
        button:
          "group flex w-full items-center rounded-lg p-2 text-base font-normal text-gray-900 transition duration-75 hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700",
        icon: {
          base: "h-6 w-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white",
          open: {
            off: "",
            on: "text-gray-900",
          },
        },
        label: {
          base: "ml-3 flex-1 whitespace-nowrap text-left",
          icon: {
            base: "h-6 w-6 transition ease-in-out delay-0",
            open: {
              on: "rotate-180",
              off: "",
            },
          },
        },
        list: "space-y-2 py-2",
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
                <SidebarFlow.Item href="#" icon={HiChartPie}>
                  Dashboard
                </SidebarFlow.Item>
                <SidebarFlow.Collapse
                  icon={HiShoppingBag}
                  label="Category"
                  renderChevronIcon={(theme, open) => {
                    const IconComponent = open
                      ? HiOutlineMinusSm
                      : HiOutlinePlusSm;

                    return (
                      <IconComponent
                        aria-hidden
                        className={twMerge(
                          theme?.label?.icon?.open?.[open ? "on" : "off"]
                        )}
                      />
                    );
                  }}
                >
                  {categories.map((category) => (
                    <SidebarFlow.Item key={category.href} href={category.href}>
                      {category.label}
                    </SidebarFlow.Item>
                  ))}
                </SidebarFlow.Collapse>
                <SidebarFlow.Item href="#" icon={HiInbox}>
                  Inbox
                </SidebarFlow.Item>
                <SidebarFlow.Item href="#" icon={HiUser}>
                  Users
                </SidebarFlow.Item>
                <SidebarFlow.Item href="#" icon={HiShoppingBag}>
                  Products
                </SidebarFlow.Item>
                <SidebarFlow.Item href="#" icon={HiArrowSmRight}>
                  Sign In
                </SidebarFlow.Item>
                <SidebarFlow.Item href="#" icon={HiTable}>
                  Sign Up
                </SidebarFlow.Item>
              </SidebarFlow.ItemGroup>
            </SidebarFlow.Items>
          </SidebarFlow>
        </Flowbite>
      )}
    </>
  );
};

export default Sidebar;
