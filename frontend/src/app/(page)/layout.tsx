import SideBar from "@/components/sidebar";

export default function PageLayout({ children }: Readonly<{ children: React.ReactNode; }>) {
    return (
        <div className="grid md:grid-cols-[240px_1fr] h-full min-h-screen max-h-screen overflow-hidden md:border-l dark:border-gray-700">
            <SideBar />
            <main className="flex flex-col w-full h-full max-h-screen min-h-screen overflow-y-scroll md:border-l md:border-gray-200 p-4">
                {children}
            </main>
        </div>
    );
}