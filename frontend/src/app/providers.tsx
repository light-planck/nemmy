"use client";

import { NextUIProvider } from "@nextui-org/react";
import { FC, ReactNode } from "react";

interface Props {
  children: ReactNode;
}

export const Providers: FC<Props> = ({ children }) => {
  return <NextUIProvider>{children}</NextUIProvider>;
};
