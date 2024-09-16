import { auth } from "@/lib/auth";
import { NextRequest, NextResponse } from "next/server";

export async function POST(request: NextRequest) {
    try {
        const session = await auth();
        const formData = await request.json();
        console.log(session, formData);

        if (!session || session.user?.data || session.user?.providerAccountId !== formData?.sub) {
            return NextResponse.json({ message: "Unauthorized" }, { status: 401 });
        }


        const backend_url = process.env.BACKEND_API_ENDPOINT ?? "http://hakushi-backend.default.svc.cluster.local/api";
        console.log(`backend_url: ${backend_url}`);
        const response = await fetch(`${backend_url}/account`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                name: formData.name,
                display_name: formData.display_name,
                icon_url: formData.icon_url,
                sub: formData.sub,
            }),
        });

        if (!response.ok) {
            return NextResponse.json({ message: "Failed to create account" }, { status: 500 });
        } else {
            return NextResponse.json({ message: "Account created" }, { status: 201 });
        }
    } catch (error) {
        console.error(error);
        return NextResponse.json({ message: "Failed to create account" }, { status: 500 });
    }
}