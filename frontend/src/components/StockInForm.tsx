import React, { useState, useEffect } from 'react'
import { StockInRequest, Product, Warehouse } from '../types'

interface StockInFormProps {
  products: Product[]
  warehouses: Warehouse[]
  onSubmit: (data: StockInRequest) => Promise<void>
  onCancel: () => void
  isLoading?: boolean
}

export default function StockInForm({ products, warehouses, onSubmit, onCancel, isLoading = false }: StockInFormProps) {
  const [formData, setFormData] = useState<StockInRequest>({
    product_id: 0,
    warehouse_id: 0,
    qty: 0,
    unit: '',
    ref_type: '',
    ref_id: undefined,
    note: ''
  })

  const [errors, setErrors] = useState<Partial<Record<keyof StockInRequest, string>>>({})
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null)

  useEffect(() => {
    if (formData.product_id > 0) {
      const product = products.find(p => p.id === formData.product_id)
      setSelectedProduct(product || null)
      if (product && !formData.unit) {
        setFormData(prev => ({ ...prev, unit: product.unit1 }))
      }
    }
  }, [formData.product_id, products])

  const validateForm = (): boolean => {
    const newErrors: Partial<Record<keyof StockInRequest, string>> = {}

    if (formData.product_id <= 0) {
      newErrors.product_id = 'Please select a product'
    }

    if (formData.warehouse_id <= 0) {
      newErrors.warehouse_id = 'Please select a warehouse'
    }

    if (formData.qty <= 0) {
      newErrors.qty = 'Quantity must be greater than 0'
    }

    if (!formData.unit.trim()) {
      newErrors.unit = 'Please select a unit'
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
      console.error('Error submitting stock in form:', error)
    }
  }

  const handleChange = (field: keyof StockInRequest, value: string | number | undefined) => {
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

  return (
    <div className="card">
      <div className="card-header">
        <h3 className="text-lg font-medium text-gray-900">Stock In</h3>
        <p className="text-sm text-gray-600">Add stock to warehouse</p>
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

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Quantity *
              </label>
              <input
                type="number"
                value={formData.qty}
                onChange={(e) => handleChange('qty', parseFloat(e.target.value) || 0)}
                className={`input ${errors.qty ? 'border-red-500' : ''}`}
                placeholder="0"
                min="0"
                step="0.01"
                disabled={isLoading}
              />
              {errors.qty && (
                <p className="text-red-500 text-sm mt-1">{errors.qty}</p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Unit *
              </label>
              <select
                value={formData.unit}
                onChange={(e) => handleChange('unit', e.target.value)}
                className={`input ${errors.unit ? 'border-red-500' : ''}`}
                disabled={isLoading || !selectedProduct}
              >
                <option value="">Select unit</option>
                {selectedProduct && (
                  <>
                    <option value={selectedProduct.unit1}>{selectedProduct.unit1}</option>
                    {selectedProduct.unit2 && (
                      <option value={selectedProduct.unit2}>{selectedProduct.unit2}</option>
                    )}
                  </>
                )}
              </select>
              {errors.unit && (
                <p className="text-red-500 text-sm mt-1">{errors.unit}</p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Reference Type
              </label>
              <select
                value={formData.ref_type}
                onChange={(e) => handleChange('ref_type', e.target.value)}
                className="input"
                disabled={isLoading}
              >
                <option value="">Select reference type</option>
                <option value="PO">Purchase Order</option>
                <option value="GRN">Goods Receipt</option>
                <option value="ADJ">Adjustment</option>
                <option value="OTHER">Other</option>
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Reference ID
              </label>
              <input
                type="number"
                value={formData.ref_id || ''}
                onChange={(e) => handleChange('ref_id', e.target.value ? parseInt(e.target.value) : undefined)}
                className="input"
                placeholder="Reference ID"
                disabled={isLoading}
              />
            </div>
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
              disabled={isLoading}
            >
              {isLoading ? 'Processing...' : 'Stock In'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}