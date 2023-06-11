<template>
  <h1 class="m-0 mb-5">Add new subscription</h1>

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

        <div class="mt-5">
          <button type="submit" class="button mr-1">Save</button>
          <RouterLink to="/" class="button button-outline">Cancel</RouterLink>
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
import { useRouter } from 'vue-router';
import { currency } from '../stores/settings';
import ValidationErrors from '../components/ValidationErrors.vue';

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

async function submitForm() {
  loading.value = false;

  const formData = { ...form.value };
  formData.price_month *= 100;

  const response = await fetch(`${process.env.VUE_APP_API_URL}/subscriptions`, {
    method: 'POST',
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
      toast('Subscription saved!', { type: 'success' });
    });
  }

  loading.value = true;
}

async function loadCategories() {
  loading.value = false;
  const response = await fetch(`${process.env.VUE_APP_API_URL}/categories`);
  categories.value = await response.json();
  loading.value = true;
}

onMounted(() => {
  document.title = `New subscription - ${process.env.VUE_APP_TITLE_POST}`;
  loadCategories();
});

</script>
