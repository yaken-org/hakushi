import PostDetailCard from "@/components/post/post_detail_card";
import { API_Post } from "@/types/api/post";

export default async function PostDetailPage({
    params: {
        post_id
    }
}: {
    params: {
        post_id: string;
    };
    }) {
    const post = await fetch('/api/backend/post/' + post_id);
    if (!post.ok) {
        return <div>Post not found</div>
    }
    
    return (
        <div>
            <PostDetailCard post={post as unknown as API_Post} />
        </div>
    )
}