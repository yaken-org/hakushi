import { API_Post } from "@/types/api/post";
import PostSimpleCard from "./post_simple_card";

export default function PostSimpleCardList({
    posts
}: {
    posts: API_Post[];
}) {
    return (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {posts.map((post) => (
                <PostSimpleCard key={post.id} post={post} />
            ))}
        </div>
    )
}