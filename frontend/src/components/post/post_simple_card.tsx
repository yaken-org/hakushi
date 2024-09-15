import { Card, CardHeader, CardIcon, CardDescription, CardImage } from "../ui/card";
import { API_Post } from "@/types/api/post";

export default function PostSimpleCard({
    post
}: {
    post: API_Post;
}) {
    return (
        <Card>
            <CardHeader>
                <CardIcon src={`/placeholder-user.jpg`} />
                <CardDescription>@{post.user_account.id}</CardDescription>
            </CardHeader>
            <CardImage src={`${process.env.NEXT_PUBLIC_R2_STATIC_URL}/fkeyhack.jpg`} alt={post.title} />
        </Card>
    )
}