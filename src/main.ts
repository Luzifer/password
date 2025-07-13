import 'bootstrap/dist/css/bootstrap.css' // Bootstrap 5 Styles
import '@fortawesome/fontawesome-free/css/all.css' // All FA free icons

import { createApp, h } from 'vue'

import i18n from './i18n'

import App from './app.vue'

const app = createApp({
  name: 'Password',

  render() {
    return h(App)
  },
})

app.use(i18n)
app.mount('#app')
