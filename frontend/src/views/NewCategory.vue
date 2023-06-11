<template>
  <h1 class="m-0 mb-5">Add new category</h1>

  <div class="grid">
    <div class="col-4">
      <ValidationErrors :errors="validationErrors"></ValidationErrors>

      <form method="POST" @submit.prevent="submitForm">
        <div>
          <label class="label" for="title">Category name</label>
          <input type="text" id="title" v-model.trim="form.title">
        </div>
        <div class="mt-4">
          <button type="submit" class="button mr-1">Save</button>
          <RouterLink :to="{ name: 'CategoriesList' }" class="button button-outline">
            Cancel
          </RouterLink>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { toast } from 'vue3-toastify';
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { setLoading } from '../stores/loader';
import ValidationErrors from '../components/ValidationErrors.vue';

onMounted(() => {
  document.title = `New category - ${process.env.VUE_APP_TITLE_POST}`;
});

const router = useRouter();

const form = ref({
  title: '',
});
const validationErrors = ref([]);

async function submitForm() {
  setLoading(true);
  const formData = { ...form.value };
  const response = await fetch(`${process.env.VUE_APP_API_URL}/categories`, {
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
    router.push({ name: 'CategoriesList', params: {} }).then(() => {
      toast('Category saved!', { type: 'success' });
    });
  }
  setLoading(false);
}

</script>
