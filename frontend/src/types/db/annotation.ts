import { z } from "zod";

const AnnotationSchema = z.object({
    id: z.number(),
    post_id: z.number(),
    product_id: z.number(),
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