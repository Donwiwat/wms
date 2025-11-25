import React, { useState, useEffect } from 'react'
import { Product, ProductFormData } from '../types'

interface ProductFormProps {
  product?: Product
  onSubmit: (data: ProductFormData) => Promise<void>
  onCancel: () => void
  isLoading?: boolean
}

export default function ProductForm({ product, onSubmit, onCancel, isLoading = false }: ProductFormProps) {
  const [formData, setFormData] = useState<ProductFormData>({
    name: '',
    short_name: '',
    brand: '',
    model: '',
    size: '',
    group: '',
    unit1: '',
    unit2: '',
    ratio: 1,
    cost: 0,
    message: '',
    note: ''
  })

  const [errors, setErrors] = useState<Partial<Record<keyof ProductFormData, string>>>({})

  useEffect(() => {
    if (product) {
      setFormData({
        name: product.name,
        short_name: product.short_name,
        brand: product.brand,
        model: product.model,
        size: product.size,
        group: product.group,
        unit1: product.unit1,
        unit2: product.unit2,
        ratio: product.ratio,
        cost: product.cost,
        message: product.message,
        note: product.note
      })
    }
  }, [product])

  const validateForm = (): boolean => {
    const newErrors: Partial<Record<keyof ProductFormData, string>> = {}

    if (!formData.name.trim()) {
      newErrors.name = 'Product name is required'
    }

    if (!formData.unit1.trim()) {
      newErrors.unit1 = 'Primary unit is required'
    }

    if (formData.ratio <= 0) {
      newErrors.ratio = 'Ratio must be greater than 0'
    }

    if (formData.cost < 0) {
      newErrors.cost = 'Cost cannot be negative'
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
      console.error('Error submitting product form:', error)
    }
  }

  const handleChange = (field: keyof ProductFormData, value: string | number) => {
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
        <h3 className="text-lg font-medium text-gray-900">
          {product ? 'Edit Product' : 'Add New Product'}
        </h3>
      </div>
      <div className="card-content">
        <form onSubmit={handleSubmit} className="space-y-6">
          {/* Basic Information */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Product Name *
              </label>
              <input
                type="text"
                value={formData.name}
                onChange={(e) => handleChange('name', e.target.value)}
                className={`input ${errors.name ? 'border-red-500' : ''}`}
                placeholder="Enter product name"
                disabled={isLoading}
              />
              {errors.name && (
                <p className="text-red-500 text-sm mt-1">{errors.name}</p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Short Name
              </label>
              <input
                type="text"
                value={formData.short_name}
                onChange={(e) => handleChange('short_name', e.target.value)}
                className="input"
                placeholder="Enter short name"
                disabled={isLoading}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Brand
              </label>
              <input
                type="text"
                value={formData.brand}
                onChange={(e) => handleChange('brand', e.target.value)}
                className="input"
                placeholder="Enter brand"
                disabled={isLoading}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Model
              </label>
              <input
                type="text"
                value={formData.model}
                onChange={(e) => handleChange('model', e.target.value)}
                className="input"
                placeholder="Enter model"
                disabled={isLoading}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Size
              </label>
              <input
                type="text"
                value={formData.size}
                onChange={(e) => handleChange('size', e.target.value)}
                className="input"
                placeholder="Enter size"
                disabled={isLoading}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Group
              </label>
              <input
                type="text"
                value={formData.group}
                onChange={(e) => handleChange('group', e.target.value)}
                className="input"
                placeholder="Enter product group"
                disabled={isLoading}
              />
            </div>
          </div>

          {/* Units and Ratio */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Primary Unit *
              </label>
              <input
                type="text"
                value={formData.unit1}
                onChange={(e) => handleChange('unit1', e.target.value)}
                className={`input ${errors.unit1 ? 'border-red-500' : ''}`}
                placeholder="e.g., pcs, kg"
                disabled={isLoading}
              />
              {errors.unit1 && (
                <p className="text-red-500 text-sm mt-1">{errors.unit1}</p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Secondary Unit
              </label>
              <input
                type="text"
                value={formData.unit2}
                onChange={(e) => handleChange('unit2', e.target.value)}
                className="input"
                placeholder="e.g., box, carton"
                disabled={isLoading}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Ratio (Unit2 to Unit1) *
              </label>
              <input
                type="number"
                value={formData.ratio}
                onChange={(e) => handleChange('ratio', parseFloat(e.target.value) || 0)}
                className={`input ${errors.ratio ? 'border-red-500' : ''}`}
                placeholder="1"
                min="0"
                step="0.01"
                disabled={isLoading}
              />
              {errors.ratio && (
                <p className="text-red-500 text-sm mt-1">{errors.ratio}</p>
              )}
            </div>
          </div>

          {/* Cost */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Cost
              </label>
              <input
                type="number"
                value={formData.cost}
                onChange={(e) => handleChange('cost', parseFloat(e.target.value) || 0)}
                className={`input ${errors.cost ? 'border-red-500' : ''}`}
                placeholder="0.00"
                min="0"
                step="0.01"
                disabled={isLoading}
              />
              {errors.cost && (
                <p className="text-red-500 text-sm mt-1">{errors.cost}</p>
              )}
            </div>
          </div>

          {/* Message and Note */}
          <div className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Message
              </label>
              <textarea
                value={formData.message}
                onChange={(e) => handleChange('message', e.target.value)}
                className="input"
                rows={2}
                placeholder="Enter product message"
                disabled={isLoading}
              />
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
          </div>

          {/* Form Actions */}
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
              {isLoading ? 'Saving...' : (product ? 'Update Product' : 'Create Product')}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}