// API Response Types
export interface ApiResponse<T> {
  data?: T
  error?: string
  message?: string
}

// User Types
export interface User {
  id: number
  username: string
  email: string
  role: string
  is_active: boolean
  created_at: string
  updated_at: string
}

export interface LoginRequest {
  username: string
  password: string
}

export interface RegisterRequest {
  username: string
  email: string
  password: string
  role?: string
}

export interface AuthResponse {
  token: string
  user: User
}

// Product Types
export interface Product {
  id: number
  name: string
  short_name: string
  brand: string
  model: string
  size: string
  group: string
  unit1: string
  unit2: string
  ratio: number
  cost: number
  message: string
  note: string
  created_at: string
  updated_at: string
}

// Customer Group Types
export interface CustomerGroup {
  id: number
  name: string
  description: string
  created_at: string
  updated_at: string
}

// Product Price Types
export interface ProductPrice {
  id: number
  product_id: number
  customer_group_id: number
  unit: string
  price: number
  created_at: string
  updated_at: string
}

// Warehouse Types
export interface Warehouse {
  id: number
  name: string
  location: string
  note: string
  created_at: string
  updated_at: string
}

// Stock Types
export interface Stock {
  id: number
  product_id: number
  warehouse_id: number
  remain1: number
  remain2: number
  total_remain: number
  updated_at: string
}

export interface StockSummary {
  product_id: number
  product_name: string
  warehouse_id: number
  warehouse_name: string
  remain1: number
  remain2: number
  total_remain: number
  unit1: string
  unit2: string
  updated_at: string
}

export interface StockCardEntry {
  date: string
  type: string
  qty: number
  unit: string
  ref_type: string
  ref_id?: number
  note: string
  created_by: string
}

// Stock Movement Types
export interface StockMovement {
  id: number
  product_id: number
  warehouse_id: number
  type: string
  qty: number
  unit: string
  ref_id?: number
  ref_type: string
  date: string
  note: string
  created_by: string
  created_at: string
}

// Stock Operation Request Types
export interface StockInRequest {
  product_id: number
  warehouse_id: number
  qty: number
  unit: string
  ref_type?: string
  ref_id?: number
  note?: string
}

export interface StockOutRequest {
  product_id: number
  warehouse_id: number
  qty: number
  unit: string
  ref_type?: string
  ref_id?: number
  note?: string
}

export interface BreakDownRequest {
  product_id: number
  warehouse_id: number
  qty_unit2: number
  note?: string
}

export interface PackUpRequest {
  product_id: number
  warehouse_id: number
  qty_unit2: number
  note?: string
}

export interface TransferRequest {
  product_id: number
  from_warehouse_id: number
  to_warehouse_id: number
  qty: number
  unit: string
  ref_type?: string
  ref_id?: number
  note?: string
}

export interface StockAdjustRequest {
  product_id: number
  warehouse_id: number
  new_remain1: number
  new_remain2: number
  reason: string
  note?: string
}

// Document Types
export interface SalesOrder {
  id: number
  so_number: string
  date: string
  customer: string
  note: string
  created_at: string
  updated_at: string
}

export interface DeliveryOrder {
  id: number
  do_number: string
  date: string
  customer: string
  warehouse_id: number
  note: string
  created_at: string
  updated_at: string
}

export interface PurchaseOrder {
  id: number
  po_number: string
  date: string
  supplier: string
  note: string
  created_at: string
  updated_at: string
}

export interface GoodsReceipt {
  id: number
  grn_number: string
  date: string
  supplier: string
  warehouse_id: number
  note: string
  created_at: string
  updated_at: string
}

export interface Transfer {
  id: number
  tf_number: string
  from_warehouse_id: number
  to_warehouse_id: number
  date: string
  note: string
  created_at: string
  updated_at: string
}

export interface StockAdjustment {
  id: number
  adj_number: string
  date: string
  warehouse_id: number
  reason: string
  note: string
  created_at: string
  updated_at: string
}

// Form Types
export interface ProductFormData {
  name: string
  short_name: string
  brand: string
  model: string
  size: string
  group: string
  unit1: string
  unit2: string
  ratio: number
  cost: number
  message: string
  note: string
}

export interface WarehouseFormData {
  name: string
  location: string
  note: string
}

export interface CustomerGroupFormData {
  name: string
  description: string
}

// Filter Types
export interface StockFilter {
  product_id?: number
  warehouse_id?: number
}

export interface StockMovementFilter {
  product_id?: number
  warehouse_id?: number
  limit?: number
  offset?: number
}

// Customer Types
export interface Customer {
  id: number
  prefix: string
  name: string
  address: string
  phone: string
  contact_person: string
  level: string
  delivery_place: string
  transport: string
  credit_limit: number
  credit_term: number
  outstanding: number
  last_contact?: string
  note: string
  created_at: string
  updated_at: string
}

export interface CustomerFormData {
  prefix: string
  name: string
  address: string
  phone: string
  contact_person: string
  level: string
  delivery_place: string
  transport: string
  credit_limit: number
  credit_term: number
  outstanding: number
  last_contact?: string
  note: string
}

// Order Types
export interface Order {
  id: number
  order_number: string
  customer_id: number
  order_date: string
  delivery_date?: string
  status: string
  total_amount: number
  discount: number
  tax_amount: number
  final_amount: number
  payment_terms: string
  delivery_address: string
  note: string
  created_by: string
  created_at: string
  updated_at: string
}

export interface OrderItem {
  id: number
  order_id: number
  product_id: number
  quantity: number
  unit: string
  unit_price: number
  total_price: number
  note: string
  created_at: string
}

export interface OrderItemWithProduct {
  order_item: OrderItem
  product_name: string
}

export interface OrderWithDetails {
  order: Order
  customer: Customer
  items: OrderItemWithProduct[]
}

export interface OrderRequest {
  customer_id: number
  order_date: string
  delivery_date?: string
  payment_terms: string
  delivery_address: string
  note: string
  items: OrderItemRequest[]
}

export interface OrderItemRequest {
  product_id: number
  quantity: number
  unit: string
  unit_price: number
  note: string
}

export interface OrderFormData {
  customer_id: number
  order_date: string
  delivery_date?: string
  payment_terms: string
  delivery_address: string
  note: string
  items: OrderItemRequest[]
}