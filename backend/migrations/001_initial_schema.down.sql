-- Drop triggers
DROP TRIGGER IF EXISTS update_stock_total_remain_trigger ON stock;
DROP TRIGGER IF EXISTS update_stock_adjustments_updated_at ON stock_adjustments;
DROP TRIGGER IF EXISTS update_transfers_updated_at ON transfers;
DROP TRIGGER IF EXISTS update_goods_receipts_updated_at ON goods_receipts;
DROP TRIGGER IF EXISTS update_purchase_orders_updated_at ON purchase_orders;
DROP TRIGGER IF EXISTS update_delivery_orders_updated_at ON delivery_orders;
DROP TRIGGER IF EXISTS update_sales_orders_updated_at ON sales_orders;
DROP TRIGGER IF EXISTS update_warehouses_updated_at ON warehouses;
DROP TRIGGER IF EXISTS update_product_prices_updated_at ON product_prices;
DROP TRIGGER IF EXISTS update_products_updated_at ON products;
DROP TRIGGER IF EXISTS update_customer_groups_updated_at ON customer_groups;
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Drop functions
DROP FUNCTION IF EXISTS update_stock_total_remain();
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop indexes
DROP INDEX IF EXISTS idx_product_prices_customer_group;
DROP INDEX IF EXISTS idx_product_prices_product;
DROP INDEX IF EXISTS idx_stock_movements_ref;
DROP INDEX IF EXISTS idx_stock_movements_date;
DROP INDEX IF EXISTS idx_stock_movements_warehouse;
DROP INDEX IF EXISTS idx_stock_movements_product;
DROP INDEX IF EXISTS idx_stock_product_warehouse;

-- Drop tables in reverse order of creation
DROP TABLE IF EXISTS stock_movements;
DROP TABLE IF EXISTS stock_adjustments;
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS goods_receipts;
DROP TABLE IF EXISTS purchase_orders;
DROP TABLE IF EXISTS delivery_orders;
DROP TABLE IF EXISTS sales_orders;
DROP TABLE IF EXISTS stock;
DROP TABLE IF EXISTS warehouses;
DROP TABLE IF EXISTS product_prices;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS customer_groups;
DROP TABLE IF EXISTS users;