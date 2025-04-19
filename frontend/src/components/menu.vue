<template>
  <div class="navbar bg-base-100 border-b border-gray-200">
    <div class="flex-1">
      <!-- 顯示使用者名稱 -->
      <span v-if="userStore.user">歡迎, {{ userStore.user.displayName }}</span>
    </div>

    <div class="flex-none space-x-2">
      <RouterLink to="/" class="btn btn-outline" :class="{ 'btn-active': $route.name === 'home' }">
        商品列表
      </RouterLink>

      <RouterLink to="/edit" class="btn btn-outline" :class="{ 'btn-active': $route.name === 'edit' }">
        規格管理
      </RouterLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/userStore'
import { useRoute } from 'vue-router'

const $route = useRoute()
const userStore = useUserStore()
</script>
