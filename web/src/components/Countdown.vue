<template>
  <div class="countdown">
    <span>{{ formatTime }}</span>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, onUnmounted, computed } from 'vue'

export default defineComponent({
  name: 'Countdown',
  props: {
    time: {
      type: Number,
      required: true
    }
  },
  setup(props) {
    const currentTime = ref(props.time)
    let timer: number

    const formatTime = computed(() => {
      const hours = Math.floor(currentTime.value / 3600000)
      const minutes = Math.floor((currentTime.value % 3600000) / 60000)
      const seconds = Math.floor((currentTime.value % 60000) / 1000)
      
      return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`
    })

    onMounted(() => {
      timer = window.setInterval(() => {
        if (currentTime.value <= 0) {
          clearInterval(timer)
          return
        }
        currentTime.value -= 1000
      }, 1000)
    })

    onUnmounted(() => {
      clearInterval(timer)
    })

    return {
      formatTime
    }
  }
})
</script> 