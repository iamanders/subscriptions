<template>
  <LoaderBar></LoaderBar>

  <div class="container">
    <div id="nav">
      <router-link :to="{ name: 'Home' }">
        <img src="../public/android-chrome-192x192.png">
      </router-link>
      <router-link :to="{ name: 'Home' }">Subscriptions</router-link>
      <router-link :to="{ name: 'CategoriesList' }">Categories</router-link>
      <router-link :to="{ name: 'Stats' }">Stats</router-link>

      <router-link class="ml-auto" :to="{ name: 'Settings' }">Settings</router-link>
    </div>
  </div>

  <main class="container pt-7 pb-5">
    <router-view/>
  </main>
</template>

<script setup>
import { onMounted } from 'vue';
import './framework.css';
import './app.css';
import LoaderBar from './components/LoaderBar.vue';

// eslint-disable-next-line object-curly-newline
import { initSettingsStore } from './stores/settings';

onMounted(async () => {
  await initSettingsStore();
});

</script>
