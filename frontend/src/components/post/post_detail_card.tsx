import { cn } from "@/lib/utils";
import { API_Post } from "@/types/api/post";
import Image from "next/image";
import Link from "next/link";

export default function PostDetailCard({
    post
}: {
    post: API_Post;
}) {
    return (
        <div className={cn("border border-slate-200")}>
            <header className={cn("flex p-2 gap-2 flex-row justify-start items-center")}>
                <Image className="rounded-full border border-slate-200" width="28" height="28" alt={`${post.user_account.display_name}'s icon`} src={`/placeholder-user.jpg`} />
                <div className="flex flex-col justify-start items-start">
                    <span className="text-xs text-slate-800">{post.user_account.display_name}</span>
                    <span className="text-xs text-slate-800">@{post.user_account.name}</span>
                </div>
            </header>
            <Image width="300" height="400" className={cn("w-full")} src={`${process.env.NEXT_PUBLIC_R2_STATIC_URL}/fkeyhack.jpg`} alt={post.title} />
            <div className={cn("p-2")}>
                <div className="flex gap-2 flex-row flex-nowrap justify-start items-baseline">
                    <h2 className="text-lg font-bold">{post.title}</h2>
                    <span className="text-xs text-slate-800">{`公開:${post.created_at.replace("T", " ").replace("Z", "")}`}</span>
                </div>
                <p className="text-sm whitespace-pre-wrap">{post.content}</p>
            </div>
            <footer className={cn("flex p-2 gap-2 flex-row justify-start items-center")}>
                {post.tags.map((tag) => (
                    <Link key={tag.id} href={`/tags/${tag.name}`} className={cn("text-sm text-blue-600 hover:text-blue-900 hover:underline")}>#{tag.name}</Link>
                ))}
            </footer>
        </div>
    )
}
