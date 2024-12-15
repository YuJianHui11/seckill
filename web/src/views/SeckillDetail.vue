<template>
  <div class="seckill-detail">
    <div class="product-info">
      <h1>{{ product.name }}</h1>
      <p class="description">{{ product.description }}</p>
      <div class="price">
        <span class="original">原价: ¥{{ product.price }}</span>
        <span class="seckill">秒杀价: ¥{{ activity.seckillPrice }}</span>
      </div>
      <div class="stock">库存: {{ activity.stock }}</div>
      <div class="countdown">
        <countdown :time="countdownTime" v-if="!started" />
        <div class="status" v-else>秒杀进行中</div>
      </div>
      <button 
        :disabled="!canBuy" 
        @click="handleBuy"
        :class="{ 'btn-disabled': !canBuy }"
      >
        {{ buyButtonText }}
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import Countdown from '@/components/Countdown.vue'
import { getSeckillDetail, createOrder } from '@/api'

export default defineComponent({
  name: 'SeckillDetail',
  components: { Countdown },
  
  setup() {
    const route = useRoute()
    const store = useStore()
    
    const product = ref({})
    const activity = ref({})
    const started = ref(false)
    
    const countdownTime = computed(() => {
      if (!activity.value.startTime) return 0
      return new Date(activity.value.startTime).getTime() - Date.now()
    })
    
    const canBuy = computed(() => {
      return started.value && activity.value.stock > 0
    })
    
    const buyButtonText = computed(() => {
      if (!started.value) return '秒杀未开始'
      if (activity.value.stock <= 0) return '已售罄'
      return '立即抢购'
    })
    
    const loadDetail = async () => {
      try {
        const { data } = await getSeckillDetail(route.params.id)
        product.value = data.product
        activity.value = data.activity
        checkStatus()
      } catch (error) {
        // 错误处理
      }
    }
    
    const checkStatus = () => {
      const now = Date.now()
      const startTime = new Date(activity.value.startTime).getTime()
      const endTime = new Date(activity.value.endTime).getTime()
      started.value = now >= startTime && now <= endTime
    }
    
    const handleBuy = async () => {
      try {
        await createOrder({
          activityId: activity.value.id,
          productId: product.value.id
        })
        // 下单成功处理
      } catch (error) {
        // 错误处理
      }
    }
    
    // 初始化
    loadDetail()
    
    return {
      product,
      activity,
      started,
      countdownTime,
      canBuy,
      buyButtonText,
      handleBuy
    }
  }
})
</script>

<style lang="scss" scoped>
.seckill-detail {
  padding: 20px;
  
  .product-info {
    max-width: 600px;
    margin: 0 auto;
    
    h1 {
      font-size: 24px;
      margin-bottom: 16px;
    }
    
    .description {
      color: #666;
      margin-bottom: 16px;
    }
    
    .price {
      margin-bottom: 16px;
      
      .original {
        color: #999;
        text-decoration: line-through;
        margin-right: 16px;
      }
      
      .seckill {
        color: #f00;
        font-size: 20px;
        font-weight: bold;
      }
    }
    
    .stock {
      margin-bottom: 16px;
    }
    
    .countdown {
      margin-bottom: 24px;
    }
    
    button {
      width: 100%;
      height: 44px;
      border: none;
      border-radius: 4px;
      background: #f00;
      color: #fff;
      font-size: 16px;
      
      &.btn-disabled {
        background: #ccc;
      }
    }
  }
}
</style> 