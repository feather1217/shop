// 商品主體
export default class Product {
    id: number;
    name: string;
    price: {
        min: number;
        max: number;
    };
    specTypes: SpecType[];              // 最多兩種
    variants: ProductVariant[];         // 所有的規格組合

    constructor(
        id: number,
        name: string,
        specTypes: SpecType[],
        variants: ProductVariant[]
    ) {
        this.id = id;
        this.name = name;
        this.specTypes = specTypes;
        this.variants = variants;

        // 自動計算 minPrice 和 maxPrice
        const prices = variants.map((variant) => variant.price);
        this.price = {
            min: Math.min(...prices), // 最小價格
            max: Math.max(...prices), // 最大價格
        };
    }
}

// 規格類型（最多兩種：例如顏色、尺寸）
export interface SpecType {
    name: string;                     // 規格名稱（例如：顏色）
    values: SpecValue[];             // 對應的值
}

// 規格值（第一個規格需要圖片）
export interface SpecValue {
    value: string;                   // 規格值（例如：紅色）
    imageUrl?: string;               // 第一種規格才會填入圖片
}

// 商品的規格組合（每一筆子商品資料）
export interface ProductVariant {
    specValue1: string;
    specValue2?: string;
    stock: number;
    price: number;
}
