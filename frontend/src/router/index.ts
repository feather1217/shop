import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '../views/login.vue'
import Product from '../views/product.vue'
import Edit from '../views/edit.vue'
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
      path: '/product',
      name: 'product',
      component: Product,
    },
    {
      path: '/edit',
      name: 'edit',
      component: Edit,
    }

  ],
})

export default router
