<template>
  <h1 class="m-0 mb-4">Settings</h1>

  <form v-on:submit.prevent="submitForm">
    <div>
      <label class="label" for="locale">Locale</label>
      <select v-model="form.newLocale" id="locale">
        <option v-for="item in localeList" :value="item.value" :key="item.value">
          {{ item.text }}
        </option>
      </select>
    </div>

    <div class="mt-4">
      <label class="label" for="currency">Currency</label>
      <select v-model="form.newCurrency" id="currency">
        <option v-for="item in currencyList" :value="item.value" :key="item.value">
          {{ item.text }}
        </option>
      </select>
    </div>

    <div class="mt-4">
      <label class="label" for="decimals">Currency decimals</label>
      <select v-model="form.newDecimals" id="decimals">
        <option v-for="item in decimalList" :value="item.value" :key="item.value">
          {{ item.text }}
        </option>
      </select>
    </div>

    <div class="mt-5">
      <input type="submit" class="button" value="Save changes">
    </div>
  </form>

</template>

<script setup>
import { toast } from 'vue3-toastify';
import { ref, watch, onMounted } from 'vue';
import {
  locale, currency, decimals, setLocale, setCurrency, setDecimals,
} from '../stores/settings';
import { setLoading } from '../stores/loader';

const form = ref({
  newLocale: locale.value,
  newCurrency: currency.value,
  newDecimals: decimals.value,
});

const localeList = ref([
  { text: 'en-US', value: 'en-US' },
  { text: 'sv-SE', value: 'sv-SE' },
]);

const currencyList = ref([
  { text: 'Euro', value: 'EUR' },
  { text: 'Swedish krona', value: 'SEK' },
  { text: 'US dollar', value: 'USD' },
]);

const decimalList = ref([
  { text: 'None (0)', value: '0' },
  { text: '1', value: '1' },
  { text: '2', value: '2' },
]);

onMounted(() => {
  document.title = `Settings - ${process.env.VUE_APP_TITLE_POST}`;
});

watch(currency, async () => {
  // Update form when settings store has finished loaded after page reload
  form.value.newLocale = locale.value;
  form.value.newCurrency = currency.value;
  form.value.newDecimals = decimals.value;
});

async function submitForm() {
  setLoading(true);
  setLocale(form.value.newLocale);
  setCurrency(form.value.newCurrency);
  setDecimals(form.value.newDecimals);

  const formData = {
    locale: form.value.newLocale,
    currency: form.value.newCurrency,
    decimals: parseInt(form.value.newDecimals, 10),
  };
  const response = await fetch(`${process.env.VUE_APP_API_URL}/settings`, {
    method: 'PUT',
    body: JSON.stringify(formData),
  });

  if (!response.ok) {
    const message = await response.json();
    console.log('Error', message);
    toast(message, { type: 'error' });
  } else {
    toast('Changes saved!', { type: 'success' });
  }
  setLoading(false);
}

</script>
