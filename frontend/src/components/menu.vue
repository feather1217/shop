<template>
    <div class="navbar bg-base-100 border-b border-gray-200">
      <div class="flex-1">

      </div>
  
      <div class="flex-none space-x-2">
        <!-- 導覽按鈕 -->
        <RouterLink
          to="/"
          class="btn  btn-outline"
          :class="{ 'btn-active': $route.name === 'home' }">
          商品列表
        </RouterLink>
  
        <RouterLink
          to="/edit"
          class="btn btn-outline"
          :class="{ 'btn-active': $route.name === 'edit' }">
          規格管理
        </RouterLink>
  
       

      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { useRoute } from 'vue-router'
  const $route = useRoute()
  </script>
  