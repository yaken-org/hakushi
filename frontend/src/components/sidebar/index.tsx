import Link from "next/link"
import { Button } from "@/components/ui/button"
import React from "react"

export default async function SideBar() {
    return (
        <div className="hidden md:flex flex-col border-r">
            <div className="flex items-center p-4 gap-4">
                <Link href="#" className="text-2xl font-bold">
                    shadcn
                </Link>
                <Button variant="ghost" size="icon" className="ml-auto rounded-full">
                    <SearchIcon className="w-4 h-4" />
                    <span className="sr-only">Search</span>
                </Button>
                <Button variant="ghost" size="icon">
                    <BellIcon className="w-4 h-4" />
                    <span className="sr-only">Notifications</span>
                </Button>
                <Button variant="ghost" size="icon">
                    <MailIcon className="w-4 h-4" />
                    <span className="sr-only">Messages</span>
                </Button>
            </div>
            <nav className="flex-1">
                <div className="flex flex-col gap-0.5">
                    <Link href="#" className="flex items-center gap-4 p-4 text-lg font-semibold" prefetch={false}>
                        <HomeIcon className="w-5 h-5 text-2xl" />
                        Home
                    </Link>
                    <Link href="#" className="flex items-center gap-4 p-4 text-lg font-semibold" prefetch={false}>
                        <HashIcon className="w-5 h-5 text-2xl" />
                        Explore
                    </Link>
                    <Link href="#" className="flex items-center gap-4 p-4 text-lg font-semibold" prefetch={false}>
                        <BellIcon className="w-5 h-5 text-2xl" />
                        Notifications
                    </Link>
                    <Link href="#" className="flex items-center gap-4 p-4 text-lg font-semibold" prefetch={false}>
                        <MailIcon className="w-5 h-5 text-2xl" />
                        Messages
                    </Link>
                </div>
                <Button size="lg" className="mx-4 my-8 justify-start" style={{}}>
                    Upgrade to Pro
                </Button>
                <div className="flex items-center justify-center w-full p-4">
                    <Button variant="outline" className="w-full">
                        Tweet
                    </Button>
                </div>
            </nav>
        </div>
    )
}

function BellIcon(props: React.SVGProps<SVGSVGElement>) {
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
            <path d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9" />
            <path d="M10.3 21a1.94 1.94 0 0 0 3.4 0" />
        </svg>
    )
}

function HashIcon(props: React.SVGProps<SVGSVGElement>) {
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
            <line x1="4" x2="20" y1="9" y2="9" />
            <line x1="4" x2="20" y1="15" y2="15" />
            <line x1="10" x2="8" y1="3" y2="21" />
            <line x1="16" x2="14" y1="3" y2="21" />
        </svg>
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


function MailIcon(props: React.SVGProps<SVGSVGElement>) {
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
            <rect width="20" height="16" x="2" y="4" rx="2" />
            <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
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
