import { type DefaultSession } from "next-auth";
import { UserAccount } from "./db/user";

declare module "next-auth" {
    interface Session {
        user: {
            providerAccountId: string, // providerAccountId(sub): google unique user id
            data: UserAccount | null, // database user data
            trigger,
        } & DefaultSession["user"];
    }
}

declare module "next-auth/jwt" {
    interface JWT {
        trigger: "signIn" | "signUp" | "update" | undefined,
    }
}