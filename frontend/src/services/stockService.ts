import api from './api'
import {
  StockSummary,
  StockCardEntry,
  StockInRequest,
  StockOutRequest,
  BreakDownRequest,
  PackUpRequest,
  TransferRequest,
  StockAdjustRequest,
  StockFilter,
} from '@/types'

export const stockService = {
  getStockSummary: async (filter: StockFilter = {}): Promise<StockSummary[]> => {
    const params = new URLSearchParams()
    if (filter.product_id) params.append('product_id', filter.product_id.toString())
    if (filter.warehouse_id) params.append('warehouse_id', filter.warehouse_id.toString())
    
    const response = await api.get<StockSummary[]>(`/stock?${params.toString()}`)
    return response.data
  },

  getStockCard: async (productId: number, warehouseId: number): Promise<StockCardEntry[]> => {
    const response = await api.get<StockCardEntry[]>(`/stock/card?product_id=${productId}&warehouse_id=${warehouseId}`)
    return response.data
  },

  stockIn: async (data: StockInRequest): Promise<void> => {
    await api.post('/stock/in', data)
  },

  stockOut: async (data: StockOutRequest): Promise<void> => {
    await api.post('/stock/out', data)
  },

  breakDown: async (data: BreakDownRequest): Promise<void> => {
    await api.post('/stock/break', data)
  },

  packUp: async (data: PackUpRequest): Promise<void> => {
    await api.post('/stock/pack', data)
  },

  transfer: async (data: TransferRequest): Promise<void> => {
    await api.post('/stock/transfer', data)
  },

  adjust: async (data: StockAdjustRequest): Promise<void> => {
    await api.post('/stock/adjust', data)
  },
}