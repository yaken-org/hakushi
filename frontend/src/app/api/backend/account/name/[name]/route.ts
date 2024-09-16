import { get_account_data_by_name } from "@/lib/get_data_utils";
import { NextRequest, NextResponse } from "next/server";

export const revalidate = 10;
export const dynamic = "force-dynamic";

export async function GET(request: NextRequest, params: { name: string }) {
    const result = await get_account_data_by_name(params.name);

    if (!result.is_success && result.status_code !== 404) {
        return NextResponse.json({ message: "Failed to fetch user data" }, { status: 500 });
    } else if (result.status_code === 404) {
        return NextResponse.json({ message: "User not found" }, { status: 200 });
    }

    return NextResponse.json({ is_exists: true, ...result.data });
}