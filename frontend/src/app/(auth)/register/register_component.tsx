"use client";
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { useForm } from "react-hook-form"
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useSession } from "next-auth/react";
import { useState } from "react";
import { cn } from "@/lib/utils";
import { useRouter } from "next/navigation";

const formSchema = z.object({
    name: z.string().min(4, {
        message: "ユーザーIDは4文字以上で入力してください",
    }).max(20, {
        message: "ユーザーIDは20文字以下で入力してください",
    }).refine(value => /^[a-zA-Z0-9_]+$/.test(value), {
        message: "ユーザーIDは半角英数字とアンダースコア(_)のみ使用できます",
    }),
    display_name: z.string().min(4, {
        message: "表示名は4文字以上で入力してください",
    }).max(20, {
        message: "表示名は20文字以下で入力してください",
    }),
    icon_url: z.string(),
    sub: z.string(), // googleのsub
});
type FormValues = z.infer<typeof formSchema>;

export default function RegistrationComponent() {
    const router = useRouter();
    const { data: session } = useSession();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const form = useForm<FormValues>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: session?.user.email?.split("@")[0].replaceAll(".", "_") ?? "",
            display_name: session?.user.name ?? "",
            icon_url: session?.user.image ?? "",
            sub: session?.user.providerAccountId ?? "",
        }
    });

    const handleSubmit = async (values: FormValues) => {
        console.log(values);
        setIsSubmitting(true);

        const isNameExists = await fetch("/api/backend/account/name/" + values.name);

        if (isNameExists.ok) {
            const isNameExistsJson = await isNameExists.json();
            if (isNameExistsJson.is_exists) {
                setIsSubmitting(false);
                form.setError("name", {
                    type: "manual",
                    message: "このユーザーIDは既に使用されています",
                });
            }
        } else {
            setIsSubmitting(false);
            form.setError("name", {
                type: "manual",
                message: "ユーザーIDの重複チェックに失敗しました",
            });
        }

        const result = await fetch("/api/backend/account", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(values),
        });

        if (result.ok) {
            setIsSubmitting(false);
            router.push("/");
        } else {
            setIsSubmitting(false);
            form.setError("name", {
                type: "manual",
                message: "アカウントの作成に失敗しました",
            });
        }
    }

    return (
        <Card className="w-full max-w-md mx-auto">
            <CardHeader>
                <CardTitle>初回登録</CardTitle>
                <CardDescription>アカウント情報を完成させてください</CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(handleSubmit)} className="space-y-8">
                        <FormField
                            control={form.control}
                            name="name"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>ユーザーID</FormLabel>
                                    <FormControl>
                                        <Input placeholder="ユーザーID" {...field} />
                                    </FormControl>
                                    <FormDescription>
                                        他のユーザーと重複しない、4文字以上20文字以下の英数字とアンダースコア(_)で入力してください
                                    </FormDescription>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="display_name"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>表示名</FormLabel>
                                    <FormControl>
                                        <Input placeholder="表示名" {...field} />
                                    </FormControl>
                                    <FormDescription>
                                        4文字以上20文字以下で入力してください
                                    </FormDescription>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <Button
                            type="submit"
                            className={cn("w-full", 
                                isSubmitting ? "bg-gray-300 text-gray-500 cursor-not-allowed" : ""
                            )}
                            disabled={isSubmitting}
                        >
                            登録
                        </Button>
                    </form>
                </Form>
            </CardContent>
        </Card>
    )
}