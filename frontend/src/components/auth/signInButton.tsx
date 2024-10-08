"use client";
import { Button } from "../ui/button";
import { signIn, signOut, useSession } from "next-auth/react";
import { cn } from "@/lib/utils";

export default function SignInButton({
    className,
}: Readonly<{
    className?: string;
}>) {
    const { data: session } = useSession();
    if (session) {
        return (
            <Button variant="outline" className={cn(className)} onClick={() => signOut()}>ログアウト</Button>
        )
    }
    return (
        <Button className={cn(className)} onClick={() => signIn("google", { callbackUrl: "/register" })}>Googleでログイン</Button>
    )
}