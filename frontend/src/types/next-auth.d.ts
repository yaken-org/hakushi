import { Account, Profile, User, type DefaultSession } from "next-auth";
import { AdapterUser } from "next-auth/adapters";

declare module "next-auth" {
    interface Session {
        user: {
            account,
            profile,
            trigger,
            data,
        } & DefaultSession["user"];
    }
}

declare module "next-auth/jwt" {
    interface JWT {
        account: Account | null;
        user: User | AdapterUser,
        profile: Profile | undefined,
        trigger: "signIn" | "signUp" | "update" | undefined,
    }
}