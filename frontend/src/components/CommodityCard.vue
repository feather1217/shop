<template>
    <div class="card bg-base-100 space-x-5 shadow-sm">
        <figure>
            <img :src="props.image" alt="" />
        </figure>
        <div class="card-body">
            <h2 class="card-title">{{ props.name }}</h2>
            <p>價格 : {{ props.minPrice }} ~ {{ props.maxPrice }}</p>
            <p></p>           
            <div class="card-actions justify-end">
                <button @click="pBtn()" class="btn ">查看詳情</button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
    import { defineProps, defineEmits } from 'vue'
    import { useRouter } from 'vue-router'
    const router = useRouter()
    const props = defineProps({
        id: Number,
        name: String,
        price: Number,
        minPrice: Number,
        maxPrice: Number,
        image: String,
    })
    console.log('image', props.image); // 印出傳入的圖片路徑
    console.log('image', props.minPrice); // 印出傳入的圖片路徑
    const emit = defineEmits(['view-detail'])
    function pBtn() {
        router.push(`/product/${props.id}`)
    }
</script>