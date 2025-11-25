import React, { useState } from 'react'
import { 
  SalesOrder, 
  DeliveryOrder, 
  PurchaseOrder, 
  GoodsReceipt, 
  Transfer, 
  StockAdjustment 
} from '../types'

type DocumentType = 'SO' | 'DO' | 'PO' | 'GRN' | 'TF' | 'ADJ'
type Document = SalesOrder | DeliveryOrder | PurchaseOrder | GoodsReceipt | Transfer | StockAdjustment

interface DocumentListProps {
  documentType: DocumentType
  documents: Document[]
  onView: (document: Document) => void
  onEdit: (document: Document) => void
  onDelete: (document: Document) => void
  onCreate: () => void
  isLoading?: boolean
}

export default function DocumentList({ 
  documentType, 
  documents, 
  onView, 
  onEdit, 
  onDelete, 
  onCreate,
  isLoading = false 
}: DocumentListProps) {
  const [searchTerm, setSearchTerm] = useState('')
  const [sortField, setSortField] = useState<string>('date')
  const [sortDirection, setSortDirection] = useState<'asc' | 'desc'>('desc')

  const getDocumentTypeInfo = (type: DocumentType) => {
    switch (type) {
      case 'SO':
        return { title: 'Sales Orders', icon: '📋', color: 'blue' }
      case 'DO':
        return { title: 'Delivery Orders', icon: '🚚', color: 'green' }
      case 'PO':
        return { title: 'Purchase Orders', icon: '🛒', color: 'purple' }
      case 'GRN':
        return { title: 'Goods Receipts', icon: '📦', color: 'indigo' }
      case 'TF':
        return { title: 'Transfers', icon: '🔄', color: 'orange' }
      case 'ADJ':
        return { title: 'Stock Adjustments', icon: '⚖️', color: 'red' }
      default:
        return { title: 'Documents', icon: '📄', color: 'gray' }
    }
  }

  const getDocumentNumber = (doc: Document) => {
    if ('so_number' in doc) return doc.so_number
    if ('do_number' in doc) return doc.do_number
    if ('po_number' in doc) return doc.po_number
    if ('grn_number' in doc) return doc.grn_number
    if ('tf_number' in doc) return doc.tf_number
    if ('adj_number' in doc) return doc.adj_number
    return 'N/A'
  }

  const getDocumentDescription = (doc: Document) => {
    if ('customer' in doc) return doc.customer
    if ('supplier' in doc) return doc.supplier
    if ('from_warehouse_id' in doc) return `Transfer from WH${doc.from_warehouse_id} to WH${doc.to_warehouse_id}`
    if ('reason' in doc) return doc.reason
    return 'N/A'
  }

  const filteredDocuments = documents.filter(doc => {
    const number = getDocumentNumber(doc).toLowerCase()
    const description = getDocumentDescription(doc).toLowerCase()
    const search = searchTerm.toLowerCase()
    return number.includes(search) || description.includes(search)
  })

  const sortedDocuments = [...filteredDocuments].sort((a, b) => {
    let aValue: any = a[sortField as keyof Document]
    let bValue: any = b[sortField as keyof Document]

    if (sortField === 'number') {
      aValue = getDocumentNumber(a)
      bValue = getDocumentNumber(b)
    } else if (sortField === 'description') {
      aValue = getDocumentDescription(a)
      bValue = getDocumentDescription(b)
    }

    if (aValue < bValue) return sortDirection === 'asc' ? -1 : 1
    if (aValue > bValue) return sortDirection === 'asc' ? 1 : -1
    return 0
  })

  const handleSort = (field: string) => {
    if (sortField === field) {
      setSortDirection(sortDirection === 'asc' ? 'desc' : 'asc')
    } else {
      setSortField(field)
      setSortDirection('asc')
    }
  }

  const typeInfo = getDocumentTypeInfo(documentType)

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <div className="flex items-center space-x-3">
          <span className="text-2xl">{typeInfo.icon}</span>
          <div>
            <h1 className="text-2xl font-bold text-gray-900">{typeInfo.title}</h1>
            <p className="text-gray-600">Manage {typeInfo.title.toLowerCase()}</p>
          </div>
        </div>
        <button
          onClick={onCreate}
          className="btn-primary"
          disabled={isLoading}
        >
          Create New
        </button>
      </div>

      {/* Search and Filters */}
      <div className="card">
        <div className="card-content">
          <div className="flex flex-col md:flex-row gap-4">
            <div className="flex-1">
              <input
                type="text"
                placeholder="Search by document number or description..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="input w-full"
                disabled={isLoading}
              />
            </div>
            <div className="flex space-x-2">
              <select
                value={sortField}
                onChange={(e) => setSortField(e.target.value)}
                className="input"
                disabled={isLoading}
              >
                <option value="date">Sort by Date</option>
                <option value="number">Sort by Number</option>
                <option value="description">Sort by Description</option>
              </select>
              <button
                onClick={() => setSortDirection(sortDirection === 'asc' ? 'desc' : 'asc')}
                className="btn-secondary px-3"
                disabled={isLoading}
              >
                {sortDirection === 'asc' ? '↑' : '↓'}
              </button>
            </div>
          </div>
        </div>
      </div>

      {/* Document List */}
      <div className="card">
        <div className="card-header">
          <h3 className="text-lg font-medium text-gray-900">
            {typeInfo.title} ({sortedDocuments.length})
          </h3>
        </div>
        <div className="card-content">
          {isLoading ? (
            <div className="text-center py-8">
              <p className="text-gray-500">Loading documents...</p>
            </div>
          ) : sortedDocuments.length === 0 ? (
            <div className="text-center py-8">
              <p className="text-gray-500">
                {searchTerm ? 'No documents found matching your search' : 'No documents found'}
              </p>
            </div>
          ) : (
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th 
                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer hover:bg-gray-100"
                      onClick={() => handleSort('number')}
                    >
                      Document Number
                      {sortField === 'number' && (
                        <span className="ml-1">{sortDirection === 'asc' ? '↑' : '↓'}</span>
                      )}
                    </th>
                    <th 
                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer hover:bg-gray-100"
                      onClick={() => handleSort('date')}
                    >
                      Date
                      {sortField === 'date' && (
                        <span className="ml-1">{sortDirection === 'asc' ? '↑' : '↓'}</span>
                      )}
                    </th>
                    <th 
                      className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider cursor-pointer hover:bg-gray-100"
                      onClick={() => handleSort('description')}
                    >
                      Description
                      {sortField === 'description' && (
                        <span className="ml-1">{sortDirection === 'asc' ? '↑' : '↓'}</span>
                      )}
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Note
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Created
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                      Actions
                    </th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {sortedDocuments.map((doc) => (
                    <tr key={doc.id} className="hover:bg-gray-50">
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm font-medium text-gray-900">
                          {getDocumentNumber(doc)}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-900">
                          {new Date(doc.date).toLocaleDateString()}
                        </div>
                      </td>
                      <td className="px-6 py-4">
                        <div className="text-sm text-gray-900 max-w-xs truncate">
                          {getDocumentDescription(doc)}
                        </div>
                      </td>
                      <td className="px-6 py-4">
                        <div className="text-sm text-gray-500 max-w-xs truncate" title={doc.note}>
                          {doc.note || '-'}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="text-sm text-gray-500">
                          {new Date(doc.created_at).toLocaleDateString()}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm font-medium space-x-2">
                        <button
                          onClick={() => onView(doc)}
                          className="text-blue-600 hover:text-blue-900"
                        >
                          View
                        </button>
                        <button
                          onClick={() => onEdit(doc)}
                          className="text-indigo-600 hover:text-indigo-900"
                        >
                          Edit
                        </button>
                        <button
                          onClick={() => onDelete(doc)}
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