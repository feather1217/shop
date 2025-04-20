import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '../views/Login.vue'
import Product from '../views/Product.vue'
import Edit from '../views/edit.vue'
import LineCallback from '../views/callBack.vue'
import { useUserStore } from '../stores/userStore'

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
    },
    {
      path: '/callback', // ✅ 這就是 LINE 登入後會跳來的頁面
      name: 'callback',
      component: LineCallback
    }

  ],
})

router.beforeEach((to, from, next) => {
  const userStore = useUserStore()

  // 排除 callback 頁面，避免登入死循環
  if (to.path === '/callback') {
    return next()
  }

  // 未登入則導向 LINE Login
  if (!userStore.user) {
    const clientId = '2007292223'
    const redirectUri = encodeURIComponent('http://localhost:5173/callback')
    const state = crypto.randomUUID()
    const scope = 'profile openid'
    const lineLoginUrl = `https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=${clientId}&redirect_uri=${redirectUri}&state=${state}&scope=${scope}`

    window.location.href = lineLoginUrl
    return // 中斷導航
  }

  // 有登入就繼續
  next()
})

export default router
