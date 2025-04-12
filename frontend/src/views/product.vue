<template>
  <div class="max-w-4xl mx-auto p-8">
    <div v-if="product" class="grid grid-cols-2 gap-8">
      <!-- 商品圖片 -->
      <img :src="product.specTypes[0]?.values[0]?.imageUrl || ''" class="w-full rounded-lg shadow-md" />

      <!-- 商品資訊 -->
      <div>
        <h1 class="text-3xl font-bold">{{ product.name }}</h1>

        <!-- 規格選擇 -->
        <div v-for="spec in product.specTypes" :key="spec.name" class="mt-4">
          <label class="block text-sm font-medium text-gray-700 mb-2">{{ spec.name }}</label>
          <div class="space-x-3">
            <button v-for="val in spec.values" :key="val.value" class="btn"
              :disabled="getStockForValue(spec.name, val.value) === 0" :class="[
                'btn',
                selected[spec.name] === val.value
                  ? 'bg-blue-400 text-white'
                  : 'bg-gray-200 text-gray-700',
                getStockForValue(spec.name, val.value) === 0 ? 'opacity-50 cursor-not-allowed' : ''
              ]" @click="() => selectSpec(spec.name, val.value)">
              {{ val.value }}
            </button>
          </div>
        </div>

        <!-- 數量選擇 -->
        <div class="mt-4">
          <label class="block text-sm font-medium text-gray-700">數量</label>
          <input type="number" v-model="quantity" :max="displayStock" min="1" class="mt-1 block w-20 border rounded p-2"
            @input="validateQuantity" />
          <span class="text-sm text-gray-500">還剩下 {{ displayStock }} 個</span>
        </div>

        <!-- 下單按鈕 -->
        <div class="mt-6 flex gap-4">
          <button class="btn text-white bg-blue-400" @click="addToCart">購買</button>
        </div>
      </div>
    </div>

    <div v-else>
      <p>找不到商品</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { ref, computed } from 'vue'
import products from '@/mock/productMock'

const route = useRoute()
const id = Number(route.params.id)
const product = products.find(p => p.id === id)

const selected = ref<Record<string, string>>({}) // 儲存使用者選擇的規格
const quantity = ref(1) // 數量

// 使用者選擇規格
function selectSpec(name: string, value: string) {
  selected.value[name] = value
}

function getStockForValue(specName: string, value: string) {
  if (!product) return 0

  const isFirst = product.specTypes[0]?.name === specName
  const isSecond = product.specTypes[1]?.name === specName

  if (isFirst) {
    const filtered = product.variants.filter(v => v.specValue1 === value)
    return filtered.reduce((sum, v) => sum + v.stock, 0)
  } else if (isSecond) {
    const filtered = product.variants.filter(v => v.specValue2 === value)
    return filtered.reduce((sum, v) => sum + v.stock, 0)
  }

  return 0
}


// 顯示的庫存數量
const displayStock = computed(() => {
  if (!product) return 0

  const spec1Name = product.specTypes[0]?.name
  const spec2Name = product.specTypes[1]?.name
  const spec1Value = selected.value[spec1Name]
  const spec2Value = selected.value[spec2Name]

  // 兩個都選了 → 找到對應的 variant
  if (spec1Value && spec2Value) {
    const match = product.variants.find(
      v => v.specValue1 === spec1Value && v.specValue2 === spec2Value
    )
    return match?.stock ?? 0
  }

  // 只選了第一個規格 → 篩選所有符合的 specValue1，加總 stock
  if (spec1Value && !spec2Value) {
    const filtered = product.variants.filter(v => v.specValue1 === spec1Value)
    return filtered.reduce((sum, v) => sum + v.stock, 0)
  }

  // 只選了第二個規格（較少見）→ 加總所有 specValue2 相符的
  if (!spec1Value && spec2Value) {
    const filtered = product.variants.filter(v => v.specValue2 === spec2Value)
    return filtered.reduce((sum, v) => sum + v.stock, 0)
  }

  // 沒選任何 → 顯示所有總庫存
  return product.variants.reduce((sum, v) => sum + v.stock, 0)
})

function validateQuantity() {
  if (quantity.value > displayStock.value) {
    quantity.value = displayStock.value
  }
  if (quantity.value < 1) {
    quantity.value = 1
  }
}


function addToCart() {
  console.log('加入購物車:', {
    productId: product?.id,
    specs: selected.value,
    quantity: quantity.value,
  })
}
</script>
