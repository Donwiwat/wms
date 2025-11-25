import React, { useState, useEffect } from 'react'
import {
  Product,
  Warehouse,
  Stock,
  StockInRequest,
  StockOutRequest,
  TransferRequest,
  BreakDownRequest,
  PackUpRequest,
  StockAdjustRequest
} from '../types'
import StockInForm from '../components/StockInForm'
import StockOutForm from '../components/StockOutForm'
import TransferForm from '../components/TransferForm'
import BreakDownForm from '../components/BreakDownForm'
import PackUpForm from '../components/PackUpForm'
import StockAdjustForm from '../components/StockAdjustForm'
import { productService } from '../services/productService'
import { warehouseService } from '../services/warehouseService'
import { stockService } from '../services/stockService'

type OperationType = 'stock-in' | 'stock-out' | 'transfer' | 'break-down' | 'pack-up' | 'adjustment' | null

export default function StockOperationsPage() {
  const [currentOperation, setCurrentOperation] = useState<OperationType>(null)
  const [products, setProducts] = useState<Product[]>([])
  const [warehouses, setWarehouses] = useState<Warehouse[]>([])
  const [isLoading, setIsLoading] = useState(false)

  useEffect(() => {
    loadData()
  }, [])

  const loadData = async () => {
    setIsLoading(true)
    try {
      const [productsData, warehousesData] = await Promise.all([
        productService.getProducts(),
        warehouseService.getWarehouses()
      ])
      setProducts(productsData)
      setWarehouses(warehousesData)
    } catch (error) {
      console.error('Error loading data:', error)
    } finally {
      setIsLoading(false)
    }
  }

  const handleStockIn = async (data: StockInRequest) => {
    setIsLoading(true)
    try {
      await stockService.stockIn(data)
      alert('Stock in operation completed successfully!')
      setCurrentOperation(null)
    } catch (error) {
      console.error('Error performing stock in:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const handleStockOut = async (data: StockOutRequest) => {
    setIsLoading(true)
    try {
      await stockService.stockOut(data)
      alert('Stock out operation completed successfully!')
      setCurrentOperation(null)
    } catch (error) {
      console.error('Error performing stock out:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const handleTransfer = async (data: TransferRequest) => {
    setIsLoading(true)
    try {
      await stockService.transfer(data)
      alert('Transfer operation completed successfully!')
      setCurrentOperation(null)
    } catch (error) {
      console.error('Error performing transfer:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const handleBreakDown = async (data: BreakDownRequest) => {
    setIsLoading(true)
    try {
      await stockService.breakDown(data)
      alert('Break down operation completed successfully!')
      setCurrentOperation(null)
    } catch (error) {
      console.error('Error performing break down:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const handlePackUp = async (data: PackUpRequest) => {
    setIsLoading(true)
    try {
      await stockService.packUp(data)
      alert('Pack up operation completed successfully!')
      setCurrentOperation(null)
    } catch (error) {
      console.error('Error performing pack up:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const handleStockAdjust = async (data: StockAdjustRequest) => {
    setIsLoading(true)
    try {
      await stockService.adjust(data)
      alert('Stock adjustment completed successfully!')
      setCurrentOperation(null)
    } catch (error) {
      console.error('Error performing stock adjustment:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const handleGetCurrentStock = async (productId: number, warehouseId: number): Promise<Stock | null> => {
    try {
      const stockSummary = await stockService.getStockSummary({ product_id: productId, warehouse_id: warehouseId })
      const summary = stockSummary.find(s => s.product_id === productId && s.warehouse_id === warehouseId)
      
      if (summary) {
        // Convert StockSummary to Stock format
        return {
          id: 0, // Not available in summary
          product_id: summary.product_id,
          warehouse_id: summary.warehouse_id,
          remain1: summary.remain1,
          remain2: summary.remain2,
          total_remain: summary.total_remain,
          updated_at: summary.updated_at
        }
      }
      
      return null
    } catch (error) {
      console.error('Error getting current stock:', error)
      return null
    }
  }

  const handleCancel = () => {
    setCurrentOperation(null)
  }

  // Render current operation form
  if (currentOperation) {
    switch (currentOperation) {
      case 'stock-in':
        return (
          <div className="space-y-6">
            <StockInForm
              products={products}
              warehouses={warehouses}
              onSubmit={handleStockIn}
              onCancel={handleCancel}
              isLoading={isLoading}
            />
          </div>
        )
      case 'stock-out':
        return (
          <div className="space-y-6">
            <StockOutForm
              products={products}
              warehouses={warehouses}
              onSubmit={handleStockOut}
              onCancel={handleCancel}
              isLoading={isLoading}
            />
          </div>
        )
      case 'transfer':
        return (
          <div className="space-y-6">
            <TransferForm
              products={products}
              warehouses={warehouses}
              onSubmit={handleTransfer}
              onCancel={handleCancel}
              isLoading={isLoading}
            />
          </div>
        )
      case 'break-down':
        return (
          <div className="space-y-6">
            <BreakDownForm
              products={products}
              warehouses={warehouses}
              onSubmit={handleBreakDown}
              onCancel={handleCancel}
              isLoading={isLoading}
            />
          </div>
        )
      case 'pack-up':
        return (
          <div className="space-y-6">
            <PackUpForm
              products={products}
              warehouses={warehouses}
              onSubmit={handlePackUp}
              onCancel={handleCancel}
              isLoading={isLoading}
            />
          </div>
        )
      case 'adjustment':
        return (
          <div className="space-y-6">
            <StockAdjustForm
              products={products}
              warehouses={warehouses}
              onSubmit={handleStockAdjust}
              onCancel={handleCancel}
              onGetCurrentStock={handleGetCurrentStock}
              isLoading={isLoading}
            />
          </div>
        )
    }
  }

  // Render operation selection
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Stock Operations</h1>
        <p className="text-gray-600">Perform stock in, out, transfers, and adjustments</p>
      </div>

      {isLoading && (
        <div className="text-center py-4">
          <p className="text-gray-500">Loading...</p>
        </div>
      )}

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div className="card hover:shadow-lg transition-shadow cursor-pointer" onClick={() => setCurrentOperation('stock-in')}>
          <div className="card-header">
            <h3 className="text-lg font-medium text-gray-900 flex items-center">
              <span className="text-green-500 mr-2">📥</span>
              Stock In
            </h3>
          </div>
          <div className="card-content">
            <p className="text-gray-600 mb-4">Add stock to warehouse from purchases or receipts</p>
            <button className="btn-primary w-full" disabled={isLoading}>
              Stock In Form
            </button>
          </div>
        </div>

        <div className="card hover:shadow-lg transition-shadow cursor-pointer" onClick={() => setCurrentOperation('stock-out')}>
          <div className="card-header">
            <h3 className="text-lg font-medium text-gray-900 flex items-center">
              <span className="text-red-500 mr-2">📤</span>
              Stock Out
            </h3>
          </div>
          <div className="card-content">
            <p className="text-gray-600 mb-4">Remove stock from warehouse for sales or deliveries</p>
            <button className="btn-primary w-full" disabled={isLoading}>
              Stock Out Form
            </button>
          </div>
        </div>

        <div className="card hover:shadow-lg transition-shadow cursor-pointer" onClick={() => setCurrentOperation('transfer')}>
          <div className="card-header">
            <h3 className="text-lg font-medium text-gray-900 flex items-center">
              <span className="text-blue-500 mr-2">🔄</span>
              Transfer
            </h3>
          </div>
          <div className="card-content">
            <p className="text-gray-600 mb-4">Transfer stock between warehouses</p>
            <button className="btn-primary w-full" disabled={isLoading}>
              Transfer Form
            </button>
          </div>
        </div>

        <div className="card hover:shadow-lg transition-shadow cursor-pointer" onClick={() => setCurrentOperation('break-down')}>
          <div className="card-header">
            <h3 className="text-lg font-medium text-gray-900 flex items-center">
              <span className="text-purple-500 mr-2">📦</span>
              Break Down
            </h3>
          </div>
          <div className="card-content">
            <p className="text-gray-600 mb-4">Convert from larger units to smaller units</p>
            <button className="btn-primary w-full" disabled={isLoading}>
              Break Down Form
            </button>
          </div>
        </div>

        <div className="card hover:shadow-lg transition-shadow cursor-pointer" onClick={() => setCurrentOperation('pack-up')}>
          <div className="card-header">
            <h3 className="text-lg font-medium text-gray-900 flex items-center">
              <span className="text-indigo-500 mr-2">📋</span>
              Pack Up
            </h3>
          </div>
          <div className="card-content">
            <p className="text-gray-600 mb-4">Convert from smaller units to larger units</p>
            <button className="btn-primary w-full" disabled={isLoading}>
              Pack Up Form
            </button>
          </div>
        </div>

        <div className="card hover:shadow-lg transition-shadow cursor-pointer" onClick={() => setCurrentOperation('adjustment')}>
          <div className="card-header">
            <h3 className="text-lg font-medium text-gray-900 flex items-center">
              <span className="text-orange-500 mr-2">⚖️</span>
              Adjustment
            </h3>
          </div>
          <div className="card-content">
            <p className="text-gray-600 mb-4">Adjust stock levels for corrections</p>
            <button className="btn-primary w-full" disabled={isLoading}>
              Adjustment Form
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}