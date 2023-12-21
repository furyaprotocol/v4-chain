const baseConfig = require('./node_modules/@furyaprotocol-indexer/dev/.eslintrc');

module.exports = {
  ...baseConfig,

  // Override the base configuration to set the correct `tsconfigRootDir`.
  parserOptions: {
    ...baseConfig.parserOptions,
    tsconfigRootDir: __dirname,
  },
};
