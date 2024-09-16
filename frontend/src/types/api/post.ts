import { z } from "zod";
import { AnnotationSchema } from "../db/annotation";
import { TagSchema } from "../db/tag";
import { UserAccountSchema } from "../db/user";

const API_PostSchema = z.object({
    id: z.number(),

    image_id: z.number(),
    title: z.string(),
    content: z.string(),
    created_at: z.string().datetime(),
    updated_at: z.string().datetime(),

    user_account: UserAccountSchema,
    annotations: z.array(AnnotationSchema),
    tags: z.array(TagSchema),
})

type API_Post = z.infer<typeof API_PostSchema>;

export { API_PostSchema };
export type { API_Post }; 