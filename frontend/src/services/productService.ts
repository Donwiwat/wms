import api from './api'
import { Product, ProductFormData } from '@/types'

export const productService = {
  getProducts: async (): Promise<Product[]> => {
    const response = await api.get<Product[]>('/products')
    return response.data
  },

  getProduct: async (id: number): Promise<Product> => {
    const response = await api.get<Product>(`/products/${id}`)
    return response.data
  },

  createProduct: async (data: ProductFormData): Promise<Product> => {
    const response = await api.post<Product>('/products', data)
    return response.data
  },

  updateProduct: async (id: number, data: ProductFormData): Promise<Product> => {
    const response = await api.put<Product>(`/products/${id}`, data)
    return response.data
  },

  deleteProduct: async (id: number): Promise<void> => {
    await api.delete(`/products/${id}`)
  },

  searchProducts: async (query: string): Promise<Product[]> => {
    const response = await api.get<Product[]>(`/products?search=${encodeURIComponent(query)}`)
    return response.data
  },
}