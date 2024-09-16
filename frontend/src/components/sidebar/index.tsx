import Link from "next/link"
import { Button } from "@/components/ui/button"
import React from "react"
import { Settings, User } from "lucide-react"
import { auth } from "@/lib/auth";
import SignInButton from "../auth/signInButton";
import Image from "next/image";
import { Popover, PopoverTrigger } from "../ui/popover";
import { PopoverContent } from "@radix-ui/react-popover";

export default async function SideBar() {
    const session = await auth();

    return (
        <>
            <nav className="w-full flex md:hidden">
                <div className="w-full flex justify-between items-center p-4 gap-4">
                    <Link href="/" className="text-2xl font-bold">
                        hakushi
                    </Link>
                    <SignInButton />
                </div>
            </nav>
            <div className="hidden h-full md:block">
                <div className="sticky h-full max-h-screen md:flex flex-col border-r">
                    <div className="flex items-center p-4 gap-4">
                        <Link href="/" className="text-2xl font-bold">
                            hakushi
                        </Link>
                        <Button variant="ghost" size="icon" className="ml-auto rounded-full">
                            <Settings className="w-4 h-4" />
                            <span className="sr-only">Settings</span>
                        </Button>
                    </div>
                    <nav className="flex-1 h-full flex flex-col justify-between">
                        <div className="flex flex-col gap-0.5">
                            <Link href="/" className="flex items-center gap-4 p-4 text-lg font-semibold" prefetch={false}>
                                <HomeIcon className="w-5 h-5 text-2xl" />
                                Home
                            </Link>
                            <Link href="/search" className="flex items-center gap-4 p-4 text-lg font-semibold" prefetch={false}>
                            <SearchIcon className="w-4 h-4" />
                                Search
                            </Link>
                            {session?.user.data  && (
                                <Link href={`/${session?.user.providerAccountId}`} className="flex items-center gap-4 p-4 text-lg font-semibold" prefetch={false}>
                                    <User className="w-5 h-5 text-2xl" />
                                    Profile
                                </Link>
                            )}
                        </div>
                        <div className="flex p-4 gap-2 flex-col justify-center items-start w-full">
                            {session?.user.data ? <>
                                <Link href="/post/create" className="w-full">
                                    <Button className="w-full">
                                        投稿する
                                    </Button>
                                </Link>
                                <Popover>
                                    <PopoverTrigger asChild className="w-full flex gap-2 p-2 flex-row items-center justify-start rounded hover:bg-gray-100">
                                        <div>
                                            <Image src={session.user.data.icon_url} width="36" height="36" alt={`${session.user.data.display_name}'s icon`} className="rounded-full" />
                                            <div className="flex flex-col items-start justify-start">
                                                <span className="font-bold">{session.user.data.display_name}</span>
                                                <span className="text-sm text-gray-500">{session.user.data.name}</span>
                                            </div>
                                        </div>
                                    </PopoverTrigger>
                                    <PopoverContent className="w-60 p-2">
                                        <SignInButton className="w-full" />
                                    </PopoverContent>
                                </Popover>
                            </> : session ? (
                                <Link href="/register" className="w-full">
                                    <Button className="w-full">
                                        プロフィールを作成
                                    </Button>
                                </Link>
                            ) : (
                                <SignInButton className="w-full" />
                            )}
                        </div>
                    </nav>
                </div>
            </div>
        </>
    )
}


function HomeIcon(props: React.SVGProps<SVGSVGElement>) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
            <polyline points="9 22 9 12 15 12 15 22" />
        </svg>
    )
}

function SearchIcon(props: React.SVGProps<SVGSVGElement>) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <circle cx="11" cy="11" r="8" />
            <path d="m21 21-4.3-4.3" />
        </svg>
    )
}
