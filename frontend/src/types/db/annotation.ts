import { z } from "zod";

const AnnotationSchema = z.object({
    id: z.string(),
    post_id: z.string(),
    product_id: z.string(),
    display_name: z.string(),
    x: z.number(),
    y: z.number(),
    created_at: z.string().datetime(),
    updated_at: z.string().datetime(),
});

type Annotation = z.infer<typeof AnnotationSchema>;

const ProductSchema = z.object({
    id: z.number(),
    name: z.string(),
    link: z.string(),
    created_at: z.string().datetime(),
    updated_at: z.string().datetime(),
});

type Product = z.infer<typeof ProductSchema>;

export { AnnotationSchema, ProductSchema };
export type { Annotation, Product };