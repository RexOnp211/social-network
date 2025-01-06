/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    domains: ['churchthemes.com'],
    remotePatterns: [
      {
        protocol: "http",
        hostname: "api",
        port: "8080",
        pathname: "/**",
      },
      {
        protocol: "http",
        hostname: "localhost",
        port: "8080",
        pathname: "/**",
      },
    ],
  },
};

export default nextConfig;
