<template>
  <div class="box" :class="`bg-${color ?? 'blue'}`">
    {{ title ?? '' }}
    <div class="value">
      {{ value ?? '0' }}
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue';

defineProps({
  title: String,
  value: String,
  color: String,
});
</script>

<style scoped>
.bg-blue { background-color: steelblue; }
.bg-green { background-color: seagreen; }
.bg-red { background-color: firebrick; }
.bg-grey { background-color: slategrey; }
.bg-orange { background-color: darkorange; }

.box {
  /* background-color: slategray; */
  border-radius: 12px;
  color: white;
  text-align: center;
  padding: 2rem .5rem;
  flex: 1 1 100%;
}

.value {
  font-size: 3rem;
  font-weight: bold;
  margin-top: 1rem;
  margin-bottom: 1rem;
}
</style>
