import { defineConfig } from 'vite'
import uni from '@dcloudio/vite-plugin-uni'
import { resolve } from 'path'

import AutoImport from 'unplugin-auto-import/vite'

import px2vw from 'postcss-px-to-viewport-8-plugin'
import tailwindcss from 'tailwindcss';
import uniTailwind from '@uni-helper/vite-plugin-uni-tailwind';
import tailwindcssConfig from './tailwind.config.cjs';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    uni(),
    uniTailwind(),
    //自动引入
    AutoImport({
      imports: ['vue'],
      dts: 'src/auto-import.d.ts',
      dirs: ['./src/hooks/**'],
      include: [/.vue$/, /.vue?vue/, /\.[tj]sx?$/],
      resolvers: [],
    }),
  ],
  server: {
    port: 6001,
    host: '0.0.0.0',
    proxy: {
      '^/app': {
        target: 'http://xxxxxxxx',
        ws: true,
        secure: false,
        changeOrigin: true,
        // rewrite: (path) => path.replace(/^\/api\/v1/, '/api/v1'),
        rewrite: (path) => path.replace(/^\/app/, '/app'),
      },
    },
  },
  resolve: {
    alias: [
      {
        find: '@',
        replacement: resolve(__dirname, 'src'),
      },
    ],
  },
  css: {
    postcss: {
      plugins: [
        tailwindcss(),
        //vw
        // px2vw({
        //   viewportWidth: 375, // 视窗的宽度，对应的是我们设计稿的宽度，一般是750
        //   viewportUnit: 'vw',
        //   unitPrecision: 5, // 单位转换后保留的小数位
        //   exclude: [/^node_modules$/], // 设置忽略文件
        //   minPixelValue: 1,
        //   mediaQuery: false,
        // }),
      ],
    },
  },
})
