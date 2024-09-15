import SideBar from "@/components/sidebar";

export default function PageLayout({ children }: Readonly<{ children: React.ReactNode; }>) {
    return (
        <div className="grid md:grid-cols-[240px_1fr] min-h-screen md:border-l dark:border-gray-700">
            <SideBar />
            <main className="flex flex-col w-full min-h-screen md:border-l md:border-gray-200 p-4">
                {children}
            </main>
        </div>
    );
}