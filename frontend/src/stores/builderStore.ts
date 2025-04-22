// stores/specBuilderStore.ts
import { defineStore } from 'pinia'
import type { SpecType, ProductVariant } from '@/model/product'
import Product from '@/model/product'

export const useSpecBuilderStore = defineStore('specBuilder', {
  state: () => ({
    id: Date.now(), //先用 timestamp 作為 id
    name: '',
    specTypes: [] as SpecType[], // 最大兩個規格
    variants: [] as ProductVariant[] // 規格組合
  }),

  actions: {
    // 新增規格類型（最多兩個）
    addSpecType(name: string) {
      if (this.specTypes.length >= 2) return
      this.specTypes.push({ name, values: [] })
    },

    // 新增規格值
    addSpecValue(specIndex: number, value: string, imageUrl?: string) {
        const spec = this.specTypes[specIndex]
        if (spec) {
          spec.values.push({ value, imageUrl })
          this.generateVariants()
        }
      },
      
    // 自動生成規格組合
    generateVariants() {
      const spec1 = this.specTypes[0]
      const spec2 = this.specTypes[1]   
      if (!spec1 || spec1.values.length === 0) {
        this.variants = []
        return
      } 
      const values1 = spec1.values.map(v => v.value)
      const oldVariantMap = new Map<string, ProductVariant>()
    
      // 把現有 variant 存起來，方便後面比對用
      for (const variant of this.variants) {
        const key = `${variant.specValue1}||${variant.specValue2 || ''}`
        oldVariantMap.set(key, variant)
      }
    
      // 單一規格情況
      if (!spec2 || spec2.values.length === 0) {
        this.variants = values1.map(v1 => {
          const key = `${v1}||`
          const old = oldVariantMap.get(key)
          return {
            specValue1: v1,
            stock: old?.stock ?? 0,
            price: old?.price ?? 0,
            imageUrl: spec1.values.find(v => v.value === v1)?.imageUrl || ''
          }
        })
        return
      }
    
      // 雙規格情況
      const values2 = spec2.values.map(v => v.value)
      const newVariants: ProductVariant[] = []
    
      for (const v1 of values1) {
        for (const v2 of values2) {
          const key = `${v1}||${v2}`
          const old = oldVariantMap.get(key)
          newVariants.push({
            specValue1: v1,
            specValue2: v2,
            stock: old?.stock ?? 0,
            price: old?.price ?? 0,
          })
        }
      }
    
      this.variants = newVariants
    },
    

    // 將暫存資料轉換為最終商品 Product 實例
    buildProduct(): Product {
      return new Product(this.id, this.name, this.specTypes, this.variants)
    },

    // 清空規格編輯狀態
    clear() {
      this.specTypes = []
      this.variants = []
      this.name = ''
    }
  }
})
