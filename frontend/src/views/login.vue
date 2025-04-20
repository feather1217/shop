<template>
    <div class="flex items-center justify-center h-screen">
      <button @click="redirectToLineLogin" class="btn btn-success text-lg">使用 LINE 登入</button>
    </div>
  </template>
  
  <script setup lang="ts">
  const redirectToLineLogin = () => {
    const clientId = '2007292223'
    const redirectUri = encodeURIComponent('http://localhost:5173/callback') 
    const state = 'abc123'
    const scope = 'profile openid'
    const responseType = 'code'
    const loginUrl = `https://access.line.me/oauth2/v2.1/authorize?response_type=${responseType}&client_id=${clientId}&redirect_uri=${redirectUri}&state=${state}&scope=${scope}`
    window.location.href = loginUrl
  }
  </script>
  