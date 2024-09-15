import { auth } from "@/lib/auth";
import { redirect } from "next/navigation";
import CreaetePost from "./create_post";

export default async function PostCreate() {
    const session = await auth();

    if (!session?.user.data) {
        redirect("/register");
    }
    return (
        <div>
            <CreaetePost />
        </div>
    );
}