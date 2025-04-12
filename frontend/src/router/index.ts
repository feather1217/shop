import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '../views/Login.vue'
import Product from '../views/Product.vue'
import Edit from '../views/Edit.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/product/:id',
      name: 'product',
      component: Product,
      props: true,
    },
    {
      path: '/edit',
      name: 'edit',
      component: Edit,
    }

  ],
})

export default router
