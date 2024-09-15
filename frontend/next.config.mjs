/** @type {import('next').NextConfig} */
const nextConfig = {
    output: "standalone",
    images: {
        remotePatterns: [
            {
                protocol: "https",
                hostname: "static.hakushi.nenrin.me",
                port: "",
            }
        ]
    }
};

export default nextConfig;
