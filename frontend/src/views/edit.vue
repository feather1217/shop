<template>
    <div class="mx-auto p-6">
        <h1 class="text-4xl font-bold mb-8 text-start">商品規格管理</h1>
        <div class="flex flex-col md:flex-row gap-6">
            <!-- 左側：商品清單 -->
            <div class="md:w-1/4 space-y-4">
                <div class="card bg-base-100 shadow p-4">
                    <h2 class="text-lg font-bold mb-2">商品清單</h2>
                    <ul class="menu bg-base-200 rounded-box">
                        <li v-for="product in mockProductsList" :key="product.id" @click="selectProduct(product.id)">
                            <a :class="{ active: selectedProductId === product.id }">
                                {{ product.name }}
                            </a>
                        </li>
                    </ul>
                </div>
                <button @click="addNewProduct" class="btn btn-primary w-full">
                    <PhPlus :size="12" weight="bold" /> 新增商品
                </button>
            </div>

            <!-- 右側：商品編輯 -->
            <div class="md:w-3/4 space-y-6">
                <!-- 商品名稱 -->
                <div class="card bg-base-100 shadow-lg p-4">
                    <div class="form-control">
                        <label class="label font-semibold text-base">商品名稱</label>
                        <input v-model="specBuilder.name" type="text" placeholder="請輸入商品名稱"
                            class="input input-bordered w-full" />
                    </div>
                </div>

                <!-- 輸入規格類型與值 -->
                <div v-for="(spec, index) in specBuilder.specTypes" :key="index" class="border rounded-lg p-4 shadow">
                    <input v-model="spec.name" type="text" placeholder="輸入規格類型名稱"
                        class="input input-sm input-bordered mb-2 w-full" />
                    <div v-for="(value, idx) in spec.values" :key="idx" class="flex items-center gap-4 mb-2">
                        <input v-model="spec.values[idx].value" type="text" placeholder="輸入規格值"
                            class="input input-sm input-bordered w-full" />
                        <input v-if="index === 0" type="file" @change="event => handleImageUpload(event, index, idx)"
                            class="file-input file-input-bordered file-input-sm w-full max-w-xs" />
                    </div>
                    <button @click="addSpecValue(index)" class="btn btn-outline btn-primary btn-sm mt-2">
                        <PhPlus :size="12" weight="bold" /> 新增規格值
                    </button>
                </div>

                <!-- 新增規格類型 -->
                <div>
                    <button @click="addSpecType" class="btn btn-primary"
                        :class="{ 'btn-disabled bg-gray-300 text-gray-500': specBuilder.specTypes.length >= 2 }">
                        <PhPlus :size="12" weight="bold" /> 新增規格類型
                    </button>
                </div>

                <!-- 規格組合表格 -->
                <div class="overflow-x-auto">
                    <h2 class="text-xl font-bold mb-4">規格組合</h2>
                    <table v-if="specBuilder.variants.length > 0" class="table table-zebra w-full">
                        <thead>
                            <tr>
                                <th>{{ specBuilder.specTypes[0]?.name || '規格 1' }}</th>
                                <th>{{ specBuilder.specTypes[1]?.name || '規格 2' }}</th>
                                <th>庫存</th>
                                <th>價格</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(variant, index) in specBuilder.variants" :key="index">
                                <td>{{ variant.specValue1 }}</td>
                                <td>{{ variant.specValue2 }}</td>
                                <td>
                                    <input v-model.number="variant.stock" type="number" min="0"
                                        class="input input-sm input-bordered w-20" />
                                </td>
                                <td>
                                    <input v-model.number="variant.price" type="number" min="0"
                                        class="input input-sm input-bordered w-20" />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <!-- 儲存按鈕 -->
                <div>
                    <button @click="saveProduct" class="btn btn-primary w-full text-lg py-2">
                        {{ selectedProductId ? '儲存商品' : '新增商品' }}
                    </button>
                </div>
            </div>
        </div>

        <!-- Toasts -->
        <div v-if="showSavedToast" class="toast toast-start">
            <div class="alert alert-primary shadow-lg">
                <PhCheckCircle :size="24" weight="bold" color="#422ad5" />
                <span>商品已成功儲存！</span>
            </div>
        </div>

        <div v-if="showErrorToast" class="toast toast-start">
            <div class="alert alert-error shadow-lg">
                <PhWarningCircle :size="24" />
                <span>請為第一個規格的第一個值上傳圖片！</span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useSpecBuilderStore } from '@/stores/builderStore'
import { useProductStore } from '@/stores/productStore'
import { PhPlus, PhCheckCircle, PhWarningCircle } from "@phosphor-icons/vue"

// Store 取用
const specBuilder = useSpecBuilderStore()
const productStore = useProductStore()

// 商品選擇控制
const selectedProductId = ref<number | null>(null)
const mockProductsList = computed(() => productStore.products)

// 監聽商品選擇變化
watch(
    () => specBuilder.specTypes.map(type => type.values.map(v => v.value)),
    () => {
        specBuilder.generateVariants()
    },
    { deep: true }
)

//商品功能
function selectProduct(id?: number) {
    const targetId = id ?? selectedProductId.value
    if (!targetId) return
    const selectedProduct = mockProductsList.value.find(p => p.id === targetId)
    if (selectedProduct) {
        selectedProductId.value = selectedProduct.id
        specBuilder.id = selectedProduct.id
        specBuilder.name = selectedProduct.name
        specBuilder.specTypes = JSON.parse(JSON.stringify(selectedProduct.specTypes))
        specBuilder.variants = JSON.parse(JSON.stringify(selectedProduct.variants))
    }
}

function addNewProduct() {
    specBuilder.id = Date.now()
    specBuilder.name = ''
    specBuilder.specTypes = [{ name: '', values: [{ value: '', imageUrl: '' }] }]
    specBuilder.variants = []
    selectedProductId.value = null
    specBuilder.generateVariants()
}

// 規格操作
function addSpecType() {
    if (specBuilder.specTypes.length >= 2) return
    specBuilder.specTypes.push({ name: '輸入規格類型', values: [{ value: '', imageUrl: '' }] })
    specBuilder.generateVariants()
}

function addSpecValue(specIndex: number) {
    specBuilder.addSpecValue(specIndex, '', '')
    specBuilder.generateVariants()
}

function handleImageUpload(event: Event, specIndex: number, valueIndex: number) {
    const fileInput = event.target as HTMLInputElement
    if (fileInput.files?.[0]) {
        const reader = new FileReader()
        reader.onload = () => {
            const imageUrl = reader.result as string
            specBuilder.specTypes[specIndex].values[valueIndex].imageUrl = imageUrl
        }
        reader.readAsDataURL(fileInput.files[0])
    }
}

// 儲存功能
const showSavedToast = ref(false)
const showErrorToast = ref(false)

function saveProduct() {
  const firstImage = specBuilder.specTypes[0]?.values[0]?.imageUrl;
  if (!firstImage) {
    showErrorToast.value = true;
    setTimeout(() => {
      showErrorToast.value = false;
    }, 3000);
    return;
  }

  const product = specBuilder.buildProduct();

  if (selectedProductId.value) {
    // 如果有 selectedProductId，代表是編輯模式 → 更新商品
    productStore.updateProduct(product);
  } else {
    // 沒有 selectedProductId → 新增商品
    productStore.addProduct(product);
  }

  showSavedToast.value = true;
  setTimeout(() => {
    showSavedToast.value = false;
  }, 2000);

  // 清空編輯區
  specBuilder.name = '';
  specBuilder.specTypes = [{ name: '', values: [{ value: '', imageUrl: undefined }] }];
  specBuilder.variants = [];
  selectedProductId.value = null;
}

</script>
