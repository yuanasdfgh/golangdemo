import { createSSRApp } from "vue";
import App from "./App.vue";
import './index.css'
export function createApp() {
  const app = createSSRApp(App);
  return {
    app,
  };
}
