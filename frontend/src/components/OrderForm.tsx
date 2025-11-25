import React, { useState, useEffect } from 'react'
import { OrderWithDetails, OrderFormData, Customer, Product, OrderItemRequest } from '../types'
import { customerService } from '../services/customerService'
import { productService } from '../services/productService'

interface OrderFormProps {
  order?: OrderWithDetails
  onSubmit: (data: OrderFormData) => Promise<void>
  onCancel: () => void
  isLoading?: boolean
}

export default function OrderForm({ order, onSubmit, onCancel, isLoading = false }: OrderFormProps) {
  const [formData, setFormData] = useState<OrderFormData>({
    customer_id: 0,
    order_date: new Date().toISOString().split('T')[0],
    delivery_date: undefined,
    payment_terms: '',
    delivery_address: '',
    note: '',
    items: []
  })

  const [customers, setCustomers] = useState<Customer[]>([])
  const [products, setProducts] = useState<Product[]>([])
  const [errors, setErrors] = useState<Partial<Record<keyof OrderFormData, string>>>({})

  useEffect(() => {
    loadCustomers()
    loadProducts()
  }, [])

  useEffect(() => {
    if (order) {
      setFormData({
        customer_id: order.order.customer_id,
        order_date: order.order.order_date.split('T')[0],
        delivery_date: order.order.delivery_date ? order.order.delivery_date.split('T')[0] : undefined,
        payment_terms: order.order.payment_terms,
        delivery_address: order.order.delivery_address,
        note: order.order.note,
        items: order.items.map(item => ({
          product_id: item.order_item.product_id,
          quantity: item.order_item.quantity,
          unit: item.order_item.unit,
          unit_price: item.order_item.unit_price,
          note: item.order_item.note?.String || ""
        }))
      })
    }
  }, [order])

  const loadCustomers = async () => {
    try {
      const customers = await customerService.getCustomers()
      setCustomers(customers)
    } catch (error) {
      console.error('Error loading customers:', error)
    }
  }

  const loadProducts = async () => {
    try {
      const products = await productService.getProducts()
      setProducts(products)
    } catch (error) {
      console.error('Error loading products:', error)
    }
  }

  const validateForm = (): boolean => {
    const newErrors: Partial<Record<keyof OrderFormData, string>> = {}

    if (!formData.customer_id) {
      newErrors.customer_id = 'Customer is required'
    }

    if (!formData.order_date) {
      newErrors.order_date = 'Order date is required'
    }

    if (formData.items.length === 0) {
      newErrors.items = 'At least one item is required'
    }

    // Validate items
    for (let i = 0; i < formData.items.length; i++) {
      const item = formData.items[i]
      if (!item.product_id || item.quantity <= 0 || item.unit_price <= 0) {
        newErrors.items = 'All items must have valid product, quantity, and price'
        break
      }
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
      console.error('Error submitting order form:', error)
    }
  }

  const handleChange = (field: keyof OrderFormData, value: any) => {
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

  const addItem = () => {
    setFormData(prev => ({
      ...prev,
      items: [...prev.items, {
        product_id: 0,
        quantity: 1,
        unit: '',
        unit_price: 0,
        note: ''
      }]
    }))
  }

  const removeItem = (index: number) => {
    setFormData(prev => ({
      ...prev,
      items: prev.items.filter((_, i) => i !== index)
    }))
  }

  const updateItem = (index: number, field: keyof OrderItemRequest, value: any) => {
    setFormData(prev => ({
      ...prev,
      items: prev.items.map((item, i) => 
        i === index ? { ...item, [field]: value } : item
      )
    }))
  }

  const getProductById = (id: number) => {
    return products.find(p => p.id === id)
  }

  const calculateTotal = () => {
    return formData.items.reduce((total, item) => total + (item.quantity * item.unit_price), 0)
  }

  return (
    <div className="card">
      <div className="card-header">
        <h3 className="text-lg font-medium text-gray-900">
          {order ? 'Edit Order' : 'Create New Order'}
        </h3>
      </div>
      <div className="card-content">
        <form onSubmit={handleSubmit} className="space-y-6">
          {/* Order Information */}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Customer *
              </label>
              <select
                value={formData.customer_id}
                onChange={(e) => handleChange('customer_id', parseInt(e.target.value))}
                className={`input ${errors.customer_id ? 'border-red-500' : ''}`}
                disabled={isLoading}
              >
                <option value={0}>Select customer</option>
                {customers.map(customer => (
                  <option key={customer.id} value={customer.id}>
                    {customer.prefix} {customer.name}
                  </option>
                ))}
              </select>
              {errors.customer_id && (
                <p className="text-red-500 text-sm mt-1">{errors.customer_id}</p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Order Date *
              </label>
              <input
                type="date"
                value={formData.order_date}
                onChange={(e) => handleChange('order_date', e.target.value)}
                className={`input ${errors.order_date ? 'border-red-500' : ''}`}
                disabled={isLoading}
              />
              {errors.order_date && (
                <p className="text-red-500 text-sm mt-1">{errors.order_date}</p>
              )}
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Delivery Date
              </label>
              <input
                type="date"
                value={formData.delivery_date || ''}
                onChange={(e) => handleChange('delivery_date', e.target.value || undefined)}
                className="input"
                disabled={isLoading}
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Payment Terms
              </label>
              <input
                type="text"
                value={formData.payment_terms}
                onChange={(e) => handleChange('payment_terms', e.target.value)}
                className="input"
                placeholder="e.g., Net 30, COD"
                disabled={isLoading}
              />
            </div>

            <div className="md:col-span-2">
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Delivery Address
              </label>
              <textarea
                value={formData.delivery_address}
                onChange={(e) => handleChange('delivery_address', e.target.value)}
                className="input"
                rows={2}
                placeholder="Enter delivery address"
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
                rows={2}
                placeholder="Enter order notes"
                disabled={isLoading}
              />
            </div>
          </div>

          {/* Order Items */}
          <div>
            <div className="flex justify-between items-center mb-4">
              <h4 className="text-md font-medium text-gray-900">Order Items</h4>
              <button
                type="button"
                onClick={addItem}
                className="btn-secondary"
                disabled={isLoading}
              >
                Add Item
              </button>
            </div>

            {errors.items && (
              <p className="text-red-500 text-sm mb-4">{errors.items}</p>
            )}

            <div className="space-y-4">
              {formData.items.map((item, index) => {
                const product = getProductById(item.product_id)
                return (
                  <div key={index} className="border rounded-lg p-4 bg-gray-50">
                    <div className="grid grid-cols-1 md:grid-cols-6 gap-4">
                      <div className="md:col-span-2">
                        <label className="block text-sm font-medium text-gray-700 mb-1">
                          Product
                        </label>
                        <select
                          value={item.product_id}
                          onChange={(e) => {
                            const productId = parseInt(e.target.value)
                            const selectedProduct = getProductById(productId)
                            updateItem(index, 'product_id', productId)
                            if (selectedProduct) {
                              updateItem(index, 'unit', selectedProduct.unit1)
                              updateItem(index, 'unit_price', selectedProduct.cost)
                            }
                          }}
                          className="input"
                          disabled={isLoading}
                        >
                          <option value={0}>Select product</option>
                          {products.map(product => (
                            <option key={product.id} value={product.id}>
                              {product.name}
                            </option>
                          ))}
                        </select>
                      </div>

                      <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">
                          Quantity
                        </label>
                        <input
                          type="number"
                          value={item.quantity}
                          onChange={(e) => updateItem(index, 'quantity', parseFloat(e.target.value) || 0)}
                          className="input"
                          min="0"
                          step="0.01"
                          disabled={isLoading}
                        />
                      </div>

                      <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">
                          Unit
                        </label>
                        <input
                          type="text"
                          value={item.unit}
                          onChange={(e) => updateItem(index, 'unit', e.target.value)}
                          className="input"
                          placeholder={product ? product.unit1 : 'Unit'}
                          disabled={isLoading}
                        />
                      </div>

                      <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">
                          Unit Price
                        </label>
                        <input
                          type="number"
                          value={item.unit_price}
                          onChange={(e) => updateItem(index, 'unit_price', parseFloat(e.target.value) || 0)}
                          className="input"
                          min="0"
                          step="0.01"
                          disabled={isLoading}
                        />
                      </div>

                      <div className="flex items-end">
                        <button
                          type="button"
                          onClick={() => removeItem(index)}
                          className="btn-danger w-full"
                          disabled={isLoading}
                        >
                          Remove
                        </button>
                      </div>
                    </div>

                    <div className="mt-4">
                      <label className="block text-sm font-medium text-gray-700 mb-1">
                        Item Note
                      </label>
                      <input
                        type="text"
                        value={item.note}
                        onChange={(e) => updateItem(index, 'note', e.target.value)}
                        className="input"
                        placeholder="Item-specific notes"
                        disabled={isLoading}
                      />
                    </div>

                    <div className="mt-2 text-right">
                      <span className="text-sm font-medium text-gray-900">
                        Total: ${(item.quantity * item.unit_price).toFixed(2)}
                      </span>
                    </div>
                  </div>
                )
              })}
            </div>

            {formData.items.length > 0 && (
              <div className="mt-4 p-4 bg-blue-50 rounded-lg">
                <div className="text-right">
                  <span className="text-lg font-bold text-blue-900">
                    Order Total: ${calculateTotal().toFixed(2)}
                  </span>
                </div>
              </div>
            )}
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
              {isLoading ? 'Saving...' : (order ? 'Update Order' : 'Create Order')}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}