import { get_account_data_by_sub } from "@/lib/get_data_utils";
import { NextRequest, NextResponse } from "next/server";

export const revalidate = 10;
export const dynamic = "force-dynamic";

export async function GET(request: NextRequest, params: { sub: string }) {
    const result = await get_account_data_by_sub(params.sub);

    if (!result.is_success && result.status_code !== 404) {
        return NextResponse.json({ message: "Failed to fetch user data" }, { status: 500 });
    } else if (result.status_code === 404) {
        return NextResponse.json({ message: "User not found", is_exists: false }, { status: 200 });
    }

    return NextResponse.json({ is_exists: true, ...result.data });
}
