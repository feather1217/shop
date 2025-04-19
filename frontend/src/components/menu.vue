<template>
  <div class="navbar bg-base-100 border-b border-gray-200">
    <div class="flex-1">
      <!-- 使用者名稱 + 下拉選單 -->
      <div v-if="userStore.user" class="dropdown dropdown-bottom">
        <label tabindex="0" class="btn btn-ghost normal-case text-lg cursor-pointer">
          {{ userStore.user.displayName }}
        </label>
        <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52 mt-2">
          <li><button @click="logout">登出</button></li>
        </ul>
      </div>
    </div>

    <div class="flex-none space-x-2">
      <RouterLink to="/" class="btn btn-outline" :class="{ 'btn-active': $route.name === 'home' }">商品列表</RouterLink>
      <RouterLink to="/edit" class="btn btn-outline" :class="{ 'btn-active': $route.name === 'edit' }">規格管理</RouterLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/userStore'
import { useRoute, useRouter } from 'vue-router'

const userStore = useUserStore()
const $route = useRoute()
const router = useRouter()

const logout = () => {
  userStore.clearUser()
  router.push('/')
}
</script>
