const NodePolyfillPlugin = require('node-polyfill-webpack-plugin');

module.exports = function override(config, env) {
    config.resolve.fallback = {
        ...config.resolve.fallback,
        // Include necessary polyfills
        http: require.resolve("stream-http"),
        https: require.resolve("https-browserify"),
        zlib: require.resolve("browserify-zlib"),
        stream: require.resolve("stream-browserify"),
        buffer: require.resolve("buffer/"),
        crypto: require.resolve("crypto-browserify"),
    };

    config.plugins = [
        ...config.plugins,
        new NodePolyfillPlugin()
    ];

    return config;
};
