import api from './api'
import { Warehouse, WarehouseFormData } from '@/types'

export const warehouseService = {
  getWarehouses: async (): Promise<Warehouse[]> => {
    const response = await api.get<Warehouse[]>('/warehouses')
    return response.data
  },

  getWarehouse: async (id: number): Promise<Warehouse> => {
    const response = await api.get<Warehouse>(`/warehouses/${id}`)
    return response.data
  },

  createWarehouse: async (data: WarehouseFormData): Promise<Warehouse> => {
    const response = await api.post<Warehouse>('/warehouses', data)
    return response.data
  },

  updateWarehouse: async (id: number, data: WarehouseFormData): Promise<Warehouse> => {
    const response = await api.put<Warehouse>(`/warehouses/${id}`, data)
    return response.data
  },

  deleteWarehouse: async (id: number): Promise<void> => {
    await api.delete(`/warehouses/${id}`)
  },
}