import React, { useState, useEffect } from 'react'
import { PackUpRequest, Product, Warehouse } from '../types'

interface PackUpFormProps {
  products: Product[]
  warehouses: Warehouse[]
  onSubmit: (data: PackUpRequest) => Promise<void>
  onCancel: () => void
  isLoading?: boolean
}

export default function PackUpForm({ products, warehouses, onSubmit, onCancel, isLoading = false }: PackUpFormProps) {
  const [formData, setFormData] = useState<PackUpRequest>({
    product_id: 0,
    warehouse_id: 0,
    qty_unit2: 0,
    note: ''
  })

  const [errors, setErrors] = useState<Partial<Record<keyof PackUpRequest, string>>>({})
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null)

  useEffect(() => {
    if (formData.product_id > 0) {
      const product = products.find(p => p.id === formData.product_id)
      setSelectedProduct(product || null)
    }
  }, [formData.product_id, products])

  const validateForm = (): boolean => {
    const newErrors: Partial<Record<keyof PackUpRequest, string>> = {}

    if (formData.product_id <= 0) {
      newErrors.product_id = 'Please select a product'
    }

    if (formData.warehouse_id <= 0) {
      newErrors.warehouse_id = 'Please select a warehouse'
    }

    if (formData.qty_unit2 <= 0) {
      newErrors.qty_unit2 = 'Quantity must be greater than 0'
    }

    if (selectedProduct && !selectedProduct.unit2) {
      newErrors.product_id = 'Selected product must have a secondary unit for packing'
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
      console.error('Error submitting pack up form:', error)
    }
  }

  const handleChange = (field: keyof PackUpRequest, value: string | number) => {
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

  const getConversionInfo = () => {
    if (!selectedProduct || !selectedProduct.unit2) return null
    
    const unit2Qty = formData.qty_unit2
    const unit1Qty = unit2Qty * selectedProduct.ratio
    
    return {
      unit2Qty,
      unit1Qty,
      unit1: selectedProduct.unit1,
      unit2: selectedProduct.unit2,
      ratio: selectedProduct.ratio
    }
  }

  const conversionInfo = getConversionInfo()

  return (
    <div className="card">
      <div className="card-header">
        <h3 className="text-lg font-medium text-gray-900">Pack Up</h3>
        <p className="text-sm text-gray-600">Convert from smaller units to larger units</p>
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
                {products.filter(p => p.unit2).map(product => (
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
                Quantity to Pack Up *
              </label>
              <input
                type="number"
                value={formData.qty_unit2}
                onChange={(e) => handleChange('qty_unit2', parseInt(e.target.value) || 0)}
                className={`input ${errors.qty_unit2 ? 'border-red-500' : ''}`}
                placeholder="0"
                min="1"
                disabled={isLoading}
              />
              {selectedProduct?.unit2 && (
                <p className="text-sm text-gray-500 mt-1">
                  Target unit: {selectedProduct.unit2}
                </p>
              )}
              {errors.qty_unit2 && (
                <p className="text-red-500 text-sm mt-1">{errors.qty_unit2}</p>
              )}
            </div>
          </div>

          {conversionInfo && (
            <div className="bg-green-50 border border-green-200 rounded-lg p-4">
              <h4 className="text-sm font-medium text-green-900 mb-2">Conversion Preview</h4>
              <div className="text-sm text-green-800">
                <p>
                  Packing up to <strong>{conversionInfo.unit2Qty} {conversionInfo.unit2}</strong> will consume:
                </p>
                <p className="mt-1">
                  <strong>{conversionInfo.unit1Qty} {conversionInfo.unit1}</strong>
                </p>
                <p className="text-xs text-green-600 mt-2">
                  Ratio: 1 {conversionInfo.unit2} = {conversionInfo.ratio} {conversionInfo.unit1}
                </p>
              </div>
            </div>
          )}

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
              {isLoading ? 'Processing...' : 'Pack Up'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}