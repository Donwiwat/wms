import React, { useState, useEffect } from 'react'
import { Customer, CustomerFormData } from '../types'

interface CustomerFormProps {
  customer?: Customer
  onSubmit: (data: CustomerFormData) => Promise<void>
  onCancel: () => void
  isLoading?: boolean
}

export default function CustomerForm({ customer, onSubmit, onCancel, isLoading = false }: CustomerFormProps) {
  const [formData, setFormData] = useState<CustomerFormData>({
    prefix: '',
    name: '',
    address: '',
    phone: '',
    contact_person: '',
    level: '',
    delivery_place: '',
    transport: '',
    credit_limit: 0,
    credit_term: 0,
    outstanding: 0,
    last_contact: undefined,
    note: ''
  })

  const [errors, setErrors] = useState<Partial<Record<keyof CustomerFormData, string>>>({})

  useEffect(() => {
    if (customer) {
      setFormData({
        prefix: customer.prefix,
        name: customer.name,
        address: customer.address,
        phone: customer.phone,
        contact_person: customer.contact_person,
        level: customer.level,
        delivery_place: customer.delivery_place,
        transport: customer.transport,
        credit_limit: customer.credit_limit,
        credit_term: customer.credit_term,
        outstanding: customer.outstanding,
        last_contact: customer.last_contact,
        note: customer.note
      })
    }
  }, [customer])

  const validateForm = (): boolean => {
    const newErrors: Partial<Record<keyof CustomerFormData, string>> = {}

    if (!formData.name.trim()) {
      newErrors.name = 'Customer name is required'
    }

    if (formData.credit_limit < 0) {
      newErrors.credit_limit = 'Credit limit cannot be negative'
    }

    if (formData.credit_term < 0) {
      newErrors.credit_term = 'Credit term cannot be negative'
    }

    if (formData.outstanding < 0) {
      newErrors.outstanding = 'Outstanding cannot be negative'
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
      console.error('Error submitting customer form:', error)
    }
  }

  const handleChange = (field: keyof CustomerFormData, value: string | number | undefined) => {
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
          {customer ? 'Edit Customer' : 'Add New Customer'}
        </h3>
      </div>
      <div className="card-content">
        <form onSubmit={handleSubmit} className="space-y-6">
          {/* Basic Information */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Prefix
              </label>
              <input
                type="text"
                value={formData.prefix}
                onChange={(e) => handleChange('prefix', e.target.value)}
                className="input"
                placeholder="Mr., Mrs., Dr., etc."
                disabled={isLoading}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Customer Name *
              </label>
              <input
                type="text"
                value={formData.name}
                onChange={(e) => handleChange('name', e.target.value)}
                className={`input ${errors.name ? 'border-red-500' : ''}`}
                placeholder="Enter customer name"
                disabled={isLoading}
              />
              {errors.name && (
                <p className="text-red-500 text-sm mt-1">{errors.name}</p>
              )}
            </div>

            <div className="md:col-span-2">
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Address
              </label>
              <textarea
                value={formData.address}
                onChange={(e) => handleChange('address', e.target.value)}
                className="input"
                rows={3}
                placeholder="Enter customer address"
                disabled={isLoading}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Phone
              </label>
              <input
                type="text"
                value={formData.phone}
                onChange={(e) => handleChange('phone', e.target.value)}
                className="input"
                placeholder="Enter phone number"
                disabled={isLoading}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Contact Person
              </label>
              <input
                type="text"
                value={formData.contact_person}
                onChange={(e) => handleChange('contact_person', e.target.value)}
                className="input"
                placeholder="Enter contact person name"
                disabled={isLoading}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Level
              </label>
              <select
                value={formData.level}
                onChange={(e) => handleChange('level', e.target.value)}
                className="input"
                disabled={isLoading}
              >
                <option value="">Select level</option>
                <option value="VIP">VIP</option>
                <option value="Premium">Premium</option>
                <option value="Standard">Standard</option>
                <option value="Basic">Basic</option>
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Transport
              </label>
              <input
                type="text"
                value={formData.transport}
                onChange={(e) => handleChange('transport', e.target.value)}
                className="input"
                placeholder="Enter transport method"
                disabled={isLoading}
              />
            </div>

            <div className="md:col-span-2">
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Delivery Place
              </label>
              <textarea
                value={formData.delivery_place}
                onChange={(e) => handleChange('delivery_place', e.target.value)}
                className="input"
                rows={2}
                placeholder="Enter delivery place"
                disabled={isLoading}
              />
            </div>
          </div>

          {/* Financial Information */}
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Credit Limit
              </label>
              <input
                type="number"
                value={formData.credit_limit}
                onChange={(e) => handleChange('credit_limit', parseFloat(e.target.value) || 0)}
                className={`input ${errors.credit_limit ? 'border-red-500' : ''}`}
                placeholder="0.00"
                min="0"
                step="0.01"
                disabled={isLoading}
              />
              {errors.credit_limit && (
                <p className="text-red-500 text-sm mt-1">{errors.credit_limit}</p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Credit Term (days)
              </label>
              <input
                type="number"
                value={formData.credit_term}
                onChange={(e) => handleChange('credit_term', parseInt(e.target.value) || 0)}
                className={`input ${errors.credit_term ? 'border-red-500' : ''}`}
                placeholder="0"
                min="0"
                disabled={isLoading}
              />
              {errors.credit_term && (
                <p className="text-red-500 text-sm mt-1">{errors.credit_term}</p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Outstanding
              </label>
              <input
                type="number"
                value={formData.outstanding}
                onChange={(e) => handleChange('outstanding', parseFloat(e.target.value) || 0)}
                className={`input ${errors.outstanding ? 'border-red-500' : ''}`}
                placeholder="0.00"
                min="0"
                step="0.01"
                disabled={isLoading}
              />
              {errors.outstanding && (
                <p className="text-red-500 text-sm mt-1">{errors.outstanding}</p>
              )}
            </div>
          </div>

          {/* Additional Information */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Last Contact
              </label>
              <input
                type="date"
                value={formData.last_contact ? formData.last_contact.split('T')[0] : ''}
                onChange={(e) => handleChange('last_contact', e.target.value || undefined)}
                className="input"
                disabled={isLoading}
              />
            </div>

            <div className="md:col-span-2">
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
              {isLoading ? 'Saving...' : (customer ? 'Update Customer' : 'Create Customer')}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}