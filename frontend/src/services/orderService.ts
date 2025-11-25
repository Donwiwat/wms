import { Order, OrderWithDetails, OrderRequest, ApiResponse } from '../types'
import api from './api'

export const orderService = {
  async getOrders(): Promise<OrderWithDetails[]> {
    const response = await api.get<ApiResponse<OrderWithDetails[]>>('/orders')
    return response.data.data || []
  },

  async getOrder(id: number): Promise<OrderWithDetails> {
    const response = await api.get<ApiResponse<OrderWithDetails>>(`/orders/${id}`)
    if (!response.data.data) {
      throw new Error('Order not found')
    }
    return response.data.data
  },

  async createOrder(data: OrderRequest): Promise<Order> {
    const response = await api.post<ApiResponse<Order>>('/orders', data)
    if (!response.data.data) {
      throw new Error('Failed to create order')
    }
    return response.data.data
  },

  async updateOrder(id: number, data: OrderRequest): Promise<Order> {
    const response = await api.put<ApiResponse<Order>>(`/orders/${id}`, data)
    if (!response.data.data) {
      throw new Error('Failed to update order')
    }
    return response.data.data
  },

  async deleteOrder(id: number): Promise<void> {
    await api.delete(`/orders/${id}`)
  },

  async getOrdersByCustomer(customerId: number): Promise<OrderWithDetails[]> {
    const response = await api.get<ApiResponse<OrderWithDetails[]>>(`/orders/customer/${customerId}`)
    return response.data.data || []
  },

  async updateOrderStatus(id: number, status: string): Promise<void> {
    await api.patch(`/orders/${id}/status`, { status })
  }
}