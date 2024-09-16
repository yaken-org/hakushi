import { auth } from "@/lib/auth";
import RegistrationComponent from "./register_component";

export default async function RegisterPage() {
    const session = await auth();

    return (
        <main className="w-full h-full min-h-screen flex justify-center items-center">
            <RegistrationComponent />
            <div className="absolute top-4 right-4">
                <div className="text-sm text-gray-400">Logged in as {session?.user.email}</div>
            </div>
        </main>
    );
}