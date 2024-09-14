import { get_user_account_data } from "@/app/api/backend/user_data/[sub]/route";
import NextAuth from "next-auth";
import google from "next-auth/providers/google";

export const { handlers, signIn, signOut, auth } = NextAuth({
    providers: [
        google({
            clientId: process.env.GOOGLE_CLIENT_ID,
            clientSecret: process.env.GOOGLE_CLIENT_SECRET,
        }),
    ],
    secret: process.env.AUTH_SECRET,
    session: {
        strategy: "jwt",
        maxAge: 7 * 24 * 60 * 60,
        updateAge: 24 * 60 * 60,
    },

    callbacks: {
        jwt: async ({ token, user, account, profile, trigger, }) => {
            token.account = account ? account : token.account;
            token.user = user ? user : token.user;
            token.profile = profile ? profile : token.profile;
            token.trigger = trigger ? trigger : token.trigger;
            return token;
        },
        session: async ({ session, token }) => {
            const user_account_data = await get_user_account_data((token.profile as { sub: string }).sub ?? "");

            session.user = {
                ...session.user,
                account: token.account,
                profile: token.profile,
                trigger: token.trigger,
                data: user_account_data.is_success ? user_account_data.data : null,
            };

            return session;
        },

        signIn: async () => {
            return true;
        },
        redirect: async ({ baseUrl }) => {
            return baseUrl;
        }
    },
})
