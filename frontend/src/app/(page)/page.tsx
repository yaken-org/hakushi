import PostSimpleCardList from "@/components/post/post_simple_list";
import { auth } from "@/lib/auth";

export default async function Home() {
    const session = await auth();
    console.log(session);

    return (
        <div className="pt-4">
            <h1 className="text-4xl font-bold">Timeline</h1>
            <PostSimpleCardList posts={[
                {
                    id: 1,
                    title: "test",
                    user_account: {
                        id: 12,
                        name: "nenrin_yearring",
                        created_at: "2021-10-01T00:00:00Z",
                        updated_at: "2021-10-01T00:00:00Z",
                        display_name: "年輪",
                        icon_url: "",
                        sub: ""
                    },
                    image_id: 0,
                    content: "aiueokakikukekosasisusesotatituteto\naaaaaaaa",
                    created_at: "2024-09-10T00:00:00Z",
                    updated_at: "",
                    annotations: [],
                    tags: [
                        {
                            id: 1,
                            name: "test1",
                            created_at: "2021-10-01T00:00:00Z",
                            updated_at: "2021-10-01T00:00:00Z"
                        },
                        {
                            id: 2,
                            name: "test2",
                            created_at: "2021-10-01T00:00:00Z",
                            updated_at: "2021-10-01T00:00:00Z"
                        }
                    ]
                },
                {
                    id: 2,
                    title: "test2",
                    user_account: {
                        id: 12,
                        name: "rokuosan",
                        created_at: "2021-10-01T00:00:00Z",
                        updated_at: "2021-10-01T00:00:00Z",
                        display_name: "ろくおー",
                        icon_url: "",
                        sub: ""
                    },
                    image_id: 0,
                    content: "aiueokakikukekosasisusesotatituteto\naaaaaaaa",
                    created_at: "2024-09-10T00:00:00Z",
                    updated_at: "",
                    annotations: [],
                    tags: []
                }
            ]} />
        </div>
    );
}