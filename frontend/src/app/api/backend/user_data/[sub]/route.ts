import { get_user_account_data } from "@/lib/get_data_utils";
import { NextRequest, NextResponse } from "next/server";

export const revalidate = 10;

export function GET(request: NextRequest, params: { sub: string }) {
    return NextResponse.json(get_user_account_data(params.sub));
}
