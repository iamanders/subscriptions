<template>
  <h1 class="m-0 mb-5">Edit subscription</h1>

  <div class="grid">
    <div class="col-4">
      <ValidationErrors :errors="validationErrors"></ValidationErrors>

      <form method="POST" @submit.prevent="submitForm">
        <div>
          <label class="label" for="title">Subscription name</label>
          <input type="text" id="title" v-model.trim="form.title">
        </div>

        <div class="mt-2">
          <label class="label" for="price_month">Price/month ({{ currency }})</label>
          <input type="number" id="price_month" v-model.number.trim="form.price_month">
        </div>

        <div class="mt-2">
          <label class="label" for="category">Category</label>
          <input
            type="text"
            id="category"
            v-model.trim="form.category"
            autocomplete="off"
            list="categories">
        </div>

        <div class="mt-4 checkbox-wrapper">
          <input type="checkbox" id="paused" v-model="form.paused">
          <label class="label" for="paused">Paused</label>
        </div>

        <div class="d-flex justify-between mt-5">
          <div>
            <button type="submit" class="button mr-1">Save</button>
            <RouterLink to="/" class="button button-outline">Cancel</RouterLink>
          </div>
          <button type="button" class="button button-danger" @click.prevent="deleteSubscription">
            Delete
          </button>
        </div>

      </form>
    </div>
  </div>

  <datalist id="categories">
    <option v-for="category in categories" :key="category.id">{{ category.title }}</option>
  </datalist>
</template>

<script setup>
import { toast } from 'vue3-toastify';
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { currency } from '../stores/settings';
import ValidationErrors from '../components/ValidationErrors.vue';

const route = useRoute();
const router = useRouter();

const form = ref({
  title: '',
  price_month: '',
  category: '',
  paused: false,
});
const validationErrors = ref([]);
const categories = ref([]);

const loading = ref(false);

async function loadSubscription() {
  const response = await fetch(`${process.env.VUE_APP_API_URL}/subscriptions/${route.params.id}`);
  const data = await response.json();
  form.value.title = data.title;
  form.value.price_month = data.price_month * 0.01;
  form.value.category = data.category.title;
  form.value.paused = data.paused;
  loading.value = true;
}

async function loadCategories() {
  loading.value = false;
  const response = await fetch(`${process.env.VUE_APP_API_URL}/categories`);
  categories.value = await response.json();
  loading.value = true;
}

async function submitForm() {
  const formData = { ...form.value };
  formData.price_month *= 100;

  const response = await fetch(`${process.env.VUE_APP_API_URL}/subscriptions/${route.params.id}`, {
    method: 'PUT',
    body: JSON.stringify(formData),
  });

  if (!response.ok) {
    if (response.status === 422) {
      window.scrollTo({ top: 0, behavior: 'smooth' });
      validationErrors.value = await response.json();
    } else {
      const message = await response.json();
      console.log('Error', message);
      toast(message, { type: 'error' });
    }
  } else {
    router.push({ name: 'Home', params: {} }).then(() => {
      toast('Changes saved!', { type: 'success' });
    });
  }
}

async function deleteSubscription() {
  // eslint-disable-next-line no-restricted-globals, no-alert
  if (!confirm('Are you sure?')) {
    return;
  }

  const response = await fetch(`${process.env.VUE_APP_API_URL}/subscriptions/${route.params.id}`, {
    method: 'DELETE',
  });

  if (!response.ok) {
    const message = await response.json();
    console.log('Error', message);
    toast(message, { type: 'error' });
  } else {
    router.push({ name: 'Home', params: {} }).then(() => {
      toast('Subscription deleted!', { type: 'success' });
    });
  }
}

onMounted(() => {
  document.title = `Edit subscription - ${process.env.VUE_APP_TITLE_POST}`;
  loadSubscription();
  loadCategories();
});

</script>
