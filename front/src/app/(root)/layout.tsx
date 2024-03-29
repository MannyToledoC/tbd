import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import HeaderBar from "@/Components/shared/HeaderBar/HeaderBar";
import LeftSidebar from "@/Components/shared/LeftSidebar/LeftSidebar";
import RightSidebar from "@/Components/shared/RightSidebar/RightSidebar";
import FooterBar from "@/Components/shared/FooterBar/FooterBar";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <HeaderBar />
        <main>
          <LeftSidebar />
          <section className="app-container">
            {children}
          </section>
          <RightSidebar />
        </main>
        <FooterBar />
      </body>
    </html>
  );
}
