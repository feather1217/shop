<!-- App.vue -->
<template>
  <Menu></Menu>
  <RouterView />
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import Menu from '@/components/menu.vue'
</script>


