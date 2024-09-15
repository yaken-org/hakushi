import { z } from "zod";
import { PostSchema } from "../db/post";

const API_UserAccountSchema = z.object({
    id: z.number(),
    name: z.string(),
    display_name: z.string(),
    icon_url: z.string(),
    sub: z.string(),
    created_at: z.string().datetime(),
    updated_at: z.string().datetime(),

    posts: z.array(PostSchema),
})

type API_UserAccount = z.infer<typeof API_UserAccountSchema>;

export { API_UserAccountSchema };
export type { API_UserAccount };
