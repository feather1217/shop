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
            <button v-for="val in spec.values" :key="val.value" class="btn "
              :disabled="getStock(spec.name, val.value) === 0"
              :class="[selectedSpecs[spec.name] === val.value ? 'btn-soft btn-primary' : 'bg-gray-100 text-gray-00']"
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
          <button class="btn text-white btn-primary disabled:bg-gray-300" :disabled="getSelected() === 0" @click="addToCart">
            購買
          </button>
        </div>
      </div>
    </div>

    <!-- Toast -->
<div v-if="toastMessage" class="toast toast-start">
  <div class="alert alert-primary shadow-lg">
    <PhCheckCircle :size="24" weight="bold" color="#422ad5" />
    <span>{{ toastMessage }}</span>
  </div>
</div>

  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useProductStore } from '@/stores/productStore'
import { PhPlus, PhCheckCircle, PhWarningCircle } from "@phosphor-icons/vue"
// 例如你在 /products.vue
import { useRoute } from 'vue-router'

const route = useRoute()
const userName = route.query.user
console.log('登入者：', userName)

const productStore = useProductStore()

const selectedSpecs = computed(() => productStore.selectedSpecs)
const selectedProduct = computed(() => productStore.getSelectedProduct())
const quantity = ref(1)
const toastMessage = ref('') // toast 訊息

function selectSpec(name: string, value: string) {
  productStore.selectSpec(name, value)
}

function getStock(specName: string, value: string) {
  return productStore.getStock(specName, value)
}

function getSelected() {
  return productStore.getSelectedProductStock()
}

function checkNum() {
  if (quantity.value > getSelected()) {
    quantity.value = getSelected()
  }
}

function addToCart() {
  const stock = getSelected()
  if (stock === 0 || quantity.value > stock) return

  productStore.decreaseStock(quantity.value)

  // 顯示 toast
  toastMessage.value = '已加入購物車！'
  setTimeout(() => {
    toastMessage.value = ''
  }, 2000)

  console.log('加入購物車', {
    productId: selectedProduct.value?.id,
    specs: selectedSpecs.value,
    quantity: quantity.value,
  })
}

</script>
