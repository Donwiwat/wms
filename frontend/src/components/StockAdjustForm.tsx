import React, { useState, useEffect } from 'react'
import { StockAdjustRequest, Product, Warehouse, Stock } from '../types'

interface StockAdjustFormProps {
  products: Product[]
  warehouses: Warehouse[]
  onSubmit: (data: StockAdjustRequest) => Promise<void>
  onCancel: () => void
  onGetCurrentStock?: (productId: number, warehouseId: number) => Promise<Stock | null>
  isLoading?: boolean
}

export default function StockAdjustForm({ 
  products, 
  warehouses, 
  onSubmit, 
  onCancel, 
  onGetCurrentStock,
  isLoading = false 
}: StockAdjustFormProps) {
  const [formData, setFormData] = useState<StockAdjustRequest>({
    product_id: 0,
    warehouse_id: 0,
    new_remain1: 0,
    new_remain2: 0,
    reason: '',
    note: ''
  })

  const [errors, setErrors] = useState<Partial<Record<keyof StockAdjustRequest, string>>>({})
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null)
  const [currentStock, setCurrentStock] = useState<Stock | null>(null)
  const [loadingStock, setLoadingStock] = useState(false)

  useEffect(() => {
    if (formData.product_id > 0) {
      const product = products.find(p => p.id === formData.product_id)
      setSelectedProduct(product || null)
    }
  }, [formData.product_id, products])

  useEffect(() => {
    const fetchCurrentStock = async () => {
      if (formData.product_id > 0 && formData.warehouse_id > 0 && onGetCurrentStock) {
        setLoadingStock(true)
        try {
          const stock = await onGetCurrentStock(formData.product_id, formData.warehouse_id)
          setCurrentStock(stock)
          if (stock) {
            setFormData(prev => ({
              ...prev,
              new_remain1: stock.remain1,
              new_remain2: stock.remain2
            }))
          }
        } catch (error) {
          console.error('Error fetching current stock:', error)
          setCurrentStock(null)
        } finally {
          setLoadingStock(false)
        }
      } else {
        setCurrentStock(null)
      }
    }

    fetchCurrentStock()
  }, [formData.product_id, formData.warehouse_id, onGetCurrentStock])

  const validateForm = (): boolean => {
    const newErrors: Partial<Record<keyof StockAdjustRequest, string>> = {}

    if (formData.product_id <= 0) {
      newErrors.product_id = 'Please select a product'
    }

    if (formData.warehouse_id <= 0) {
      newErrors.warehouse_id = 'Please select a warehouse'
    }

    if (formData.new_remain1 < 0) {
      newErrors.new_remain1 = 'Primary unit quantity cannot be negative'
    }

    if (formData.new_remain2 < 0) {
      newErrors.new_remain2 = 'Secondary unit quantity cannot be negative'
    }

    if (!formData.reason.trim()) {
      newErrors.reason = 'Reason is required for stock adjustment'
    }

    setErrors(newErrors)
    return Object.keys(newErrors).length === 0
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    
    if (!validateForm()) {
      return
    }

    try {
      await onSubmit(formData)
    } catch (error) {
      console.error('Error submitting stock adjustment form:', error)
    }
  }

  const handleChange = (field: keyof StockAdjustRequest, value: string | number) => {
    setFormData(prev => ({
      ...prev,
      [field]: value
    }))

    // Clear error when user starts typing
    if (errors[field]) {
      setErrors(prev => ({
        ...prev,
        [field]: undefined
      }))
    }
  }

  const getDifference = () => {
    if (!currentStock) return null
    
    const diff1 = formData.new_remain1 - currentStock.remain1
    const diff2 = formData.new_remain2 - currentStock.remain2
    
    return { diff1, diff2 }
  }

  const difference = getDifference()

  return (
    <div className="card">
      <div className="card-header">
        <h3 className="text-lg font-medium text-gray-900">Stock Adjustment</h3>
        <p className="text-sm text-gray-600">Adjust stock levels for inventory corrections</p>
      </div>
      <div className="card-content">
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Product *
              </label>
              <select
                value={formData.product_id}
                onChange={(e) => handleChange('product_id', parseInt(e.target.value))}
                className={`input ${errors.product_id ? 'border-red-500' : ''}`}
                disabled={isLoading}
              >
                <option value={0}>Select a product</option>
                {products.map(product => (
                  <option key={product.id} value={product.id}>
                    {product.name} ({product.short_name})
                  </option>
                ))}
              </select>
              {errors.product_id && (
                <p className="text-red-500 text-sm mt-1">{errors.product_id}</p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Warehouse *
              </label>
              <select
                value={formData.warehouse_id}
                onChange={(e) => handleChange('warehouse_id', parseInt(e.target.value))}
                className={`input ${errors.warehouse_id ? 'border-red-500' : ''}`}
                disabled={isLoading}
              >
                <option value={0}>Select a warehouse</option>
                {warehouses.map(warehouse => (
                  <option key={warehouse.id} value={warehouse.id}>
                    {warehouse.name}
                  </option>
                ))}
              </select>
              {errors.warehouse_id && (
                <p className="text-red-500 text-sm mt-1">{errors.warehouse_id}</p>
              )}
            </div>
          </div>

          {loadingStock && (
            <div className="text-center py-4">
              <p className="text-gray-500">Loading current stock...</p>
            </div>
          )}

          {currentStock && selectedProduct && (
            <div className="bg-gray-50 border border-gray-200 rounded-lg p-4">
              <h4 className="text-sm font-medium text-gray-900 mb-3">Current Stock</h4>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-gray-600">
                    {selectedProduct.unit1}: <span className="font-medium">{currentStock.remain1}</span>
                  </p>
                </div>
                {selectedProduct.unit2 && (
                  <div>
                    <p className="text-sm text-gray-600">
                      {selectedProduct.unit2}: <span className="font-medium">{currentStock.remain2}</span>
                    </p>
                  </div>
                )}
              </div>
            </div>
          )}

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                New {selectedProduct?.unit1 || 'Primary Unit'} Quantity *
              </label>
              <input
                type="number"
                value={formData.new_remain1}
                onChange={(e) => handleChange('new_remain1', parseInt(e.target.value) || 0)}
                className={`input ${errors.new_remain1 ? 'border-red-500' : ''}`}
                placeholder="0"
                min="0"
                disabled={isLoading || loadingStock}
              />
              {difference && (
                <p className={`text-sm mt-1 ${difference.diff1 >= 0 ? 'text-green-600' : 'text-red-600'}`}>
                  {difference.diff1 >= 0 ? '+' : ''}{difference.diff1}
                </p>
              )}
              {errors.new_remain1 && (
                <p className="text-red-500 text-sm mt-1">{errors.new_remain1}</p>
              )}
            </div>

            {selectedProduct?.unit2 && (
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  New {selectedProduct.unit2} Quantity *
                </label>
                <input
                  type="number"
                  value={formData.new_remain2}
                  onChange={(e) => handleChange('new_remain2', parseInt(e.target.value) || 0)}
                  className={`input ${errors.new_remain2 ? 'border-red-500' : ''}`}
                  placeholder="0"
                  min="0"
                  disabled={isLoading || loadingStock}
                />
                {difference && (
                  <p className={`text-sm mt-1 ${difference.diff2 >= 0 ? 'text-green-600' : 'text-red-600'}`}>
                    {difference.diff2 >= 0 ? '+' : ''}{difference.diff2}
                  </p>
                )}
                {errors.new_remain2 && (
                  <p className="text-red-500 text-sm mt-1">{errors.new_remain2}</p>
                )}
              </div>
            )}
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Reason *
            </label>
            <select
              value={formData.reason}
              onChange={(e) => handleChange('reason', e.target.value)}
              className={`input ${errors.reason ? 'border-red-500' : ''}`}
              disabled={isLoading}
            >
              <option value="">Select reason</option>
              <option value="Physical Count">Physical Count</option>
              <option value="Damaged Goods">Damaged Goods</option>
              <option value="Expired Items">Expired Items</option>
              <option value="System Error">System Error</option>
              <option value="Theft/Loss">Theft/Loss</option>
              <option value="Other">Other</option>
            </select>
            {errors.reason && (
              <p className="text-red-500 text-sm mt-1">{errors.reason}</p>
            )}
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Note
            </label>
            <textarea
              value={formData.note}
              onChange={(e) => handleChange('note', e.target.value)}
              className="input"
              rows={3}
              placeholder="Enter additional notes"
              disabled={isLoading}
            />
          </div>

          <div className="flex justify-end space-x-3 pt-4 border-t">
            <button
              type="button"
              onClick={onCancel}
              className="btn-secondary"
              disabled={isLoading}
            >
              Cancel
            </button>
            <button
              type="submit"
              className="btn-primary"
              disabled={isLoading || loadingStock}
            >
              {isLoading ? 'Processing...' : 'Adjust Stock'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}