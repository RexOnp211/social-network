/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    domains: ['churchthemes.com', 'social-network-backend.onrender.com'],
    remotePatterns: [
      {
        protocol: "https",
        hostname: "social-network-backend.onrender.com",
        pathname: "/**",
      },
    ],
  },
};

export default nextConfig;