import PostSimpleCardList from "@/components/post/post_simple_list";
import { auth } from "@/lib/auth";
import { API_Post } from "@/types/api/post";

export default async function Home() {
    const session = await auth();
    console.log(session);

    const posts = await fetch('/api/backend/post');
    const data = await posts.json();

    return (
        <div className="pt-4">
            <h1 className="text-4xl font-bold">Timeline</h1>
            <PostSimpleCardList posts={data as unknown as API_Post[]} />
        </div>
    );
}