module.exports = {
  webpack: (config) => {
    const oldConfig = config;
    // Fixes npm packages that depend on `fs` module
    oldConfig.node = {
      fs: 'empty',
    };

    return oldConfig;
  },
};
