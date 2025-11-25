import React, { useState, useEffect } from 'react'
import { Customer, CustomerFormData } from '../types'
import CustomerForm from '../components/CustomerForm'
import { customerService } from '../services/customerService'

export default function CustomersPage() {
  const [customers, setCustomers] = useState<Customer[]>([])
  const [showForm, setShowForm] = useState(false)
  const [editingCustomer, setEditingCustomer] = useState<Customer | undefined>()
  const [isLoading, setIsLoading] = useState(false)
  const [searchTerm, setSearchTerm] = useState('')

  useEffect(() => {
    loadCustomers()
  }, [])

  const loadCustomers = async () => {
    setIsLoading(true)
    try {
      const customers = await customerService.getCustomers()
      setCustomers(customers)
    } catch (error) {
      console.error('Error loading customers:', error)
    } finally {
      setIsLoading(false)
    }
  }

  const handleCreateCustomer = async (data: CustomerFormData) => {
    setIsLoading(true)
    try {
      const customer = await customerService.createCustomer(data)
      setCustomers(prev => [...prev, customer])
      setShowForm(false)
    } catch (error) {
      console.error('Error creating customer:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const handleUpdateCustomer = async (data: CustomerFormData) => {
    if (!editingCustomer) return
    
    setIsLoading(true)
    try {
      const customer = await customerService.updateCustomer(editingCustomer.id, data)
      setCustomers(prev => prev.map(c => c.id === editingCustomer.id ? customer : c))
      setShowForm(false)
      setEditingCustomer(undefined)
    } catch (error) {
      console.error('Error updating customer:', error)
      throw error
    } finally {
      setIsLoading(false)
    }
  }

  const handleDeleteCustomer = async (customer: Customer) => {
    if (!confirm(`Are you sure you want to delete "${customer.name}"?`)) {
      return
    }

    setIsLoading(true)
    try {
      await customerService.deleteCustomer(customer.id)
      setCustomers(prev => prev.filter(c => c.id !== customer.id))
    } catch (error) {
      console.error('Error deleting customer:', error)
    } finally {
      setIsLoading(false)
    }
  }

  const handleEditCustomer = (customer: Customer) => {
    setEditingCustomer(customer)
    setShowForm(true)
  }

  const handleCancelForm = () => {
    setShowForm(false)
    setEditingCustomer(undefined)
  }

  const handleSearch = async () => {
    if (!searchTerm.trim()) {
      loadCustomers()
      return
    }

    setIsLoading(true)
    try {
      const customers = await customerService.searchCustomers(searchTerm)
      setCustomers(customers)
    } catch (error) {
      console.error('Error searching customers:', error)
    } finally {
      setIsLoading(false)
    }
  }

  const filteredCustomers = searchTerm ? customers : customers.filter(customer =>
    customer.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    customer.contact_person.toLowerCase().includes(searchTerm.toLowerCase()) ||
    customer.phone.toLowerCase().includes(searchTerm.toLowerCase())
  )

  if (showForm) {
    return (
      <div className="space-y-6">
        <CustomerForm
          customer={editingCustomer}
          onSubmit={editingCustomer ? handleUpdateCustomer : handleCreateCustomer}
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
          <h1 className="text-2xl font-bold text-gray-900">Customers</h1>
          <p className="text-gray-600">Manage your customer database</p>
        </div>
        <button
          onClick={() => setShowForm(true)}
          className="btn-primary"
          disabled={isLoading}
        >
          Add Customer
        </button>
      </div>

      {/* Search */}
      <div className="card">
        <div className="card-content">
          <div className="flex gap-2">
            <input
              type="text"
              placeholder="Search customers by name, contact person, or phone..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="input flex-1"
              disabled={isLoading}
              onKeyPress={(e) => e.key === 'Enter' && handleSearch()}
            />
            <button
              onClick={handleSearch}
              className="btn-primary"
              disabled={isLoading}
            >
              Search
            </button>
            {searchTerm && (
              <button
                onClick={() => {
                  setSearchTerm('')
                  loadCustomers()
                }}
                className="btn-secondary"
                disabled={isLoading}
              >
                Clear
              </button>
            )}
          </div>
        </div>
      </div>

      {/* Customer List */}
      <div className="card">
        <div className="card-header">
          <h3 className="text-lg font-medium text-gray-900">
            Customer List ({filteredCustomers.length})
          </h3>
        </div>
        <div className="card-content">
          {isLoading ? (
            <div className="text-center py-8">
              <p className="text-gray-500">Loading customers...</p>
            </div>
          ) : filteredCustomers.length === 0 ? (
            <div className="text-center py-8">
              <p className="text-gray-500">
                {searchTerm ? 'No customers found matching your search' : 'No customers found'}
              </p>
              {!searchTerm && (
                <button
                  onClick={() => setShowForm(true)}
                  className="btn-primary mt-4"
                >
                  Add Your First Customer
                </button>
              )}
            </div>
          ) : (
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Customer
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Contact
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Level
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Credit Info
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Outstanding
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Actions
                    </th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {filteredCustomers.map((customer) => (
                    <tr key={customer.id} className="hover:bg-gray-50">
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div>
                          <div className="text-sm font-medium text-gray-900">
                            {customer.prefix} {customer.name}
                          </div>
                          <div className="text-sm text-gray-500">
                            {customer.address}
                          </div>
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-900">
                          {customer.contact_person}
                        </div>
                        <div className="text-sm text-gray-500">
                          {customer.phone}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <span className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${
                          customer.level === 'VIP' ? 'bg-purple-100 text-purple-800' :
                          customer.level === 'Premium' ? 'bg-blue-100 text-blue-800' :
                          customer.level === 'Standard' ? 'bg-green-100 text-green-800' :
                          'bg-gray-100 text-gray-800'
                        }`}>
                          {customer.level || 'Basic'}
                        </span>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-900">
                          Limit: ${customer.credit_limit.toFixed(2)}
                        </div>
                        <div className="text-sm text-gray-500">
                          Term: {customer.credit_term} days
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className={`text-sm font-medium ${
                          customer.outstanding > 0 ? 'text-red-600' : 'text-green-600'
                        }`}>
                          ${customer.outstanding.toFixed(2)}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                        <button
                          onClick={() => handleEditCustomer(customer)}
                          className="text-indigo-600 hover:text-indigo-900"
                        >
                          Edit
                        </button>
                        <button
                          onClick={() => handleDeleteCustomer(customer)}
                          className="text-red-600 hover:text-red-900"
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