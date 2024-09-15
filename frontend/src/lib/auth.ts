import { get_user_account_data } from "@/lib/get_data_utils";
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
    trustHost: true,
    callbacks: {
        jwt: async ({ token, user, account, trigger }) => {
            token.user = user ? {
                ...user,
                providerAccountId: account?.providerAccountId,
            } : token.user;
            token.trigger = trigger ? trigger : token.trigger;
            return token;
        },
        session: async ({ session, token }) => {
            const providerAccountId = (token.user as { providerAccountId: string }).providerAccountId;
            const user_account_data = await get_user_account_data(providerAccountId);

            session.user = {
                ...session.user,
                providerAccountId,
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
