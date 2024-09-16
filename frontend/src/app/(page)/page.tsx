import PostSimpleCardList from "@/components/post/post_simple_list";
import { auth } from "@/lib/auth";
import { API_Post } from "@/types/api/post";

export default async function Home() {
    const session = await auth();
    console.log(session);

    const backend_url = process.env.BACKEND_API_ENDPOINT ?? "http://hakushi-backend.default.svc.cluster.local/api";
    const result = await fetch(`${backend_url}/post`);
    const data = await result.json();

    return (
        <div className="pt-4">
            <h1 className="text-4xl font-bold">Timeline</h1>
            <PostSimpleCardList posts={data as unknown as API_Post[]} />
        </div>
    );
}