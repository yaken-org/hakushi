import { NextResponse } from "next/server";

export async function GET({params: {post_id}}: {params: {post_id: string}}) {
    const post = await fetch(`/api/backend/post/${post_id}`);
    if (!post.ok) {
        return NextResponse.json({message: "Post not found"}, {status: 404});
    }
    const data = await post.json();
    return NextResponse.json(data);
}