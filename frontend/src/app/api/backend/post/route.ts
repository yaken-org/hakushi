import { NextRequest, NextResponse } from "next/server";

export async function POST(request: NextRequest) {
    try {
        const body = await request.json();
        console.log(body);

        return NextResponse.json({ message: "OK" }, { status: 200 });
    } catch (e) {
        return NextResponse.json({ message: "Bad Request" }, { status: 400 });
    }
}