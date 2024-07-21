//@ts-ignore
import { defineConfig } from 'vite';
import { resolve } from 'path';

export default defineConfig({
  base: '/public/', // Ensure base path is set correctly
  build: {
    manifest: true,
    outDir: 'assets',
    rollupOptions: {
      input: {
        index: resolve(__dirname, 'views/scripts/index.ts'),
        moduleLoader: resolve(__dirname, 'views/scripts/module-loader.ts'),
      },
      output: {
        entryFileNames: (chunk) => {
          // Name the module-loader file without a hash
          if (chunk.name === 'moduleLoader') {
            return 'module-loader.js';
          }
          // Use hashes for other entry files
          return '[name]-[hash].js';
        },
        chunkFileNames: '[name]-[hash].js',
        assetFileNames: '[name]-[hash][extname]'
      },
    },
  },
});
