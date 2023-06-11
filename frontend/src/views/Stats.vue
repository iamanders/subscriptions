<!-- eslint-disable max-len -->
<template>
  <h1 class="m-0 mb-4">Stats</h1>

  <div v-if="subscriptions.length > 0" class="d-flex stat-boxes-container">
    <StatsBox class="mr-2 ml-2"
      color="grey"
      title="Total subscriptions"
      :value="String(subscriptionsCount)">
    </StatsBox>
    <StatsBox class="mr-2 ml-2"
      color="blue"
      title="Paused subscriptions"
      :value="String(pausedSubscriptionsCount)">
    </StatsBox>
    <StatsBox class="mr-2 ml-2"
      color="orange"
      title="Average subscription cost"
      :value="$filters.formatCurrency(Math.round(averageSubscriptionPrice), locale, currency, decimals)">
    </StatsBox>
  </div>

  <div class="d-flex stat-boxes-container mt-5">
    <StatsBoxBig class="mr-2 ml-2"
      v-if="leastExpensiveSubscription"
      title="Cheapest subscription"
      color="green"
      :value="leastExpensiveSubscription.title"
      :value2="$filters.formatCurrency(Math.round(leastExpensiveSubscription.price_month), locale, currency, decimals) + ' / month'">
    </StatsBoxBig>

    <StatsBoxBig class="mr-2 ml-2"
      v-if="mostExpensiveSubscription"
      title="Most expensive subscription"
      color="red"
      :value="mostExpensiveSubscription.title"
      :value2="$filters.formatCurrency(Math.round(mostExpensiveSubscription.price_month), locale, currency, decimals) + ' / month'">
    </StatsBoxBig>
  </div>

  <div class="panel mt-5 pl-4 pr-4 pt-4 pb-4">
    <div v-if="Object.keys(categorySums).length > 0 && subscriptionsCount > 0">
      <h2 class="m-0 mb-4" style="font-size: 1.2rem;">Category summary</h2>
      <table>
        <thead>
          <tr>
            <th>Category</th>
            <th class="text-right">Price / month</th>
            <th class="text-right">Price / year</th>
            <th>% of total price</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="price, category in categorySums" :key="category">
            <td>{{ category }}</td>
            <td class="text-right">
              {{ $filters.formatCurrency(price, locale, currency, decimals) }}
            </td>
            <td class="text-right">
              {{ $filters.formatCurrency(price * 12, locale, currency, decimals) }}
            </td>
            <td>
              <LineChart
                :percent="Math.round((price / subscriptionsCost) * 100)"
                :value="Math.round((price / subscriptionsCost) * 100) + '%'">
              </LineChart>
            </td>
          </tr>
        </tbody>
        <tfoot>
          <tr>
            <td></td>
            <td class="text-right">
              {{ $filters.formatCurrency(subscriptionsCost, locale, currency, decimals) }}
            </td>
            <td class="text-right">
              {{ $filters.formatCurrency(subscriptionsCost * 12, locale, currency, decimals) }}
            </td>
            <td></td>
          </tr>
        </tfoot>
      </table>
    </div>
  </div>

</template>

<script setup>
/* eslint-disable arrow-body-style */
/* eslint-disable implicit-arrow-linebreak */
/* eslint-disable no-param-reassign */

import { ref, computed, onMounted } from 'vue';
import { setLoading } from '../stores/loader';
import { locale, currency, decimals } from '../stores/settings';
import StatsBox from '../components/StatsBox.vue';
import StatsBoxBig from '../components/StatsBoxBig.vue';
import LineChart from '../components/LineChart.vue';

const loaded = ref(false);
const subscriptions = ref([]);

async function loadData() {
  loaded.value = false;
  const response = await fetch(`${process.env.VUE_APP_API_URL}/subscriptions`);
  const responseData = await response.json();
  subscriptions.value = responseData;
  loaded.value = true;
}

onMounted(async () => {
  document.title = `Stats - ${process.env.VUE_APP_TITLE_POST}`;
  setLoading(true);
  await loadData();
  setLoading(false);
});

const subscriptionsCount = computed(() =>
  subscriptions.value.filter((item) => !item.paused).length);

const pausedSubscriptionsCount = computed(() =>
  subscriptions.value.filter((item) => item.paused).length);

const subscriptionsCost = computed(() =>
  subscriptions.value.reduce((carry, item) => {
    carry += item.paused ? 0 : item.price_month;
    return carry;
  }, 0));

const averageSubscriptionPrice = computed(() =>
  subscriptions.value.reduce((carry, item) => {
    carry += item.paused ? 0 : item.price_month;
    return carry;
  }, 0) / subscriptionsCount.value);

const mostExpensiveSubscription = computed(() => {
  return subscriptions.value.reduce((carry, item) => {
    if (!carry || item.price_month > carry.price_month) {
      carry = item;
    }
    return carry;
  }, null);
});

const leastExpensiveSubscription = computed(() => {
  return subscriptions.value.reduce((carry, item) => {
    if (!carry || item.price_month < carry.price_month) {
      carry = item;
    }
    return carry;
  }, null);
});

const categorySums = computed(() => {
  let categories = subscriptions.value.reduce((carry, item) => {
    if (item.paused) {
      return carry;
    }

    if (item.category?.id !== '') {
      if (carry[item.category.title] === undefined) {
        carry[item.category.title] = 0;
      }
      carry[item.category.title] += item.price_month;
    } else {
      carry.Uncategorized += item.price_month;
    }

    return carry;
  }, { Uncategorized: 0 });

  // Order by key/title
  categories = Object.fromEntries(Object.entries(categories).sort((a, b) =>
    a[0].localeCompare(b[0])));

  return categories;
});

</script>

<style scoped>
.stat-boxes-container {
  margin-left: -.75rem;
  margin-right: -.75rem;
}

table td,
table th {
  width: 25%;
}
</style>
