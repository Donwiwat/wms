import { useQuery } from 'react-query'
import { stockService } from '@/services/stockService'
import { productService } from '@/services/productService'
import { warehouseService } from '@/services/warehouseService'
import { PackageIcon, WarehouseIcon, BarChart3Icon, TrendingUpIcon } from 'lucide-react'

export default function DashboardPage() {
  const { data: stockSummary, isLoading: stockLoading } = useQuery(
    'stockSummary',
    () => stockService.getStockSummary(),
    { refetchInterval: 30000 }
  )

  const { data: products, isLoading: _productsLoading } = useQuery(
    'products',
    () => productService.getProducts()
  )

  const { data: warehouses, isLoading: _warehousesLoading } = useQuery(
    'warehouses',
    () => warehouseService.getWarehouses()
  )

  const totalProducts = products?.length || 0
  const totalWarehouses = warehouses?.length || 0
  const totalStockItems = stockSummary?.length || 0
  const lowStockItems = stockSummary?.filter(item => item.total_remain < 10).length || 0

  const stats = [
    {
      name: 'Total Products',
      value: totalProducts,
      icon: PackageIcon,
      color: 'bg-blue-500',
    },
    {
      name: 'Total Warehouses',
      value: totalWarehouses,
      icon: WarehouseIcon,
      color: 'bg-green-500',
    },
    {
      name: 'Stock Items',
      value: totalStockItems,
      icon: BarChart3Icon,
      color: 'bg-purple-500',
    },
    {
      name: 'Low Stock Items',
      value: lowStockItems,
      icon: TrendingUpIcon,
      color: 'bg-red-500',
    },
  ]

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Dashboard</h1>
        <p className="text-gray-600">Overview of your warehouse management system</p>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
        {stats.map((stat) => (
          <div key={stat.name} className="card">
            <div className="card-content">
              <div className="flex items-center">
                <div className="flex-shrink-0">
                  <div className={`p-3 rounded-lg ${stat.color}`}>
                    <stat.icon className="h-6 w-6 text-white" />
                  </div>
                </div>
                <div className="ml-5 w-0 flex-1">
                  <dl>
                    <dt className="text-sm font-medium text-gray-500 truncate">
                      {stat.name}
                    </dt>
                    <dd className="text-lg font-medium text-gray-900">
                      {stat.value}
                    </dd>
                  </dl>
                </div>
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Recent Stock Summary */}
      <div className="card">
        <div className="card-header">
          <h3 className="text-lg font-medium text-gray-900">Recent Stock Summary</h3>
        </div>
        <div className="card-content">
          {stockLoading ? (
            <div className="text-center py-4">Loading...</div>
          ) : (
            <div className="overflow-x-auto">
              <table className="table">
                <thead className="table-header">
                  <tr>
                    <th className="table-head">Product</th>
                    <th className="table-head">Warehouse</th>
                    <th className="table-head">Stock (Unit1)</th>
                    <th className="table-head">Stock (Unit2)</th>
                    <th className="table-head">Total</th>
                    <th className="table-head">Status</th>
                  </tr>
                </thead>
                <tbody>
                  {stockSummary?.slice(0, 10).map((item) => (
                    <tr key={`${item.product_id}-${item.warehouse_id}`} className="table-row">
                      <td className="table-cell font-medium">{item.product_name}</td>
                      <td className="table-cell">{item.warehouse_name}</td>
                      <td className="table-cell">{item.remain1} {item.unit1}</td>
                      <td className="table-cell">{item.remain2} {item.unit2}</td>
                      <td className="table-cell">{item.total_remain}</td>
                      <td className="table-cell">
                        <span
                          className={`inline-flex px-2 py-1 text-xs font-semibold rounded-full ${
                            item.total_remain < 10
                              ? 'bg-red-100 text-red-800'
                              : item.total_remain < 50
                              ? 'bg-yellow-100 text-yellow-800'
                              : 'bg-green-100 text-green-800'
                          }`}
                        >
                          {item.total_remain < 10
                            ? 'Low Stock'
                            : item.total_remain < 50
                            ? 'Medium'
                            : 'Good'}
                        </span>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
              {!stockSummary?.length && (
                <div className="text-center py-8 text-gray-500">
                  No stock data available
                </div>
              )}
            </div>
          )}
        </div>
      </div>
    </div>
  )
}