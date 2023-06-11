<template>
  <div class="line-chart">
    <div class="line" :style="{ width: `${currentPercent}%` }">{{ value }}</div>
  </div>
</template>

<script setup>
// eslint-disable-next-line object-curly-newline
import { defineProps, ref, watch } from 'vue';

// eslint-disable-next-line no-unused-vars
const currentPercent = ref(0);

const props = defineProps({
  percent: Number,
  value: String,
});

watch(
  () => props.percent,
  () => {
    setTimeout(() => {
      currentPercent.value = props.percent;
    }, 10);
  },
  { immediate: true },
);
</script>

<style scoped>
.line-chart {
  background-color: white;
  border-radius: 1rem;
  width: 100%;
  position: relative;
  height: 2rem;
}

.line {
  border-radius: 1rem;
  height: 2rem;
  background-color: tomato;
  color: white;
  font-size: .75rem;
  font-weight: bold;
  line-height: 2rem;
  text-indent: -8px;
  padding-right: 10px;
  padding: 0 1rem;
  width: 1px;
  transition: width 1s ease-out;

  text-align: right;
}
</style>
