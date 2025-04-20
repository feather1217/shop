<template>
  <div class="grid grid-cols-4 p-4">
    <CommodityCard
      v-for="product in products"
      :key="product.id"
      :id="product.id"
      :name="product.name"
      :min-price="product.price.min"
      :max-price="product.price.max"
      :image="product.specTypes[0]?.values[0]?.imageUrl || ''"
      :total-stock="product.variants.reduce((sum, v) => sum + v.stock, 0)"
      @view-detail="handleViewDetail"
    />
  </div>
</template>

<script setup lang="ts">
import { onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import CommodityCard from '@/components/productCard.vue'
import { useProductStore } from '@/stores/productStore'
import { useUserStore } from '@/stores/userStore'

const router = useRouter()
const productStore = useProductStore()
const userStore = useUserStore()

// 直接從 Pinia Store 獲取產品資料
const products = productStore.products

function handleViewDetail(id: number) {
  productStore.selectedSpecs['id'] = String(id)
  router.push(`/product/${id}`)
}

// 監聽 products 更新
watch(() => productStore.products, (newProducts) => {
  console.log('監聽到 products 更新:', newProducts)
})

onMounted(() => {
  // 載入商品資料
  productStore.loadProducts()

  // 無登入則跳轉至 LINE 登入
  if (!userStore.user) {
    const clientId = '2007292223'
    const redirectUri = encodeURIComponent('http://localhost:5173/callback')
    const state = crypto.randomUUID()
    const scope = 'profile openid'

    const lineLoginUrl = `https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=${clientId}&redirect_uri=${redirectUri}&state=${state}&scope=${scope}`

    window.location.href = lineLoginUrl
  }
})
</script>

