<template>
  <h1 class="m-0 mb-5">Edit category</h1>

  <div class="grid">
    <div class="col-4">
      <ValidationErrors :errors="validationErrors"></ValidationErrors>

      <form method="POST" @submit.prevent="submitForm()">
        <div>
          <label class="label" for="title">Category name</label>
          <input type="text" id="title" v-model.trim="form.title">
        </div>

        <div class="d-flex justify-between mt-5">
          <div>
            <button type="submit" class="button mr-1">Save</button>
            <RouterLink :to="{ name: 'CategoriesList' }" class="button button-outline">
              Cancel
            </RouterLink>
          </div>
          <button type="button" class="button button-danger" @click.prevent="deleteCategory">
            Delete
          </button>
        </div>
      </form>
    </div>
  </div>

</template>

<script setup>
import { toast } from 'vue3-toastify';
import { ref, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { setLoading } from '../stores/loader';
import ValidationErrors from '../components/ValidationErrors.vue';

const route = useRoute();
const router = useRouter();

const form = ref({
  title: '',
});
const validationErrors = ref([]);

async function loadCategory() {
  const response = await fetch(`${process.env.VUE_APP_API_URL}/categories/${route.params.id}`);
  const category = await response.json();
  form.value.title = category.title;
}

async function submitForm() {
  const formData = { ...form.value };
  formData.price_month *= 100;

  const response = await fetch(`${process.env.VUE_APP_API_URL}/categories/${route.params.id}`, {
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
    router.push({ name: 'CategoriesList', params: {} }).then(() => {
      toast('Changes saved!', { type: 'success' });
    });
  }
}

async function deleteCategory() {
  // eslint-disable-next-line no-restricted-globals, no-alert
  if (!confirm('Are you sure?')) {
    return;
  }

  const response = await fetch(`${process.env.VUE_APP_API_URL}/categories/${route.params.id}`, {
    method: 'DELETE',
  });

  if (!response.ok) {
    const message = await response.json();
    console.log('Error', message);
    toast(message, { type: 'error' });
  } else {
    router.push({ name: 'CategoriesList', params: {} }).then(() => {
      toast('Category deleted!', { type: 'success' });
    });
  }
}

onMounted(async () => {
  document.title = `Edit category - ${process.env.VUE_APP_TITLE_POST}`;
  setLoading(true);
  await loadCategory();
  setLoading(false);
});

</script>
