import { defineStore } from 'pinia'
import productMock from '@/mock/productMock'
import Product from '@/model/product'
export const useProductStore = defineStore('product', {
  state: () => ({ // 定義狀態
    products: [] as Product[], // 商品列表，初始為空[]
    selectedSpecs: {} as Record<string, string>, // 用戶所選
  }),

  actions: {
    loadProducts() { //從mock載入資料
      this.products = productMock
      console.log('載入商品資料:', this.products)
    },

    selectSpec(name: string, value: string) { // 紀錄選擇規格
      this.selectedSpecs[name] = value
      console.log('選擇的規格:', this.selectedSpecs[name])
    },

    getStock(specName: string, value: string): number { // 獲取總庫存
      const product = this.getSelectedProduct() // 獲取當前選擇的商品
      if (!product) return 0

      const spec1 = product.specTypes[0]?.name
      const spec2 = product.specTypes[1]?.name

      if (spec1 === specName) {
        return product.variants
          .filter(v => v.specValue1 === value)
          .reduce((sum, v) => sum + v.stock, 0)
      } else if (spec2 === specName) {
        return product.variants
          .filter(v => v.specValue2 === value)
          .reduce((sum, v) => sum + v.stock, 0)
      }

      return 0
    },

    getSelectedProductStock(): number { // 獲取選擇的商品庫存
      const product = this.getSelectedProduct()
      if (!product) return 0

      const spec1 = product.specTypes[0]?.name
      const spec2 = product.specTypes[1]?.name

      const val1 = this.selectedSpecs[spec1]
      const val2 = this.selectedSpecs[spec2]

      const matched = product.variants.find(v => v.specValue1 === val1 && v.specValue2 === val2)
      return matched?.stock ?? 0
    },

    getSelectedProduct(): Product | undefined {// 獲取選擇的商品
      const id = Number(this.selectedSpecs['id']) // 用 id 來找商品
      return this.products.find(p => p.id === id)
    },

    addProduct(product: Product) {
      this.products.push(product)
    },

    updateProduct(product: Product) {
      const index = this.products.findIndex(p => p.id === product.id)
      if (index !== -1) {
        this.products[index] = product
      }
    },

    decreaseStock(count: number) { // 減少庫存
      const selected = this.getSelectedProduct()
      if (!selected) return

      const variant = selected.variants.find(v => {
        return (
          v.specValue1 === this.selectedSpecs[selected.specTypes[0].name] &&
          (selected.specTypes[1] ? v.specValue2 === this.selectedSpecs[selected.specTypes[1].name] : true)
        )
      })

      if (variant && variant.stock >= count) {
        variant.stock -= count
      }
    },




  }

}) 