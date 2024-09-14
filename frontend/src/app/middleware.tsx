import { auth } from "@/lib/auth"

export const config = {
    matcher: ["/((?!api|_next/static|_next/image|favicon.ico).*)"],
}

export default auth((request) => {
    const require_registered_page = [
        "/post/create",
        "/post/[id]/edit",
        "/user/[id]/edit",
    ]
    if (request.auth && !request.auth.user.data && require_registered_page.includes(request.nextUrl.pathname)) {
        const newURL = new URL("/register", request.nextUrl.origin);
        return Response.redirect(newURL.toString(), 302);
    }
})