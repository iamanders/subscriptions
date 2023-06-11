<template>
  <div class="d-flex justify-between align-center mb-6">
    <h1 class="m-0">Subscriptions</h1>
    <RouterLink :to="{ name: 'NewSubscription' }" class="button">New subscription</RouterLink>
  </div>

  <div v-if="!loaded">
    <p>Loading ..</p>
  </div>
  <div v-else>

    <div class="d-flex justify-end mb-4">
      <div class="checkbox-wrapper">
        <label for="paused">Show paused subscriptions ({{ noPausedSubscriptions }})</label>
        <input type="checkbox" class="ml-1" id="paused" v-model="hidePaused">
      </div>
      <div class="ml-4">
        <select v-model="filterCategory">
          <option value="">- All categories -</option>
          <option v-for="category in categories" :key="category.id">{{ category }}</option>
        </select>
      </div>
    </div>

    <div v-if="filteredSubscriptions.length < 1">
      <p>No subscriptions ..</p>
    </div>
    <div v-else>

      <table>
        <thead>
          <tr>
            <th>Subscription</th>
            <th>Category</th>
            <th class="text-right">Price / month</th>
            <th class="text-right">Price / year</th>
            <th class="text-center">Paused</th>
            <th class="text-right"></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="subscription in filteredSubscriptions" :key="subscription.id"
            :class="{ 'paused-row': subscription.paused }">

            <td :class="{ 'text-striked': subscription.paused }">{{ subscription.title }}</td>
            <td>{{ subscription.category.id ? subscription.category.title : '' }}</td>
            <td :class="{ 'text-striked': subscription.paused }" class="text-right">
              {{ $filters.formatCurrency(
                subscription.price_month,
                locale,
                currency,
                decimals)}}
            </td>
            <td :class="{ 'text-striked': subscription.paused }" class="text-right">
              {{ $filters.formatCurrency(
                subscription.price_month * 12,
                locale,
                currency,
                decimals)}}
            </td>
            <td class="text-center">{{ subscription.paused ? "Yes" : "No" }}</td>
            <td class="text-right">
              <RouterLink :to="{ name: 'EditSubscription', params: { id: subscription.id } }">
                Edit
              </RouterLink>
            </td>
          </tr>
        </tbody>
        <tfoot>
          <tr>
            <td></td>
            <td></td>
            <td class="text-right">
              {{ $filters.formatCurrency(totalPerMonth, locale, currency, decimals) }}
            </td>
            <td class="text-right">
              {{ $filters.formatCurrency(totalPerMonth * 12, locale, currency, decimals) }}
            </td>
            <td></td>
            <td></td>
          </tr>
        </tfoot>
      </table>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { setLoading } from '../stores/loader';
import { locale, currency, decimals } from '../stores/settings';

const loaded = ref(false);
const subscriptions = ref([]);
const hidePaused = ref(false);
const filterCategory = ref('');

async function loadSubscriptions() {
  loaded.value = false;
  const response = await fetch(`${process.env.VUE_APP_API_URL}/subscriptions`);
  const responseData = await response.json();
  subscriptions.value = responseData;
  loaded.value = true;
}

onMounted(async () => {
  document.title = `Subscriptions - ${process.env.VUE_APP_TITLE_POST}`;
  setLoading(true);
  await loadSubscriptions();
  setLoading(false);
});

const filteredSubscriptions = computed(() => {
  let filtered = subscriptions.value;

  // Filter paused
  if (!hidePaused.value) {
    filtered = filtered.filter((s) => !s.paused);
  }

  // Filter category
  if (filterCategory.value !== '') {
    filtered = filtered.filter((s) => s.category.title === filterCategory.value);
  }

  return filtered;
});

const categories = computed(() => {
  const categoryTitles = subscriptions.value.reduce((carry, item) => {
    if (item.category.title !== '') {
      carry.push(item.category.title);
    }
    return carry;
  }, []);
  return [...new Set(categoryTitles)];
});

const totalPerMonth = computed(() =>
  // eslint-disable-next-line implicit-arrow-linebreak, no-return-assign, no-param-reassign, max-len
  filteredSubscriptions.value.reduce((acc, item) => acc += (!item.paused ? item.price_month : 0), 0));

const noPausedSubscriptions = computed(() =>
  // eslint-disable-next-line implicit-arrow-linebreak, no-return-assign, no-param-reassign
  subscriptions.value.reduce((acc, item) => acc += (item.paused ? 1 : 0), 0));

</script>

<style scoped>
.paused-row {
  opacity: .33;
  text-decoration: line-through;
}

table tbody tr:first-child td {
  padding-top: 1rem;
}
</style>
