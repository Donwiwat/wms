export default function StockMovementsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-2xl font-bold text-gray-900">Stock Movements</h1>
        <p className="text-gray-600">View all stock movement history</p>
      </div>

      <div className="card">
        <div className="card-header">
          <h3 className="text-lg font-medium text-gray-900">Movement History</h3>
        </div>
        <div className="card-content">
          <div className="text-center py-8 text-gray-500">
            Stock movements interface will be implemented here.
            <br />
            Features: Filter movements, view transaction history, export reports.
          </div>
        </div>
      </div>
    </div>
  )
}