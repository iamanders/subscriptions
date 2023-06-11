<template>
  <div class="box" :class="`bg-${color ?? 'blue'}`">
    {{ title ?? '' }}
    <div class="value">
      {{ value ?? '0' }}
    </div>
    <div v-if="value2" class="value2">
      {{ value2 ?? '0' }}
    </div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue';

defineProps({
  title: String,
  value: String,
  value2: String,
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
  border-radius: 12px;
  color: white;
  text-align: center;
  padding: 2rem .5rem;
  flex: 1 1 100%;
}

.value {
  font-size: 2rem;
  font-weight: bold;
  margin-top: 1rem;
  margin-bottom: 1rem;
}
.value2 {
  font-size: 1.2rem;
  font-weight: bold;
}
</style>
