import { API_Post } from "@/types/api/post";
import Image from "next/image";
import Link from "next/link";

export default function PostSimpleCard({
    post
}: {
    post: API_Post;
}) {
    return (
        <Link href={`/post/${post.id}`} className="border border-slate-200 transition-shadow hover:shadow-lg group">
            <div className="overflow-hidden">
                <Image width="300" height="400" className="w-full object-cover transition-transform group-hover:scale-110" src={`${process.env.NEXT_PUBLIC_R2_STATIC_URL}/fkeyhack.jpg`} alt={post.title} />
            </div>
            <header className="flex p-2 gap-2 flex-row justify-start items-center">
                <Image className="rounded-full border border-slate-200" width="28" height="28" alt={`${post.user_account.display_name}'s icon`} src={`/placeholder-user.jpg`} />
                <div className="flex flex-col justify-start items-start">
                    <span className="text-xs text-slate-800">{post.user_account.display_name}</span>
                    <span className="text-xs text-slate-800">@{post.user_account.name}</span>
                </div>
            </header>
        </Link>
    )
}
