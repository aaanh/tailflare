import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { ThemeProvider } from "next-themes";
import Sidebar from "@/components/sidebar";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Tailflare UI",
  description: "From Tailscale to Cloudflare, hopefully without any issues."
};

export default function RootLayout({
  children
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <html lang="en" suppressHydrationWarning>
        <head />
        <body>
          <ThemeProvider
            attribute="class"
            defaultTheme="system"
            enableSystem
            disableTransitionOnChange
          >
            <div className="grid grid-cols-[75px_1fr] gap-2">
              <Sidebar></Sidebar>
              <div>{children}</div>
            </div>
          </ThemeProvider>
        </body>
      </html>
    </>
  );
}
