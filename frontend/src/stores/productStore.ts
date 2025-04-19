// stores/productStore.ts
import { defineStore } from 'pinia'
import productMock from '@/mock/productMock'
import Product from '@/model/product'

export const useProductStore = defineStore('product', {
  state: () => ({
    products: [] as Product[],
    selectedSpecs: {} as Record<string, string>,
  }),

  actions: {
    loadProducts() { //從mock載入資料
      this.products = productMock
    },

    selectSpec(name: string, value: string) { // 選擇規格
      this.selectedSpecs[name] = value
    },

    getStock(specName: string, value: string): number { // 獲取庫存
      const product = this.getSelectedProduct()
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

    getSelectedProduct(): Product | undefined {
      const id = Number(this.selectedSpecs['id']) // 假設用 id 來找商品
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
      
      getProductById(id: number) {
        return this.products.find(p => p.id === id)
      },

      decreaseStock(count: number) {
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