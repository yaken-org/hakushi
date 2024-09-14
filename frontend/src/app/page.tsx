import SignInButton from "@/components/auth/signInButton";
import { auth } from "@/lib/auth";

export default async function Home() {
    const session = await auth();
    return (
        <main className="w-full h-full">
            <h1 className="text-2xl font-bold">hakushi</h1>
            <SignInButton />
            {session && <p>{ session.user?.name }</p> }
        </main>
    );
}