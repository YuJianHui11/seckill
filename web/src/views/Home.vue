<template>
  <div class="home">
    <h1>秒杀商品列表</h1>
    <div class="product-list">
      <el-card v-for="item in products" :key="item.id" class="product-item">
        <img :src="item.image" class="product-image">
        <h2>{{ item.name }}</h2>
        <p class="price">
          <span class="original">原价: ¥{{ item.price }}</span>
          <span class="seckill">秒杀价: ¥{{ item.seckillPrice }}</span>
        </p>
        <el-button type="primary" @click="goToDetail(item.id)">查看详情</el-button>
      </el-card>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getProductList } from '@/api'

export default defineComponent({
  name: 'Home',
  setup() {
    const router = useRouter()
    const products = ref([])

    const loadProducts = async () => {
      try {
        const { data } = await getProductList()
        products.value = data
      } catch (error) {
        console.error('Failed to load products:', error)
      }
    }

    const goToDetail = (id: number) => {
      router.push(`/seckill/${id}`)
    }

    onMounted(() => {
      loadProducts()
    })

    return {
      products,
      goToDetail
    }
  }
})
</script>

<style lang="scss" scoped>
.home {
  padding: 20px;

  .product-list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 20px;
    margin-top: 20px;
  }

  .product-item {
    .product-image {
      width: 100%;
      height: 200px;
      object-fit: cover;
    }

    h2 {
      margin: 10px 0;
      font-size: 18px;
    }

    .price {
      margin: 10px 0;

      .original {
        color: #999;
        text-decoration: line-through;
        margin-right: 10px;
      }

      .seckill {
        color: #f00;
        font-weight: bold;
      }
    }
  }
}
</style> 