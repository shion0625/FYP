"use client";

import NextImage from "next/image";
import { Tab } from "@headlessui/react";

import GalleryTab from "./gallery-tab";

interface GalleryProps {
  id: string;
  urls: string[];
}

const Gallery: React.FC<GalleryProps> = ({ id, urls = [] }) => {
  return (
    <Tab.Group as="div" className="flex flex-col-reverse">
      <div className="mx-auto mt-6 hidden w-full max-w-2xl sm:block lg:max-w-none">
        <Tab.List className="grid grid-cols-4 gap-6">
          {urls.map((url, i) => (
            <GalleryTab key={id + i} url={url} />
          ))}
        </Tab.List>
      </div>
      <Tab.Panels className="aspect-square w-full">
        {urls.map((url, i) => (
          <Tab.Panel key={id + i}>
            <div className="aspect-square relative h-full w-full sm:rounded-lg overflow-hidden">
              <NextImage
                fill
                sizes="(max-width: 600px) 100vw, 600px"
                src={url}
                alt="Image"
                className="object-cover object-center"
              />
            </div>
          </Tab.Panel>
        ))}
      </Tab.Panels>
    </Tab.Group>
  );
};

export default Gallery;
