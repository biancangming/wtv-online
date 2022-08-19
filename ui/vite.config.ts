import { fileURLToPath, URL } from "url";
import Components from "unplugin-vue-components/vite";
import AutoImport from "unplugin-auto-import/vite";
import { AntDesignVueResolver } from "unplugin-vue-components/resolvers";
import type { ConfigEnv, UserConfigExport } from "vite";
import { loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import vueJsx from "@vitejs/plugin-vue-jsx";

// https://vitejs.dev/config/
export default ({ command, mode }: ConfigEnv): UserConfigExport => {
  const root = __dirname;
  const env = loadEnv(mode, root);

  return {
    root,
    base: env.VITE_PUBLIC_PATH,
    build: {
      emptyOutDir: true,
      assetsInlineLimit: 1024 * 128,
      assetsDir: ""
    },
    plugins: [
      vue(),
      vueJsx(),
      Components({
        resolvers: [AntDesignVueResolver()],
      }),
      AutoImport({
        imports: ["vue"],
        dts: "./auto-import.d.ts",
      }),
    ],
    resolve: {
      alias: {
        "@": fileURLToPath(new URL("./src", import.meta.url)),
      },
    },
    server: {
      proxy: {
        "/api": {
          target: "http://127.0.0.1:8888/v1/api/",
          ws: true,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ""),
        },
        "/static": {
          target: "http://127.0.0.1:8888/",
          ws: true,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ""),
        },
      },
    },
  };
};
