import { auth } from "@/lib/auth";
import { NextRequest, NextResponse } from "next/server";
import sharp from "sharp";
import { S3Client, PutObjectCommand } from "@aws-sdk/client-s3";

export async function GET() {
    const backend_url = process.env.BACKEND_API_ENDPOINT ?? "http://hakushi-backend.default.svc.cluster.local/api";
    const result = await fetch(`${backend_url}/post`);
    const data = await result.json();
    return NextResponse.json(data);
}

export async function POST(request: NextRequest) {
    try {
        const session = await auth();

        if (!session?.user.data) {
            return NextResponse.json({ message: "Unauthorized" }, { status: 401 });
        }
        const body = await request.json();
        
        const image = body.image; //base64
        const title = body.title;
        const content = body.content;
        const tags = body.tags;
        const annotations = body.annotations;

        const image_id = (new Date()).getTime();

        const backend_url = process.env.BACKEND_API_ENDPOINT ?? "http://hakushi-backend.default.svc.cluster.local/api";
        console.log(backend_url);
        const send_body = JSON.stringify({
            title: title,
            user_account_id: session?.user.providerAccountId,
            content: content,
            tags: tags.map((tag: string) => {
                return { name: tag }
            }),
            annotations: annotations.map((annotation: { x: number, y: number, text: string }) => {
                return { x: annotation.x, y: annotation.y, display_name: annotation.text }
            }),
            image_id: image_id,
        });
        console.log(send_body)
        const result = await fetch(`${backend_url}/post`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: send_body,
        });

        if (!result.ok) {
            return NextResponse.json({ message: "Bad Request" }, { status: 400 });
        }

        const imageBuffer = Buffer.from(image.replace(/^data:image\/\w+;base64,/, ''), 'base64');
        const imageBufferResized = await sharp(imageBuffer)
            .resize(1024, 1024, { fit: 'contain' })
            .webp({
                quality: 80,
            })
            .toBuffer();

        if (!process.env.R2_ACCESS_KEY_ID || !process.env.R2_SECRET_ACCESS_KEY) {
            return NextResponse.json({ message: "Bad Env" }, { status: 500 });
        }

        const S3 = new S3Client({
            region: 'auto',
            endpoint: process.env.R2_API_ENDPOINT,
            credentials: {
                accessKeyId: process.env.R2_ACCESS_KEY_ID,
                secretAccessKey: process.env.R2_SECRET_ACCESS_KEY,
            }
        })
        const _s3_output = await S3.send(
            new PutObjectCommand({
                Bucket: 'hakushi-image',
                Key: `post/${session?.user.providerAccountId}_${image_id}.webp`,
                Body: imageBufferResized,
                ContentEncoding: 'base64',
                ContentType: 'image/webp',
            })
        )
        console.log('R2 uploaded: ' + `post/${session?.user.providerAccountId}_${image_id}.webp`);

        if (_s3_output.$metadata.httpStatusCode !== 200) {
            return { status: 'error', message: 'S3 error' };
        }

        return NextResponse.json({ message: "OK" }, { status: 200 });
    } catch (e) {
        return NextResponse.json({ message: "Bad Request" }, { status: 400 });
    }
}