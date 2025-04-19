// stores/userStore.ts

import { defineStore } from 'pinia'
import User from '@/model/user'

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null as User | null
  }),
  actions: {
    setUser(user: User) {
      this.user = user
    },
    clearUser() {
      this.user = null
    }
  }
})
