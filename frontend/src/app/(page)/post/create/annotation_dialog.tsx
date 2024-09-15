/* eslint-disable @next/next/no-img-element */
import * as React from "react"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
} from "@/components/ui/dialog"
import {
    Drawer,
    DrawerClose,
    DrawerContent,
    DrawerDescription,
    DrawerFooter,
    DrawerHeader,
    DrawerTitle,
} from "@/components/ui/drawer"
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { Input } from "@/components/ui/input"
import { useEffect } from "react"

export function DrawerDialog({
    open,
    setOpen,
    image,
    annotations,
    setAnnotations,
}: {
    open: boolean
    setOpen: React.Dispatch<React.SetStateAction<boolean>>
    image: string | null
    annotations: {
        x: number
        y: number
        text: string
    }[]
    setAnnotations: React.Dispatch<
        React.SetStateAction<{
            x: number
            y: number
            text: string
        }[]>
    >
}) {

    const [isDesktop, setIsDesktop] = React.useState(window.innerWidth >= 768)
    useEffect(() => {
        const handleResize = () => {
            setIsDesktop(window.innerWidth >= 768)
        }

        window.addEventListener('resize', handleResize)
        return () => {
            window.removeEventListener('resize', handleResize)
        }
    }, [])

    if (image === null) {
        return null
    }

    if (isDesktop) {
        return (
            <Dialog open={open} onOpenChange={setOpen}>
                <DialogContent className="max-h-screen overflow-y-scroll sm:max-w-[400px] md:max-w-[600px] lg:max-w-[800px]">
                    <DialogHeader>
                        <DialogTitle>製品情報を編集</DialogTitle>
                        <DialogDescription>
                            画像をクリックで、その地点に注釈を追加できます。
                        </DialogDescription>
                    </DialogHeader>
                    <AnnotationForm image={image} annotations={annotations} setAnnotations={setAnnotations} setOpen={setOpen} />
                </DialogContent>
            </Dialog>
        )
    }

    return (
        <Drawer open={open} onOpenChange={setOpen}>
            <DrawerContent className="max-h-screen overflow-y-scroll">
                <DrawerHeader className="text-left">
                    <DrawerTitle>製品情報を編集</DrawerTitle>
                    <DrawerDescription>
                        画像をクリックで、その地点に注釈を追加できます。
                    </DrawerDescription>
                </DrawerHeader>
                <AnnotationForm image={image} annotations={annotations} setAnnotations={setAnnotations} setOpen={setOpen} className="px-4" />
                <DrawerFooter className="pt-2">
                    <DrawerClose asChild>
                        <Button variant="outline">Cancel</Button>
                    </DrawerClose>
                </DrawerFooter>
            </DrawerContent>
        </Drawer>
    )
}

function AnnotationForm({
    className,
    image,
    annotations,
    setAnnotations,
    setOpen
}: {
        className?: string,
        image?: string,
        annotations?: {
            x: number, y: number, text: string
        }[],
        setAnnotations?: React.Dispatch<React.SetStateAction<{ x: number, y: number, text: string }[]>>
        setOpen?: React.Dispatch<React.SetStateAction<boolean>>
}) {
    const [selectedAnnotation, setSelectedAnnotation] = React.useState<{ x: number, y: number } | null>(null)

    const image_click_handler = (e: React.MouseEvent<HTMLImageElement>) => {
        const rect = e.currentTarget.getBoundingClientRect()
        const x = e.clientX - rect.left
        const y = e.clientY - rect.top
        console.log(`x: ${x}, y: ${y}`)

        const percentage_x = parseInt(x / rect.width * 1000 + "")
        const percentage_y = parseInt(y / rect.height * 1000 + "")
        console.log(`percentage_x: ${percentage_x}, percentage_y: ${percentage_y}`)

        setAnnotations([...annotations, { x: percentage_x, y: percentage_y, text: "" }])
    }

    return (
        <form className={cn("grid items-start gap-4", className)}>
            <div className="relative">
                <img src={image} alt="アップロード画像" className="aspect-square object-cover rounded-md" onClick={image_click_handler} />
                <div className="absolute top-0 left-0 w-full h-full" onClick={image_click_handler}>
                    {/* Annotation */}
                    {annotations && annotations.map((annotation, index) => (
                        <div key={index} className="absolute" style={{ left: `${annotation.x}%`, top: `${annotation.y}%` }}>
                            <div className="absolute w-8 h-8 bg-gray-800 rounded-full" />
                            <div className="absolute w-8 h-8 bg-white rounded-full" />
                        </div>
                    ))}
                </div>
            </div>

            <div>
                <Select>
                    <SelectTrigger className="w-[280px]">
                        <SelectValue placeholder="追加した点を選択" value={selectedAnnotation ? `${selectedAnnotation.x}, ${selectedAnnotation.y}` : ""} />
                    </SelectTrigger>
                    <SelectContent>
                        {annotations && annotations.map((annotation) => (
                            <SelectItem key={`${annotation.x},${annotation.y}`} onClick={() => setSelectedAnnotation(annotation)}>
                                {annotation.x}, {annotation.y}
                            </SelectItem>
                        ))}
                    </SelectContent>
                </Select>
            </div>
            <div>
                <label htmlFor="title" className="block">Title</label>
                <Input type="text" id="title" name="title" className="w-full" onChange={e => {
                    if (selectedAnnotation) {
                        setAnnotations(annotations.map(annotation => {
                            if (annotation.x === selectedAnnotation.x && annotation.y === selectedAnnotation.y) {
                                return { ...annotation, text: e.target.value }
                            }
                            return annotation
                        }))
                    }
                }} />
            </div>
            <Button onClick={e => {
                e.preventDefault()
                setSelectedAnnotation(null)
                setOpen(false)
            }}>Save changes</Button>
        </form>
    )
}
