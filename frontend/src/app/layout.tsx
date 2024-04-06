import type { Metadata } from "next";

import "./globals.css";
import { Header } from "@/components/layouts/header";

export const metadata: Metadata = {
  title: "nemmy",
  description: "nemmy is an awesome web app",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ja">
      <body>
        <Header />
        <main>{children}</main>
      </body>
    </html>
  );
}
