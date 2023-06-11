/* eslint-disable no-return-assign */
import { ref } from 'vue';

export const locale = ref('en-US');
export const currency = ref('USD');
export const decimals = ref(2);

export const setLocale = (value) => locale.value = value;
export const setCurrency = (value) => currency.value = value;
export const setDecimals = (value) => decimals.value = value;

export async function initSettingsStore() {
  const response = await fetch(`${process.env.VUE_APP_API_URL}/settings`);
  const settingValues = await response.json();
  setLocale(settingValues.locale);
  setCurrency(settingValues.currency);
  setDecimals(settingValues.decimals);
}
