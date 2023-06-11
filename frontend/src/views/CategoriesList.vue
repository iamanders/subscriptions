<template>
  <div class="d-flex justify-between align-center mb-5">
    <h1 class="m-0">Categories A-Z</h1>
    <RouterLink :to="{ name: 'NewCategory' }" class="button">New category</RouterLink>
  </div>

  <div v-if="!loaded">
    <p>Loading ..</p>
  </div>
  <div v-else>
    <div v-if="categories.length < 1">
      <p>No categories ..</p>
    </div>
    <div v-else>
      <ul>
        <li v-for="category in categories" :key="category.id">
          <RouterLink :to="{ name: 'EditCategory', params: { id: category.id } }">
            {{ category.title }}
          </RouterLink>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { setLoading } from '../stores/loader';

const loaded = ref(false);
const categories = ref([]);

async function loadCategories() {
  loaded.value = false;
  const response = await fetch(`${process.env.VUE_APP_API_URL}/categories`);
  categories.value = await response.json();
  loaded.value = true;
}

onMounted(async () => {
  document.title = `Categories - ${process.env.VUE_APP_TITLE_POST}`;
  setLoading(true);
  await loadCategories();
  setLoading(false);
});

</script>

<style scoped>
ul {
  padding: 0;
  list-style: none;
}
</style>
