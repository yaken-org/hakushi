import { auth } from "@/lib/auth"

export const config = {
    matcher: ["/((?!api|_next/static|_next/image|favicon.ico).*)"],
}

export default auth((request) => {
    console.log(request.auth, !request.auth?.user.data, request.nextUrl.pathname !== "/register");
    if (request.auth && !request.auth.user.data && request.nextUrl.pathname !== "/register") {
        console.log("redirecting to /register");
        const newURL = new URL("/register", request.nextUrl.origin);
        return Response.redirect(newURL.toString(), 302);
    }

    const has_user_data = request.auth && request.auth.user.data;
    if (has_user_data && request.nextUrl.pathname === "/register") {
        console.log("redirecting to /");
        const newURL = new URL("/", request.nextUrl.origin);
        return Response.redirect(newURL.toString(), 302);
    }
})