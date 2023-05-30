/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  transpilePackages: [
    '@douyinfe/semi-ui',
    '@douyinfe/semi-icons',
    '@douyinfe/semi-illustrations',
  ],
  async rewrites() {
    return {
      fallback: [
        {
          source: '/api/:path*',
          destination: `http:///localhost:8888/api/:path*`,
        },
      ],
    }
  },
}

module.exports = nextConfig
