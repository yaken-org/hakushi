import { NextResponse } from "next/server"

export async function GET() {
    const hello_backend = await fetch(process.env.BACKEND_API_ENDPOINT ?? "http://hakushi-backend.default.svc.cluster.local" + "/health");

    if (hello_backend.ok) {
        const data = await hello_backend.text();
        return NextResponse.json({ message: "Backend is available", data: data });
    } else {
        return NextResponse.json({ message: "Backend is not available" }, { status: 500 });
    }
}