import { createApp } from 'vue';
import App from './App.vue';
import router from './router';

// eslint-disable-next-line import/no-extraneous-dependencies, import/order
import Vue3Toasity from 'vue3-toastify';
// eslint-disable-next-line import/no-extraneous-dependencies
import 'vue3-toastify/dist/index.css';

const app = createApp(App);
app.use(router);
app.use(Vue3Toasity, {
  autoClose: 3000,
  hideProgressBar: true,
  theme: 'colored',
});

app.config.globalProperties.$filters = {
  formatCurrency(value, locale, currency, decimals) {
    return new Intl
      .NumberFormat(locale, {
        style: 'currency',
        currency: currency ?? '',
        currencyDisplay: 'narrowSymbol',
        minimumIntegerDigits: 1,
        minimumFractionDigits: decimals,
        maximumFractionDigits: decimals,
      })
      .format(value * 0.01);
  },
};

app.mount('#app');
