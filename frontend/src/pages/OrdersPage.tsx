import React, { useState, useEffect } from 'react'
import { OrderWithDetails, OrderFormData } from '../types'
import OrderForm from '../components/OrderForm'
import { orderService } from '../services/orderService'

export default function OrdersPage() {
  const [orders, setOrders] = useState<OrderWithDetails[]>([])
  const [showForm, setShowForm] = useState(false)
  const [editingOrder, setEditingOrder] = useState<OrderWithDetails | undefined>()
  const [isLoading, setIsLoading] = useState(false)
  const [statusFilter, setStatusFilter] = useState('')

  useEffect(() => {
    loadOrders()
  }, [])

  const loadOrders = async () => {
    setIsLoading(true)
    try {
      const orders = await orderService.getOrders()
      setOrders(orders)
    } catch (error) {
      console.error('Error loading orders:', error)
    } finally {
      setIsLoading(false)
    }
  }

  const handleCreateOrder = async (data: OrderFormData) => {
    setIsLoading(true)
    try {
      const payload = {
        ...data,
        order_date: new Date(data.order_date).toISOString(),
        delivery_date: data.delivery_date
          ? new Date(data.delivery_date).toISOString()
          : null
      }

      await orderService.createOrder(payload)
      await loadOrders() // Reload to get the full order details
      setShowForm(false)
    } catch (error) {
      console.error('Error creating order:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const handleUpdateOrder = async (data: OrderFormData) => {
    if (!editingOrder) return
    
    setIsLoading(true)
    try {
      const payload = {
        ...data,
        order_date: new Date(data.order_date).toISOString(),
        delivery_date: data.delivery_date
          ? new Date(data.delivery_date).toISOString()
          : null
      }

      await orderService.updateOrder(editingOrder.order.id, payload)

      await loadOrders() // Reload to get the updated order details
      setShowForm(false)
      setEditingOrder(undefined)
    } catch (error) {
      console.error('Error updating order:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const handleDeleteOrder = async (order: OrderWithDetails) => {
    if (!confirm(`Are you sure you want to delete order "${order.order.order_number}"?`)) {
      return
    }

    setIsLoading(true)
    try {
      await orderService.deleteOrder(order.order.id)
      setOrders(prev => prev.filter(o => o.order.id !== order.order.id))
    } catch (error) {
      console.error('Error deleting order:', error)
    } finally {
      setIsLoading(false)
    }
  }

  const handleEditOrder = (order: OrderWithDetails) => {
    setEditingOrder(order)
    setShowForm(true)
  }

  const handleCancelForm = () => {
    setShowForm(false)
    setEditingOrder(undefined)
  }

  const handleStatusChange = async (orderId: number, newStatus: string) => {
    setIsLoading(true)
    try {
      await orderService.updateOrderStatus(orderId, newStatus)
      setOrders(prev => prev.map(o => 
        o.order.id === orderId 
          ? { ...o, order: { ...o.order, status: newStatus } }
          : o
      ))
    } catch (error) {
      console.error('Error updating order status:', error)
    } finally {
      setIsLoading(false)
    }
  }

  const getStatusColor = (status: string) => {
    switch (status.toLowerCase()) {
      case 'pending':
        return 'bg-yellow-100 text-yellow-800'
      case 'confirmed':
        return 'bg-blue-100 text-blue-800'
      case 'shipped':
        return 'bg-purple-100 text-purple-800'
      case 'delivered':
        return 'bg-green-100 text-green-800'
      case 'cancelled':
        return 'bg-red-100 text-red-800'
      default:
        return 'bg-gray-100 text-gray-800'
    }
  }

  const filteredOrders = statusFilter 
    ? orders.filter(order => order.order.status === statusFilter)
    : orders

  if (showForm) {
    return (
      <div className="space-y-6">
        <OrderForm
          order={editingOrder}
          onSubmit={editingOrder ? handleUpdateOrder : handleCreateOrder}
          onCancel={handleCancelForm}
          isLoading={isLoading}
        />
      </div>
    )
  }

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Orders</h1>
          <p className="text-gray-600">Manage customer orders</p>
        </div>
        <button
          onClick={() => setShowForm(true)}
          className="btn-primary"
          disabled={isLoading}
        >
          Create Order
        </button>
      </div>

      {/* Filters */}
      <div className="card">
        <div className="card-content">
          <div className="flex gap-4 items-center">
            <label className="text-sm font-medium text-gray-700">Filter by Status:</label>
            <select
              value={statusFilter}
              onChange={(e) => setStatusFilter(e.target.value)}
              className="input w-48"
              disabled={isLoading}
            >
              <option value="">All Statuses</option>
              <option value="pending">Pending</option>
              <option value="confirmed">Confirmed</option>
              <option value="shipped">Shipped</option>
              <option value="delivered">Delivered</option>
              <option value="cancelled">Cancelled</option>
            </select>
          </div>
        </div>
      </div>

      {/* Orders List */}
      <div className="card">
        <div className="card-header">
          <h3 className="text-lg font-medium text-gray-900">
            Orders ({filteredOrders.length})
          </h3>
        </div>
        <div className="card-content">
          {isLoading ? (
            <div className="text-center py-8">
              <p className="text-gray-500">Loading orders...</p>
            </div>
          ) : filteredOrders.length === 0 ? (
            <div className="text-center py-8">
              <p className="text-gray-500">
                {statusFilter ? `No orders found with status "${statusFilter}"` : 'No orders found'}
              </p>
              {!statusFilter && (
                <button
                  onClick={() => setShowForm(true)}
                  className="btn-primary mt-4"
                >
                  Create Your First Order
                </button>
              )}
            </div>
          ) : (
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Order
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Customer
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Dates
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Status
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Amount
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Items
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Actions
                    </th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {filteredOrders.map((orderDetail) => (
                    <tr key={orderDetail.order.id} className="hover:bg-gray-50">
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div>
                          <div className="text-sm font-medium text-gray-900">
                            {orderDetail.order.order_number}
                          </div>
                          <div className="text-sm text-gray-500">
                            {orderDetail.order.payment_terms}
                          </div>
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div>
                          <div className="text-sm font-medium text-gray-900">
                            {orderDetail.customer.prefix} {orderDetail.customer.name}
                          </div>
                          <div className="text-sm text-gray-500">
                            {orderDetail.customer.phone}
                          </div>
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div>
                          <div className="text-sm text-gray-900">
                            Order: {new Date(orderDetail.order.order_date).toLocaleDateString()}
                          </div>
                          {orderDetail.order.delivery_date && (
                            <div className="text-sm text-gray-500">
                              Delivery: {new Date(orderDetail.order.delivery_date).toLocaleDateString()}
                            </div>
                          )}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <select
                          value={orderDetail.order.status}
                          onChange={(e) => handleStatusChange(orderDetail.order.id, e.target.value)}
                          className={`text-xs font-semibold rounded-full px-2 py-1 border-0 ${getStatusColor(orderDetail.order.status)}`}
                          disabled={isLoading}
                        >
                          <option value="pending">Pending</option>
                          <option value="confirmed">Confirmed</option>
                          <option value="shipped">Shipped</option>
                          <option value="delivered">Delivered</option>
                          <option value="cancelled">Cancelled</option>
                        </select>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm font-medium text-gray-900">
                          ${orderDetail.order.final_amount.toFixed(2)}
                        </div>
                        {orderDetail.order.discount > 0 && (
                          <div className="text-sm text-gray-500">
                            Discount: ${orderDetail.order.discount.toFixed(2)}
                          </div>
                        )}
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-900">
                          {orderDetail.items.length} item(s)
                        </div>
                        <div className="text-sm text-gray-500">
                          {orderDetail.items.slice(0, 2).map(item => item.product_name).join(', ')}
                          {orderDetail.items.length > 2 && '...'}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                        <button
                          onClick={() => handleEditOrder(orderDetail)}
                          className="text-indigo-600 hover:text-indigo-900"
                          disabled={orderDetail.order.status === 'delivered' || orderDetail.order.status === 'cancelled'}
                        >
                          Edit
                        </button>
                        <button
                          onClick={() => handleDeleteOrder(orderDetail)}
                          className="text-red-600 hover:text-red-900"
                          disabled={orderDetail.order.status === 'delivered'}
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}