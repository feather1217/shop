<template>
    <div class=" flex   items-center justify-center">
        <h1 class=" text-4xl">商品列表</h1>
    </div>
    <div class="grid grid-cols-4 p-4">
    <CommodityCard
      v-for="product in product"
      :key="product.id"
      :id ="product.id"
      :name="product.name"
      :min-price="product.price.min"
      :max-price="product.price.max"
      :image="product.specTypes[0]?.values[0]?.imageUrl || '/1.jpg'"     />
  </div>
</template>

<script setup lang="ts">
    import CommodityCard from '@/components/CommodityCard.vue';
    import product from '@/mock/productMock'

</script>