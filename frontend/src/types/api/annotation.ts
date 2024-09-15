import { z } from "zod";
import { ProductSchema } from "../db/annotation";
import { PostSchema } from "../db/post";

const API_AnnotationSchema = z.object({
    id: z.number(),


    display_name: z.string(),
    x: z.number(),
    y: z.number(),
    created_at: z.string().datetime(),
    updated_at: z.string().datetime(),

    post: PostSchema,
    product: ProductSchema,
});

type API_Annotation = z.infer<typeof API_AnnotationSchema>;

export { API_AnnotationSchema };
export type { API_Annotation };
