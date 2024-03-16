import type { Metadata } from "next";

import "./globals.css";

export const runtime = "edge"

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
      <body>{children}</body>
    </html>
  );
}
