module.exports = {
  webpack: (config) => {
    const oldConfig = config

    oldConfig.node = {
      fs: 'empty',
    }

    return oldConfig
  },
}
