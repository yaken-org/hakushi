"use client";
import { usePathname } from "next/navigation";
import { Button } from "../ui/button";
import { signIn, signOut, useSession } from "next-auth/react";

export default function SignInButton() {
    const nowPath = usePathname();
    const { data: session } = useSession();
    if (session) {
        return (
            <Button onClick={() => signOut()}>SignOut</Button>
        )
    }
    return (
        <Button onClick={() => signIn("google", { redirectTo: nowPath})}>SignIn</Button>
    )
}