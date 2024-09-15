"use client";
import { usePathname } from "next/navigation";
import { Button } from "../ui/button";
import { signIn, signOut, useSession } from "next-auth/react";
import { cn } from "@/lib/utils";

export default function SignInButton({
    className,
}: Readonly<{
    className?: string;
}>) {
    const nowPath = usePathname();
    const { data: session } = useSession();
    console.log(session);
    if (session) {
        return (
            <Button variant="outline" className={cn(className)} onClick={() => signOut()}>ログアウト</Button>
        )
    }
    return (
        <Button className={cn(className)} onClick={() => signIn("google", { redirectTo: nowPath})}>Googleでログイン</Button>
    )
}