import { z } from "zod";

const PostTagSchema = z.object({
    post_id: z.number(),
    tag_id: z.number(),
})

type PostTag = z.infer<typeof PostTagSchema>;

const TagSchema = z.object({
    id: z.number(),
    name: z.string(),
    created_at: z.string().datetime(),
    updated_at: z.string().datetime(),
})

type Tag = z.infer<typeof TagSchema>;

export { PostTagSchema, TagSchema };
export type { PostTag, Tag };