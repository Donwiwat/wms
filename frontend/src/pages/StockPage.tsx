import React, { useState, useEffect } from 'react'
import { StockSummary as StockSummaryType, Product, Warehouse, StockFilter, StockCardEntry } from '../types'
import StockSummary from '../components/StockSummary'
import StockCard from '../components/StockCard'
import { stockService } from '../services/stockService'
import { productService } from '../services/productService'
import { warehouseService } from '../services/warehouseService'

export default function StockPage() {
  const [stockSummary, setStockSummary] = useState<StockSummaryType[]>([])
  const [products, setProducts] = useState<Product[]>([])
  const [warehouses, setWarehouses] = useState<Warehouse[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [showStockCard, setShowStockCard] = useState(false)
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null)
  const [selectedWarehouse, setSelectedWarehouse] = useState<Warehouse | null>(null)
  const [stockCardEntries, setStockCardEntries] = useState<StockCardEntry[]>([])
  const [stockCardLoading, setStockCardLoading] = useState(false)

  useEffect(() => {
    loadInitialData()
  }, [])

  const loadInitialData = async () => {
    setIsLoading(true)
    try {
      const [productsData, warehousesData, stockData] = await Promise.all([
        productService.getProducts(),
        warehouseService.getWarehouses(),
        stockService.getStockSummary()
      ])

      setProducts(Array.isArray(productsData) ? productsData : [])
      setWarehouses(Array.isArray(warehousesData) ? warehousesData : [])
      setStockSummary(Array.isArray(stockData) ? stockData : [])
    
    } catch (error) {
      console.error('Error loading initial data:', error)
      setStockSummary([])   
    } finally {
      setIsLoading(false)
    }
  }

  const handleFilterChange = async (filter: StockFilter) => {
    setIsLoading(true)
    try {
      const stockData = await stockService.getStockSummary(filter)
      setStockSummary(stockData)
    } catch (error) {
      console.error('Error loading filtered stock data:', error)
    } finally {
      setIsLoading(false)
    }
  }

  const handleViewStockCard = async (productId: number, warehouseId: number) => {
    const product = products.find(p => p.id === productId)
    const warehouse = warehouses.find(w => w.id === warehouseId)
    
    if (!product || !warehouse) {
      console.error('Product or warehouse not found')
      return
    }

    setSelectedProduct(product)
    setSelectedWarehouse(warehouse)
    setStockCardLoading(true)
    setShowStockCard(true)

    try {
      const entries = await stockService.getStockCard(productId, warehouseId)
      setStockCardEntries(entries)
    } catch (error) {
      console.error('Error loading stock card:', error)
      setStockCardEntries([])
    } finally {
      setStockCardLoading(false)
    }
  }

  const handleCloseStockCard = () => {
    setShowStockCard(false)
    setSelectedProduct(null)
    setSelectedWarehouse(null)
    setStockCardEntries([])
  }

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Stock Summary</h1>
        <p className="text-gray-600">View current stock levels across all warehouses</p>
      </div>

      <StockSummary
        stockSummary={stockSummary}
        products={products}
        warehouses={warehouses}
        onFilterChange={handleFilterChange}
        onViewStockCard={handleViewStockCard}
        isLoading={isLoading}
      />

      {showStockCard && selectedProduct && selectedWarehouse && (
        <StockCard
          product={selectedProduct}
          warehouse={selectedWarehouse}
          entries={stockCardEntries}
          onClose={handleCloseStockCard}
          isLoading={stockCardLoading}
        />
      )}
    </div>
  )
}