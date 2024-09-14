"use client";
import { usePathname } from "next/navigation";
import { Button } from "../ui/button";
import { signIn, signOut, useSession } from "next-auth/react";

export default function SignInButton() {
    const nowPath = usePathname();
    console.log(nowPath);
    const { data: session } = useSession();
    if (session) {
        console.log(session);
        return (
            <Button onClick={() => signOut()}>SignOut</Button>
        )
    }
    return (
        <Button onClick={() => signIn("google", { redirectTo: "/"})}>SignIn</Button>
    )
}