<template>
    <div class="bar" :class="{ 'active': isLoading }"></div>
</template>

<script setup>
import { isLoading } from '../stores/loader';
</script>

<style scoped>
    .bar {
        background-color: #eee;
        height: 6px;
    }
    .bar.active {
        background-color: #eee;
        animation: pulse .66s infinite;
    }

    @keyframes pulse {
        0% {
            background-color: #eee;
        }
        70% {
            background-color: #2c3e50;
        }
        100% {
            background-color: #eee;
        }
    }
</style>
