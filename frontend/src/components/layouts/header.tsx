"use client";

import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
} from "@radix-ui/react-navigation-menu";
import Link from "next/link";
import { useState } from "react";

import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { Button } from "../ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { navigationMenuTriggerStyle } from "../ui/navigation-menu";
import { Separator } from "../ui/separator";
import { Skeleton } from "../ui/skeleton";

const navigationItems = [
  { href: "/", label: "ホーム" },
  { href: "/subjects", label: "教材一覧" },
] as const;

export const Header = () => {
  const [loggedIn, setLoggedIn] = useState(false);

  return (
    <>
      <header className="flex p-5 gap-2 justify-between">
        <h1 className="text-2xl">nemmy</h1>
        <NavigationMenu>
          <NavigationMenuList className="flex">
            {navigationItems.map((item) => (
              <NavigationMenuItem key={item.href}>
                <Link href={item.href} legacyBehavior passHref>
                  <NavigationMenuLink
                    className={`text-base ${navigationMenuTriggerStyle()}`}
                  >
                    {item.label}
                  </NavigationMenuLink>
                </Link>
              </NavigationMenuItem>
            ))}
          </NavigationMenuList>
        </NavigationMenu>
        {loggedIn ? (
          <DropdownMenu>
            <DropdownMenuTrigger>
              <Avatar>
                <AvatarImage src="https://github.com/light-planck.png" />
                <AvatarFallback>
                  <Skeleton className="h-12 w-12 rounded-full" />
                </AvatarFallback>
              </Avatar>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem>プロフィール</DropdownMenuItem>
              <DropdownMenuItem>設定</DropdownMenuItem>
              <DropdownMenuItem onClick={() => setLoggedIn(false)}>
                ログアウト
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        ) : (
          <Button onClick={() => setLoggedIn(true)}>ログイン</Button>
        )}
      </header>
      <Separator />
    </>
  );
};
