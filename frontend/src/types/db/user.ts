import { z } from "zod";

const UserAccountSchema = z.object({
    id: z.number(),
    name: z.string(),
    display_name: z.string(),
    icon_url: z.string(),
    sub: z.string(),
    created_at: z.string().datetime(),
    updated_at: z.string().datetime(),
})

type UserAccount = z.infer<typeof UserAccountSchema>;

export { UserAccountSchema };
export type { UserAccount };