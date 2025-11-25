import api from './api'
import {
  SalesOrder,
  DeliveryOrder,
  PurchaseOrder,
  GoodsReceipt,
  Transfer,
  StockAdjustment
} from '../types'

// Generic document form data types
export interface SalesOrderFormData {
  so_number: string
  date: string
  customer: string
  note: string
}

export interface DeliveryOrderFormData {
  do_number: string
  date: string
  customer: string
  warehouse_id: number
  note: string
}

export interface PurchaseOrderFormData {
  po_number: string
  date: string
  supplier: string
  note: string
}

export interface GoodsReceiptFormData {
  grn_number: string
  date: string
  supplier: string
  warehouse_id: number
  note: string
}

export interface TransferFormData {
  tf_number: string
  from_warehouse_id: number
  to_warehouse_id: number
  date: string
  note: string
}

export interface StockAdjustmentFormData {
  adj_number: string
  date: string
  warehouse_id: number
  reason: string
  note: string
}

export const documentService = {
  // Sales Orders
  getSalesOrders: async (): Promise<SalesOrder[]> => {
    const response = await api.get<SalesOrder[]>('/documents/sales-orders')
    return response.data
  },

  getSalesOrder: async (id: number): Promise<SalesOrder> => {
    const response = await api.get<SalesOrder>(`/documents/sales-orders/${id}`)
    return response.data
  },

  createSalesOrder: async (data: SalesOrderFormData): Promise<SalesOrder> => {
    const response = await api.post<SalesOrder>('/documents/sales-orders', data)
    return response.data
  },

  updateSalesOrder: async (id: number, data: SalesOrderFormData): Promise<SalesOrder> => {
    const response = await api.put<SalesOrder>(`/documents/sales-orders/${id}`, data)
    return response.data
  },

  deleteSalesOrder: async (id: number): Promise<void> => {
    await api.delete(`/documents/sales-orders/${id}`)
  },

  // Delivery Orders
  getDeliveryOrders: async (): Promise<DeliveryOrder[]> => {
    const response = await api.get<DeliveryOrder[]>('/documents/delivery-orders')
    return response.data
  },

  getDeliveryOrder: async (id: number): Promise<DeliveryOrder> => {
    const response = await api.get<DeliveryOrder>(`/documents/delivery-orders/${id}`)
    return response.data
  },

  createDeliveryOrder: async (data: DeliveryOrderFormData): Promise<DeliveryOrder> => {
    const response = await api.post<DeliveryOrder>('/documents/delivery-orders', data)
    return response.data
  },

  updateDeliveryOrder: async (id: number, data: DeliveryOrderFormData): Promise<DeliveryOrder> => {
    const response = await api.put<DeliveryOrder>(`/documents/delivery-orders/${id}`, data)
    return response.data
  },

  deleteDeliveryOrder: async (id: number): Promise<void> => {
    await api.delete(`/documents/delivery-orders/${id}`)
  },

  // Purchase Orders
  getPurchaseOrders: async (): Promise<PurchaseOrder[]> => {
    const response = await api.get<PurchaseOrder[]>('/documents/purchase-orders')
    return response.data
  },

  getPurchaseOrder: async (id: number): Promise<PurchaseOrder> => {
    const response = await api.get<PurchaseOrder>(`/documents/purchase-orders/${id}`)
    return response.data
  },

  createPurchaseOrder: async (data: PurchaseOrderFormData): Promise<PurchaseOrder> => {
    const response = await api.post<PurchaseOrder>('/documents/purchase-orders', data)
    return response.data
  },

  updatePurchaseOrder: async (id: number, data: PurchaseOrderFormData): Promise<PurchaseOrder> => {
    const response = await api.put<PurchaseOrder>(`/documents/purchase-orders/${id}`, data)
    return response.data
  },

  deletePurchaseOrder: async (id: number): Promise<void> => {
    await api.delete(`/documents/purchase-orders/${id}`)
  },

  // Goods Receipts
  getGoodsReceipts: async (): Promise<GoodsReceipt[]> => {
    const response = await api.get<GoodsReceipt[]>('/documents/goods-receipts')
    return response.data
  },

  getGoodsReceipt: async (id: number): Promise<GoodsReceipt> => {
    const response = await api.get<GoodsReceipt>(`/documents/goods-receipts/${id}`)
    return response.data
  },

  createGoodsReceipt: async (data: GoodsReceiptFormData): Promise<GoodsReceipt> => {
    const response = await api.post<GoodsReceipt>('/documents/goods-receipts', data)
    return response.data
  },

  updateGoodsReceipt: async (id: number, data: GoodsReceiptFormData): Promise<GoodsReceipt> => {
    const response = await api.put<GoodsReceipt>(`/documents/goods-receipts/${id}`, data)
    return response.data
  },

  deleteGoodsReceipt: async (id: number): Promise<void> => {
    await api.delete(`/documents/goods-receipts/${id}`)
  },

  // Transfers
  getTransfers: async (): Promise<Transfer[]> => {
    const response = await api.get<Transfer[]>('/documents/transfers')
    return response.data
  },

  getTransfer: async (id: number): Promise<Transfer> => {
    const response = await api.get<Transfer>(`/documents/transfers/${id}`)
    return response.data
  },

  createTransfer: async (data: TransferFormData): Promise<Transfer> => {
    const response = await api.post<Transfer>('/documents/transfers', data)
    return response.data
  },

  updateTransfer: async (id: number, data: TransferFormData): Promise<Transfer> => {
    const response = await api.put<Transfer>(`/documents/transfers/${id}`, data)
    return response.data
  },

  deleteTransfer: async (id: number): Promise<void> => {
    await api.delete(`/documents/transfers/${id}`)
  },

  // Stock Adjustments
  getStockAdjustments: async (): Promise<StockAdjustment[]> => {
    const response = await api.get<StockAdjustment[]>('/documents/stock-adjustments')
    return response.data
  },

  getStockAdjustment: async (id: number): Promise<StockAdjustment> => {
    const response = await api.get<StockAdjustment>(`/documents/stock-adjustments/${id}`)
    return response.data
  },

  createStockAdjustment: async (data: StockAdjustmentFormData): Promise<StockAdjustment> => {
    const response = await api.post<StockAdjustment>('/documents/stock-adjustments', data)
    return response.data
  },

  updateStockAdjustment: async (id: number, data: StockAdjustmentFormData): Promise<StockAdjustment> => {
    const response = await api.put<StockAdjustment>(`/documents/stock-adjustments/${id}`, data)
    return response.data
  },

  deleteStockAdjustment: async (id: number): Promise<void> => {
    await api.delete(`/documents/stock-adjustments/${id}`)
  },
}