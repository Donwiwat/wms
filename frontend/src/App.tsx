import { Routes, Route, Navigate } from 'react-router-dom'
import { useAuthStore } from '@/stores/authStore'
import Layout from '@/components/Layout'
import LoginPage from '@/pages/LoginPage'
import DashboardPage from '@/pages/DashboardPage'
import ProductsPage from '@/pages/ProductsPage'
import CustomersPage from '@/pages/CustomersPage'
import OrdersPage from '@/pages/OrdersPage'
import WarehousesPage from '@/pages/WarehousesPage'
import StockPage from '@/pages/StockPage'
import StockMovementsPage from '@/pages/StockMovementsPage'
import StockOperationsPage from '@/pages/StockOperationsPage'

function App() {
  const { isAuthenticated } = useAuthStore()

  if (!isAuthenticated) {
    return (
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route path="*" element={<Navigate to="/login" replace />} />
      </Routes>
    )
  }

  return (
    <Layout>
      <Routes>
        <Route path="/" element={<Navigate to="/dashboard" replace />} />
        <Route path="/dashboard" element={<DashboardPage />} />
        <Route path="/products" element={<ProductsPage />} />
        <Route path="/customers" element={<CustomersPage />} />
        <Route path="/orders" element={<OrdersPage />} />
        <Route path="/warehouses" element={<WarehousesPage />} />
        <Route path="/stock" element={<StockPage />} />
        <Route path="/stock-movements" element={<StockMovementsPage />} />
        <Route path="/stock-operations" element={<StockOperationsPage />} />
        <Route path="/login" element={<Navigate to="/dashboard" replace />} />
        <Route path="*" element={<Navigate to="/dashboard" replace />} />
      </Routes>
    </Layout>
  )
}

export default App