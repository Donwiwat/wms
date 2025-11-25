import { Customer, CustomerFormData, ApiResponse } from '../types'
import api from './api'

export const customerService = {
  async getCustomers(): Promise<Customer[]> {
    const response = await api.get<ApiResponse<Customer[]>>('/customers')
    return response.data.data || []
  },

  async getCustomer(id: number): Promise<Customer> {
    const response = await api.get<ApiResponse<Customer>>(`/customers/${id}`)
    if (!response.data.data) {
      throw new Error('Customer not found')
    }
    return response.data.data
  },

  async createCustomer(data: CustomerFormData): Promise<Customer> {
    const response = await api.post<ApiResponse<Customer>>('/customers', data)
    if (!response.data.data) {
      throw new Error('Failed to create customer')
    }
    return response.data.data
  },

  async updateCustomer(id: number, data: CustomerFormData): Promise<Customer> {
    const response = await api.put<ApiResponse<Customer>>(`/customers/${id}`, data)
    if (!response.data.data) {
      throw new Error('Failed to update customer')
    }
    return response.data.data
  },

  async deleteCustomer(id: number): Promise<void> {
    await api.delete(`/customers/${id}`)
  },

  async searchCustomers(query: string): Promise<Customer[]> {
    const response = await api.get<ApiResponse<Customer[]>>(`/customers/search?q=${encodeURIComponent(query)}`)
    return response.data.data || []
  }
}