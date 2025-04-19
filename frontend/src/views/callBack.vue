<template>
    <div class="p-8 text-center">
      <p>登入中，請稍候...</p>
    </div>
  </template>
  
  <script setup lang="ts">
  import axios from 'axios'
  import { onMounted } from 'vue'
  import { useRoute, useRouter } from 'vue-router'
  import { useUserStore } from '@/stores/userStore'
  import User from '@/model/user'
  
  const route = useRoute()
  const router = useRouter()
  const userStore = useUserStore()
  
onMounted(() => {
  const userStr = route.query.user as string
  if (!userStr) return console.error('缺少 user 參數')

  try {
    const parsed = JSON.parse(decodeURIComponent(userStr))
    const user = new User(parsed.userId, parsed.displayName, parsed.pictureUrl)
    userStore.setUser(user)
    router.push('/') // 回首頁
  } catch (e) {
    console.error('解析 user 失敗', e)
  }
})
  </script>
  