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
              :disabled="getStock(spec.name, val.value) === 0" :class="[
                'btn ',
                selected[spec.name] === val.value
                  ? 'btn-soft btn-info'
                  : 'bg-gray-100 text-gray-00',
                getStock(spec.name, val.value) === 0 ? 'opacity-50 cursor-not-allowed' : ''
              ]" @click="() => selectSpec(spec.name, val.value)">
              {{ val.value }}
            </button>
          </div>
        </div>

        <!-- 數量選擇 -->
        <div class="mt-4">
          <label class="block text-sm font-medium text-gray-700">數量</label>
          <div class="flex items-center gap-2">
            <input type="number" v-model="quantity" :max="displayStock" min="1"
              class="mt-1 block w-20 border rounded p-2" @input="validateQuantity" />
            <!-- 警示 -->
            <div v-if="showAlert" class="alert alert-error alert-soft ">
              <PhWarning :size="20" color="#ff637e" />
              <span>{{ alertMessage }}</span>
            </div>
          </div>

          <span class="text-sm text-gray-500">還剩下 {{ displayStock }} 個</span>


        </div>

        <!-- 下單按鈕 -->
        <div class="mt-6 flex gap-4 items-center">
          <button class="btn text-white bg-blue-400 disabled:bg-gray-300" :disabled="displayStock === 0"
            @click="addToCart">
            購買
          </button>
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
import { PhWarning } from "@phosphor-icons/vue"

const route = useRoute()
const id = Number(route.params.id)
const product = products.find(p => p.id === id)
const showAlert = ref(false)
const alertMessage = ref('')
const previousQuantity = ref(1)



const selected = ref<Record<string, string>>({}) // 儲存使用者選擇的規格
const quantity = ref(0) // 數量

// 使用者選擇規格
function selectSpec(name: string, value: string) {
  selected.value[name] = value
}

// 取得庫存數量
function getStock(specName: string, value: string) {
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

  // 兩個都選了
  if (spec1Value && spec2Value) {
    const match = product.variants.find(
      v => v.specValue1 === spec1Value && v.specValue2 === spec2Value
    )
    return match?.stock ?? 0
  }

  // 只選了第一個規格 
  if (spec1Value && !spec2Value) {
    const filtered = product.variants.filter(v => v.specValue1 === spec1Value)
    return filtered.reduce((sum, v) => sum + v.stock, 0)
  }

  // 只選了第二個規格
  if (!spec1Value && spec2Value) {
    const filtered = product.variants.filter(v => v.specValue2 === spec2Value)
    return filtered.reduce((sum, v) => sum + v.stock, 0)
  }

  // 沒選任何 
  return product.variants.reduce((sum, v) => sum + v.stock, 0)
})

function validateQuantity() {
  if (quantity.value > displayStock.value) {
    if (displayStock.value === 0) {
      alertMessage.value = '已售完'
    } else {
      alertMessage.value = `超出庫存數量（最多可購買 ${displayStock.value} 件）`
    }
    showAlert.value = true
    quantity.value = displayStock.value

    setTimeout(() => (showAlert.value = false), 3000)
  } else if (quantity.value < 1) {
    quantity.value = 1
  } else {
    // 合法範圍才更新 previous
    previousQuantity.value = quantity.value
  }
}




function addToCart() {
  if (displayStock.value === 0) {
    showAlert.value = true
    setTimeout(() => {
      showAlert.value = false
    }, 1000)
    return
  }
  // 模擬加入購物車
  console.log('加入購物車:', {
    productId: product?.id,
    specs: selected.value,
    quantity: quantity.value,
  })
}

</script>
