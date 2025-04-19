// stores/user.ts
import { defineStore } from 'pinia'
import User from '../model/user'

export const useUserStore = defineStore('user', {
  state: () => ({
    user: null as User | null, // 預設為 null，當使用者登入後才設置
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
