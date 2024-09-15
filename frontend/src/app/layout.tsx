import type { Metadata } from "next";
import { IBM_Plex_Sans_JP } from "next/font/google";
import "./globals.css";
import AuthSessionProvider from "@/components/auth/sessionProvider";

const ibmplexsansjp = IBM_Plex_Sans_JP({
    weight: "400",
    subsets: ["latin"],
    display: "swap",
});

export const metadata: Metadata = {
    title: "hakushi",
    description: "hakushiプロジェクトです",
};

export default function RootLayout({
    children,
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="ja">
            <body
                className={`${ibmplexsansjp.className} antialiased overflow-hidden`}
            >
                <AuthSessionProvider>
                    {children}
                </AuthSessionProvider>
            </body>
        </html>
    );
}
