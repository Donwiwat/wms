import React, { useState, useEffect } from 'react'
import { WarehouseFormData, Warehouse } from '@/types'

interface WarehouseFormProps {
  warehouse?: Warehouse
  onSubmit: (data: WarehouseFormData) => Promise<void>
  onCancel: () => void
  isLoading?: boolean
}

export default function WarehouseForm({ warehouse, onSubmit, onCancel, isLoading = false }: WarehouseFormProps) {
  const [formData, setFormData] = useState<WarehouseFormData>({
    name: '',
    location: '',
    note: ''
  })
  const [errors, setErrors] = useState<Partial<WarehouseFormData>>({})

  useEffect(() => {
    if (warehouse) {
      setFormData({
        name: warehouse.name,
        location: warehouse.location,
        note: warehouse.note
      })
    }
  }, [warehouse])

  const validateForm = (): boolean => {
    const newErrors: Partial<WarehouseFormData> = {}

    if (!formData.name.trim()) {
      newErrors.name = 'Warehouse name is required'
    }

    if (!formData.location.trim()) {
      newErrors.location = 'Location is required'
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
      console.error('Error submitting warehouse form:', error)
    }
  }

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target
    setFormData(prev => ({ ...prev, [name]: value }))
    
    // Clear error when user starts typing
    if (errors[name as keyof WarehouseFormData]) {
      setErrors(prev => ({ ...prev, [name]: undefined }))
    }
  }

  return (
    <div className="card">
      <div className="card-header">
        <h3 className="text-lg font-medium text-gray-900">
          {warehouse ? 'Edit Warehouse' : 'Add New Warehouse'}
        </h3>
      </div>
      <div className="card-content">
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label htmlFor="name" className="block text-sm font-medium text-gray-700 mb-1">
              Warehouse Name *
            </label>
            <input
              type="text"
              id="name"
              name="name"
              value={formData.name}
              onChange={handleChange}
              className={`input ${errors.name ? 'border-red-500' : ''}`}
              placeholder="Enter warehouse name"
              disabled={isLoading}
            />
            {errors.name && (
              <p className="mt-1 text-sm text-red-600">{errors.name}</p>
            )}
          </div>

          <div>
            <label htmlFor="location" className="block text-sm font-medium text-gray-700 mb-1">
              Location *
            </label>
            <input
              type="text"
              id="location"
              name="location"
              value={formData.location}
              onChange={handleChange}
              className={`input ${errors.location ? 'border-red-500' : ''}`}
              placeholder="Enter warehouse location"
              disabled={isLoading}
            />
            {errors.location && (
              <p className="mt-1 text-sm text-red-600">{errors.location}</p>
            )}
          </div>

          <div>
            <label htmlFor="note" className="block text-sm font-medium text-gray-700 mb-1">
              Note
            </label>
            <textarea
              id="note"
              name="note"
              value={formData.note}
              onChange={handleChange}
              rows={3}
              className="input"
              placeholder="Enter additional notes (optional)"
              disabled={isLoading}
            />
          </div>

          <div className="flex justify-end space-x-3 pt-4">
            <button
              type="button"
              onClick={onCancel}
              className="btn btn-secondary"
              disabled={isLoading}
            >
              Cancel
            </button>
            <button
              type="submit"
              className="btn btn-primary"
              disabled={isLoading}
            >
              {isLoading ? 'Saving...' : warehouse ? 'Update Warehouse' : 'Add Warehouse'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}