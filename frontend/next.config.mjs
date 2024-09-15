/** @type {import('next').NextConfig} */
const nextConfig = {
    output: "standalone",
    images: {
        remotePatterns: [
            {
                protocol: "https",
                hostname: "static.hakushi.nenrin.me",
                port: "",
            },
            {
                protocol: "https",
                hostname: "lh3.googleusercontent.com",
                port: "",
            }
        ]
    },
    images: {
        unoptimized: true,
    },
};

export default nextConfig;
