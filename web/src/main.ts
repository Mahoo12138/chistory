import { createApp } from "vue";
import { createI18n } from "vue-i18n";

import "./style.less";
import App from "./App.vue";

import en from "./locales/en.json";
import zhCN from "./locales/zh-CN.json";

const app = createApp(App);

const i18n = createI18n({
  legacy: false,
  locale: "zh-CN",
  fallbackLocale: "en",
  messages: { en, "zh-CN": zhCN },
});

import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const vuetify = createVuetify({
  components,
  directives,
})

app.use(i18n);
app.use(vuetify);

app.mount("#app");
