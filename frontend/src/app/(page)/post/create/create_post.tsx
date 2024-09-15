/* eslint-disable @next/next/no-img-element */
"use client";
import React, { useState } from 'react'
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Trash, Upload } from 'lucide-react'
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { DrawerDialog } from './annotation_dialog';
import { useRouter } from 'next/navigation';


export default function CreaetePost() {
    const [formData, setFormData] = useState({
        title: '',
        content: '',
    });
    const [rawImage, setRawImage] = useState<File | null>(null);
    const [image, setImage] = useState<string | null>(null);
    const [annotationOpen, setAnnotationOpen] = useState(false);
    const [annotations, setAnnotations] = useState<{
        x: number;
        y: number;
        text: string;
    }[]>([]);
    const [tags, setTags] = useState<{
        tag: string;
        date: string;
    }[]>([])
    const router = useRouter();

    const handleImageUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
        const file = event.target.files?.[0]
        setRawImage(file as File)
        if (file) {
            const reader = new FileReader()
            reader.onloadend = () => {
                setImage(reader.result as string)
            }
            reader.readAsDataURL(file)
        }
    }

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        if (!rawImage) {
            return alert('画像を選択してください')
        }

        const uint8Array = Array.from(new Uint8Array(await rawImage.arrayBuffer()));
        let encodedStr = '';
        for (let i = 0; i < uint8Array.length; i += 1024) {
            encodedStr += String.fromCharCode.apply(null, uint8Array.slice(i, i + 1024));
        }
        const base64Str = `data:${rawImage.type};base64,` + btoa(encodedStr);

        const body = {
            image: base64Str,
            title: formData.title,
            content: formData.content,
            tags: tags.map(tag => tag.tag).filter(tag => tag !== ''),
            annotations,
        }

        console.log(body)

        const response = await fetch('/api/backend/post', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(body)
        })

        if (response.ok) {
            alert('投稿しました！')
            router.push('/')
        } else {
            alert('投稿に失敗しました')
        }

    }

    return (
        <Card className="w-full max-w-md mx-auto">
            <CardHeader>
                <CardTitle>新規投稿</CardTitle>
            </CardHeader>
            <CardContent>
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div className="grid w-full items-center gap-10">
                        <div className="flex flex-col space-y-1.5">
                            <label htmlFor="image-upload" className="cursor-pointer">
                                {image ? (
                                    <img src={image} alt="アップロード画像" className="w-full aspect-square object-cover rounded-md" />
                                ) : (
                                    <div className="w-full aspect-square border-2 border-dashed border-gray-300 rounded-md flex items-center justify-center">
                                        <Upload className="h-12 w-12 text-gray-400" />
                                    </div>
                                )}
                            </label>
                            <input
                                id="image-upload"
                                type="file"
                                accept="image/*"
                                onChange={handleImageUpload}
                                className="hidden"
                            />
                        </div>
                        <DrawerDialog open={annotationOpen} setOpen={setAnnotationOpen} image={image} annotations={annotations} setAnnotations={setAnnotations} />
                        {image && (
                            <div className="flex flex-row justify-between items-center">
                                <Button
                                    onClick={() => setImage(null)}
                                    type='button'
                                    variant="outline"
                                >
                                    画像を削除
                                </Button>
                                <Button
                                    onClick={() => setAnnotationOpen(true)}
                                    type='button'
                                    variant="outline"
                                >
                                    製品情報を追加する
                                </Button>
                            </div>
                        )}
                        {image && annotations && annotations.map((annotation, index) => (
                            <div key={index} className="absolute" style={{ left: `${annotation.x}%`, top: `${annotation.y}%` }}>
                                <div className="absolute w-8 h-8 bg-gray-800 rounded-full" />
                                <div className="absolute w-8 h-8 bg-white rounded-full" />
                            </div>
                        ))}
                        <div className="flex flex-col space-y-1.5">
                            <label htmlFor="title" className="text-sm font-semibold text-gray-600">タイトル</label>
                            <Input
                                id="title"
                                name='title'
                                placeholder="タイトル"
                                value={formData.title}
                                onChange={(event) => setFormData({ ...formData, title: event.target.value })}
                            />
                        </div>
                        <div className="flex flex-col space-y-1.5">
                            <label htmlFor="content" className="text-sm font-semibold text-gray-600">本文</label>
                            <Textarea
                                id="content"
                                name='content'
                                placeholder="本文"
                                rows={6}
                                value={formData.content}
                                onChange={(event) => setFormData({ ...formData, content: event.target.value })}
                            />
                        </div>
                        
                        {tags.map(tag => (
                            <div key={`${tag.date}`} className="flex flex-col space-y-1.5">
                                <label htmlFor={`tag-${tag.date}`} className="text-sm font-semibold text-gray-600">タグ</label>
                                <div className="flex justify-between items-center">
                                    <Input
                                        id={`tag-${tag.date}`}
                                        name={`tag-${tag.date}`}
                                        placeholder="タグ"
                                        onChange={(event) => {
                                            setTags(tags.map(v => v.date === tag.date ? { ...v, tag: event.target.value } : v))
                                        }}
                                    />
                                    <Button
                                        onClick={() => {
                                            setTags(tags.filter(v => v.date !== tag.date))
                                        }}
                                        type='button'
                                        variant="outline"
                                    >
                                        <Trash className="h-5 w-5" />
                                    </Button>
                                </div>
                            </div>
                        ))}

                        <Button
                            onClick={() => {
                                setTags([...tags, {
                                    tag: '',
                                    date: new Date().toISOString()
                                }])
                            }}
                            type='button'
                            variant="outline"
                        >
                            タグを追加
                        </Button>
                    </div>
                    <Button className="w-full" type="submit" onClick={() => handleSubmit} >
                        投稿する
                    </Button>
                </form>
            </CardContent>
        </Card>
    )
}