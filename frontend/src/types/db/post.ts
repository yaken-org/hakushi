import { z } from "zod";

const PostSchema = z.object({
    id: z.number(),
    user_account_id: z.number(),
    image_id: z.number(),
    title: z.string(),
    content: z.string(),
    created_at: z.string().datetime(),
    updated_at: z.string().datetime(),
})

type Post = z.infer<typeof PostSchema>;

export { PostSchema };
export type { Post };