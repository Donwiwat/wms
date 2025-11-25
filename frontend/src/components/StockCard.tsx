import React from 'react'
import { StockCardEntry, Product, Warehouse } from '../types'

interface StockCardProps {
  product: Product
  warehouse: Warehouse
  entries: StockCardEntry[]
  onClose: () => void
  isLoading?: boolean
}

export default function StockCard({ product, warehouse, entries, onClose, isLoading = false }: StockCardProps) {
  const getTypeColor = (type: string) => {
    switch (type) {
      case 'IN':
        return 'text-green-600 bg-green-100'
      case 'OUT':
        return 'text-red-600 bg-red-100'
      case 'BREAK':
        return 'text-blue-600 bg-blue-100'
      case 'PACK':
        return 'text-purple-600 bg-purple-100'
      case 'TF-IN':
        return 'text-indigo-600 bg-indigo-100'
      case 'TF-OUT':
        return 'text-orange-600 bg-orange-100'
      case 'ADJUST':
        return 'text-yellow-600 bg-yellow-100'
      default:
        return 'text-gray-600 bg-gray-100'
    }
  }

  const getTypeLabel = (type: string) => {
    switch (type) {
      case 'IN':
        return 'Stock In'
      case 'OUT':
        return 'Stock Out'
      case 'BREAK':
        return 'Break Down'
      case 'PACK':
        return 'Pack Up'
      case 'TF-IN':
        return 'Transfer In'
      case 'TF-OUT':
        return 'Transfer Out'
      case 'ADJUST':
        return 'Adjustment'
      default:
        return type
    }
  }

  const getRefTypeLabel = (refType: string) => {
    switch (refType) {
      case 'PO':
        return 'Purchase Order'
      case 'SO':
        return 'Sales Order'
      case 'DO':
        return 'Delivery Order'
      case 'GRN':
        return 'Goods Receipt'
      case 'TF':
        return 'Transfer'
      case 'ADJ':
        return 'Adjustment'
      default:
        return refType
    }
  }

  const calculateRunningBalance = () => {
    let balance = 0
    return entries.map(entry => {
      if (['IN', 'TF-IN', 'BREAK', 'PACK'].includes(entry.type)) {
        balance += entry.qty
      } else if (['OUT', 'TF-OUT'].includes(entry.type)) {
        balance -= entry.qty
      }
      // For ADJUST, the qty represents the new balance, not the change
      if (entry.type === 'ADJUST') {
        balance = entry.qty
      }
      return { ...entry, balance }
    })
  }

  const entriesWithBalance = calculateRunningBalance()

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-lg shadow-xl max-w-6xl w-full max-h-[90vh] overflow-hidden">
        <div className="card-header border-b">
          <div className="flex justify-between items-center">
            <div>
              <h3 className="text-lg font-medium text-gray-900">Stock Card</h3>
              <p className="text-sm text-gray-600">
                {product.name} - {warehouse.name}
              </p>
            </div>
            <button
              onClick={onClose}
              className="text-gray-400 hover:text-gray-600"
            >
              <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        <div className="p-6">
          {/* Product Info */}
          <div className="bg-gray-50 rounded-lg p-4 mb-6">
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
              <div>
                <p className="text-sm font-medium text-gray-700">Product</p>
                <p className="text-sm text-gray-900">{product.name}</p>
                <p className="text-xs text-gray-500">{product.short_name}</p>
              </div>
              <div>
                <p className="text-sm font-medium text-gray-700">Units</p>
                <p className="text-sm text-gray-900">
                  Primary: {product.unit1}
                  {product.unit2 && (
                    <span className="ml-2">Secondary: {product.unit2}</span>
                  )}
                </p>
                {product.unit2 && (
                  <p className="text-xs text-gray-500">
                    Ratio: 1 {product.unit2} = {product.ratio} {product.unit1}
                  </p>
                )}
              </div>
              <div>
                <p className="text-sm font-medium text-gray-700">Warehouse</p>
                <p className="text-sm text-gray-900">{warehouse.name}</p>
                <p className="text-xs text-gray-500">{warehouse.location}</p>
              </div>
            </div>
          </div>

          {/* Stock Movements */}
          <div className="overflow-y-auto max-h-96">
            {isLoading ? (
              <div className="text-center py-8">
                <p className="text-gray-500">Loading stock movements...</p>
              </div>
            ) : entries.length === 0 ? (
              <div className="text-center py-8">
                <p className="text-gray-500">No stock movements found</p>
              </div>
            ) : (
              <div className="overflow-x-auto">
                <table className="min-w-full divide-y divide-gray-200">
                  <thead className="bg-gray-50 sticky top-0">
                    <tr>
                      <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Date
                      </th>
                      <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Type
                      </th>
                      <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Quantity
                      </th>
                      <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Balance
                      </th>
                      <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Reference
                      </th>
                      <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Note
                      </th>
                      <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        Created By
                      </th>
                    </tr>
                  </thead>
                  <tbody className="bg-white divide-y divide-gray-200">
                    {entriesWithBalance.map((entry, index) => (
                      <tr key={index} className="hover:bg-gray-50">
                        <td className="px-4 py-4 whitespace-nowrap">
                          <div className="text-sm text-gray-900">
                            {new Date(entry.date).toLocaleDateString()}
                          </div>
                          <div className="text-xs text-gray-500">
                            {new Date(entry.date).toLocaleTimeString()}
                          </div>
                        </td>
                        <td className="px-4 py-4 whitespace-nowrap">
                          <span className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${getTypeColor(entry.type)}`}>
                            {getTypeLabel(entry.type)}
                          </span>
                        </td>
                        <td className="px-4 py-4 whitespace-nowrap">
                          <div className={`text-sm font-medium ${
                            ['IN', 'TF-IN', 'BREAK', 'PACK'].includes(entry.type) 
                              ? 'text-green-600' 
                              : ['OUT', 'TF-OUT'].includes(entry.type)
                              ? 'text-red-600'
                              : 'text-gray-900'
                          }`}>
                            {['IN', 'TF-IN', 'BREAK', 'PACK'].includes(entry.type) && '+'}
                            {['OUT', 'TF-OUT'].includes(entry.type) && '-'}
                            {entry.qty} {entry.unit}
                          </div>
                        </td>
                        <td className="px-4 py-4 whitespace-nowrap">
                          <div className="text-sm font-medium text-gray-900">
                            {entry.balance} {entry.unit}
                          </div>
                        </td>
                        <td className="px-4 py-4 whitespace-nowrap">
                          <div className="text-sm text-gray-900">
                            {entry.ref_type && (
                              <>
                                {getRefTypeLabel(entry.ref_type)}
                                {entry.ref_id && ` #${entry.ref_id}`}
                              </>
                            )}
                          </div>
                        </td>
                        <td className="px-4 py-4">
                          <div className="text-sm text-gray-900 max-w-xs truncate" title={entry.note}>
                            {entry.note}
                          </div>
                        </td>
                        <td className="px-4 py-4 whitespace-nowrap">
                          <div className="text-sm text-gray-900">{entry.created_by}</div>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            )}
          </div>
        </div>

        <div className="border-t px-6 py-4">
          <div className="flex justify-end">
            <button
              onClick={onClose}
              className="btn-secondary"
            >
              Close
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}