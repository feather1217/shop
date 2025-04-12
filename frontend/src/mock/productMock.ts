// mock/productMock.ts
import Product from '../model/product'

const mockProducts: Product[] = [
    new Product(
      1,
      '素面 T-shirt',
      [
        {
          name: '顏色',
          values: [
            { value: '白色', imageUrl: '/1.jpg' },
            { value: '黑色', imageUrl: '' }
          ]
        },
        {
          name: '尺寸',
          values: [
            { value: 'S' },
            { value: 'M' },
          ]
        }
      ],
      [
        { specValue1: '白色', specValue2: 'S', stock: 10, price: 100 },
        { specValue1: '白色', specValue2: 'M', stock: 8, price: 199 },
        { specValue1: '黑色', specValue2: 'S', stock: 6, price: 299 },
        { specValue1: '黑色', specValue2: 'M', stock: 5, price: 299 }
      ]
    ),
    new Product(
        2,
        '印花帽T',
        [
          {
            name: '顏色',
            values: [
              { value: '藍色', imageUrl: '/2.jpg' },
              { value: '灰色', imageUrl: '' }
            ]
          },
          {
            name: '尺寸',
            values: [
              { value: 'M' },
              { value: 'L' },
            ]
          }
        ],
        [
          { specValue1: '藍色', specValue2: 'M', stock: 0, price: 350 },
          { specValue1: '藍色', specValue2: 'L', stock: 7, price: 350 },
          { specValue1: '灰色', specValue2: 'M', stock: 5, price: 300 },
          { specValue1: '灰色', specValue2: 'L', stock: 3, price: 300 }
        ]
      )
  ]
  

export default mockProducts
