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
import CommodityCard from '@/components/productCard.vue'
import { useProductStore } from '@/stores/productStore'
import { useRouter } from 'vue-router'

const router = useRouter()
const productStore = useProductStore()
productStore.loadProducts()
const products = productStore.products

function handleViewDetail(id: number) {
  productStore.selectedSpecs['id'] = String(id)
  router.push(`/product/${id}`)
}
</script>
