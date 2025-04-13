<template>
  <div class="max-w-4xl mx-auto p-8">
    <div v-if="selectedProduct" class="grid grid-cols-2 gap-8">
      <!-- 商品圖片 -->
      <img :src="selectedProduct.specTypes[0]?.values[0]?.imageUrl || ''" class="w-full rounded-lg shadow-md" />

      <!-- 商品資訊 -->
      <div>
        <h1 class="text-3xl font-bold">{{ selectedProduct.name }}</h1>

        <!-- 規格選擇 -->
        <div v-for="spec in selectedProduct.specTypes" :key="spec.name" class="mt-4">
          <label class="block text-sm font-medium text-gray-700 mb-2">{{ spec.name }}</label>
          <div class="space-x-3">
            <button v-for="val in spec.values" :key="val.value" class="btn"
              :disabled="getStock(spec.name, val.value) === 0"
              :class="[selectedSpecs[spec.name] === val.value ? 'btn-soft btn-info' : 'bg-gray-100 text-gray-00']"
              @click="selectSpec(spec.name, val.value)">
              {{ val.value }}
            </button>
          </div>
        </div>

        <!-- 數量選擇 -->
        <div class="mt-4">
          <label class="block text-sm font-medium text-gray-700">數量</label>
          <div class="flex items-center gap-2">
            <input type="number" v-model="quantity" :max="getSelected" min="1"
              class="mt-1 block w-20 border rounded p-2" @input="checkNum" />
          </div>
          <span class="text-sm text-gray-500">還剩下 {{ getSelected() }} 個</span>
        </div>

        <!-- 下單按鈕 -->
        <div class="mt-6 flex gap-4 items-center">
          <button class="btn text-white bg-blue-400 disabled:bg-gray-300" :disabled="getSelected() === 0" @click="addToCart">
            購買
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useProductStore } from '@/stores/productStore'
import { ref, computed } from 'vue'

const productStore = useProductStore()

const selectedSpecs = computed(() => productStore.selectedSpecs)
const selectedProduct = computed(() => productStore.getSelectedProduct())
const quantity = ref(1)
/// 獲取所有商品規格
function selectSpec(name: string, value: string) {
  productStore.selectSpec(name, value)
}
/// 獲取庫存數量
function getStock(specName: string, value: string) {
  return productStore.getStock(specName, value)
}
/// 獲取選中商品的庫存數量
function getSelected() {
  return productStore.getSelectedProductStock()
}
/// 驗證數量輸入
function checkNum() {
  if (quantity.value > getSelected()) {
    quantity.value = getSelected()
  }
}

function addToCart() {
  // 模擬加入購物車的動作
  console.log('加入購物車', {
    productId: selectedProduct.value?.id,
    specs: selectedSpecs.value,
    quantity: quantity.value,
  })
}
</script>
