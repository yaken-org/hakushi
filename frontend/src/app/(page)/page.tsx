import PostSimpleCard from "@/components/post/post_simple_card";
import PostSimpleCardList from "@/components/post/post_simple_list";
import { auth } from "@/lib/auth";

export default async function Home() {
    const session = await auth();
    return (
        <div className="pt-4">
            <h1 className="text-4xl font-bold">Timeline</h1>
            <PostSimpleCardList posts={[
                {
                    id: 1,
                    title: "test",
                    user_account: {
                        id: 12,
                        name: "test",
                        created_at: "2021-10-01T00:00:00Z",
                        updated_at: "2021-10-01T00:00:00Z",
                        display_name: "",
                        icon_url: "",
                        sub: ""
                    },
                    image_id: 0,
                    content: "",
                    created_at: "",
                    updated_at: "",
                    annotations: [],
                    tags: []
                },
                {
                    id: 2,
                    title: "test2",
                    user_account: {
                        id: 12,
                        name: "test",
                        created_at: "2021-10-01T00:00:00Z",
                        updated_at: "2021-10-01T00:00:00Z",
                        display_name: "",
                        icon_url: "",
                        sub: ""
                    },
                    image_id: 0,
                    content: "",
                    created_at: "",
                    updated_at: "",
                    annotations: [],
                    tags: []
                }
            ]} />
        </div>
    );
}