import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import SeckillDetail from '../views/SeckillDetail.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/seckill/:id',
    name: 'SeckillDetail',
    component: SeckillDetail
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]

export default createRouter({
  history: createWebHistory(),
  routes
}) 