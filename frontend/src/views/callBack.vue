<template>
    <div class="p-8 text-center">
        <span class="loading loading-spinner text-primary"></span>
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

onMounted(async () => {
    const code = route.query.code //取得Line返回的code
    if (!code || typeof code !== 'string') { 
        console.error('找不到 code')
        return
    }
    try {
        // 用 axios 向後端換使用者資訊
        const res = await axios.get(`http://localhost:8080/api/line/callback?code=${code}`)
        const data = res.data.user
        console.log('用戶資料:', res.data.user)
        console.log('access_token:', res.data.accessToken)
        if (!data) throw new Error('缺少 user 資料')
        const user = new User(data.userId, data.displayName, data.pictureUrl)
        console.log('登入成功:', user)
        userStore.setUser(user)
        router.push('/') //登入成功導回首頁
    } catch (err) {
        console.error('登入錯誤:', err)
    }
})
</script>