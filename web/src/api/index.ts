import axios from 'axios'
import { AxiosResponse } from 'axios'

interface OrderData {
  activityId: number
  productId: number
}

interface SeckillDetail {
  product: any
  activity: any
}

interface Product {
  id: number
  name: string
  price: number
  seckillPrice: number
  image: string
}

const api = axios.create({
  baseURL: '/api'
})

export const getProductList = (): Promise<AxiosResponse<Product[]>> => {
  return api.get('/seckill/products')
}

export const getSeckillDetail = (id: string): Promise<AxiosResponse<SeckillDetail>> => {
  return api.get(`/seckill/${id}`)
}

export const createOrder = (data: OrderData): Promise<AxiosResponse<any>> => {
  return api.post('/seckill/order', data)
} 