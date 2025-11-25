import React, { useState, useEffect } from 'react'
import { StockSummary as StockSummaryType, Product, Warehouse, StockFilter } from '../types'

interface StockSummaryProps {
  stockSummary: StockSummaryType[]
  products: Product[]
  warehouses: Warehouse[]
  onFilterChange: (filter: StockFilter) => void
  onViewStockCard: (productId: number, warehouseId: number) => void
  isLoading?: boolean
}

export default function StockSummary({ 
  stockSummary, 
  products, 
  warehouses, 
  onFilterChange, 
  onViewStockCard,
  isLoading = false 
}: StockSummaryProps) {
  const [filter, setFilter] = useState<StockFilter>({
    product_id: undefined,
    warehouse_id: undefined
  })

  useEffect(() => {
    onFilterChange(filter)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [filter])


  const handleFilterChange = (field: keyof StockFilter, value: number | undefined) => {
    const newFilter = { ...filter, [field]: value }
    setFilter(newFilter)
  }

  const getTotalStock = () => {
    return (stockSummary ?? []).reduce(
        (acc, item) => {
        acc.totalRemain += item.total_remain || 0
        return acc
        },
        { totalRemain: 0 }
    )
    }

  const totals = getTotalStock()

  return (
    <div className="space-y-6">
      {/* Filters */}
      <div className="card">
        <div className="card-header">
          <h3 className="text-lg font-medium text-gray-900">Filters</h3>
        </div>
        <div className="card-content">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Product
              </label>
              <select
                value={filter.product_id || ''}
                onChange={(e) => handleFilterChange('product_id', e.target.value ? parseInt(e.target.value) : undefined)}
                className="input"
                disabled={isLoading}
              >
                <option value="">All Products</option>
                {products.map(product => (
                  <option key={product.id} value={product.id}>
                    {product.name} ({product.short_name})
                  </option>
                ))}
              </select>
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Warehouse
              </label>
              <select
                value={filter.warehouse_id || ''}
                onChange={(e) => handleFilterChange('warehouse_id', e.target.value ? parseInt(e.target.value) : undefined)}
                className="input"
                disabled={isLoading}
              >
                <option value="">All Warehouses</option>
                {warehouses.map(warehouse => (
                  <option key={warehouse.id} value={warehouse.id}>
                    {warehouse.name}
                  </option>
                ))}
              </select>
            </div>

            <div className="flex items-end">
              <button
                onClick={() => setFilter({ product_id: undefined, warehouse_id: undefined })}
                className="btn-secondary w-full"
                disabled={isLoading}
              >
                Clear Filters
              </button>
            </div>
          </div>
        </div>
      </div>

      {/* Summary Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="card">
          <div className="card-content">
            <div className="text-center">
              <p className="text-2xl font-bold text-blue-600">{(stockSummary ?? []).length}</p>
              <p className="text-sm text-gray-600">Stock Items</p>
            </div>
          </div>
        </div>

        <div className="card">
          <div className="card-content">
            <div className="text-center">
              <p className="text-2xl font-bold text-green-600">{totals.totalRemain}</p>
              <p className="text-sm text-gray-600">Total Units</p>
            </div>
          </div>
        </div>

        <div className="card">
          <div className="card-content">
            <div className="text-center">
              <p className="text-2xl font-bold text-purple-600">
                {new Set((stockSummary ?? []).map(s => s.product_id)).size}
              </p>
              <p className="text-sm text-gray-600">Unique Products</p>
            </div>
          </div>
        </div>
      </div>

      {/* Stock Summary Table */}
      <div className="card">
        <div className="card-header">
          <h3 className="text-lg font-medium text-gray-900">Stock Summary</h3>
        </div>
        <div className="card-content">
          {isLoading ? (
            <div className="text-center py-8">
              <p className="text-gray-500">Loading stock summary...</p>
            </div>
          ) : (stockSummary ?? []).length === 0 ? (
            <div className="text-center py-8">
              <p className="text-gray-500">No stock data found</p>
            </div>
          ) : (
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Product
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Warehouse
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Primary Unit
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Secondary Unit
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Total
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Last Updated
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Actions
                    </th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {stockSummary.map((item) => (
                    <tr key={`${item.product_id}-${item.warehouse_id}`} className="hover:bg-gray-50">
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div>
                          <div className="text-sm font-medium text-gray-900">
                            {item.product_name}
                          </div>
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-900">{item.warehouse_name}</div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-900">
                          {item.remain1} {item.unit1}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-900">
                          {item.unit2 ? `${item.remain2} ${item.unit2}` : '-'}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm font-medium text-gray-900">
                          {item.total_remain}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-500">
                          {new Date(item.updated_at).toLocaleDateString()}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <button
                          onClick={() => onViewStockCard(item.product_id, item.warehouse_id)}
                          className="text-blue-600 hover:text-blue-900"
                        >
                          View Card
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