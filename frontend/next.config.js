module.exports = {
  webpack: (config) => {
    const oldConfig = config;

    oldConfig.resolve.alias = {
      ...(config.resolve.alias || {}),
      'react-native$': 'react-native-web',
    };
    oldConfig.resolve.extensions = [
      '.web.js',
      '.web.ts',
      '.web.tsx',
      ...config.resolve.extensions,
    ];

    oldConfig.node = {
      fs: 'empty',
    };

    return oldConfig;
  },
};
