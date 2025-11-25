# API draft

# API Format

- Base URL: `/api/v1/`
- Response: JSON
- Auth: Token หรือ JWT (optional)

---

# ✅ **1) PRODUCTS API**

### **GET /products**

```
200 OK
[
  {
    "id": 1,
    "name": "Screw",
    "short_name": "SCR",
    "brand": "ABC",
    "model": "M8",
    "size": "8mm",
    "group": "Hardware",
    "unit1": "ชิ้น",
    "unit2": "กล่อง",
    "ratio": 100,
    "cost": 1.2,
    "message": "",
    "note": ""
  }
]

```

### **GET /products/{id}**

### **POST /products**

```
{
  "name": "...",
  "short_name": "...",
  "brand": "...",
  "model": "...",
  "size": "...",
  "group": "...",
  "unit1": "ชิ้น",
  "unit2": "ลัง",
  "ratio": 12,
  "cost": 50,
  "message": "",
  "note": ""
}

```

---

# ✅ **2) PRODUCT_PRICES API**

### **GET /products/{id}/prices**

### **POST /product-prices**

```
{
  "product_id": 1,
  "customer_group_id": 2,
  "unit": "ชิ้น",
  "price": 95
}

```

---

# ✅ **3) CUSTOMER GROUPS API**

### **GET /customer-groups**

### **POST /customer-groups**

```
{
  "name": "Dealer A",
  "description": "Large distributor"
}

```

---

# ✅ **4) WAREHOUSES API**

### **GET /warehouses**

### **POST /warehouses**

```
{
  "name": "Main Warehouse",
  "location": "Bangkok",
  "note": ""
}

```

---

# 🟦 **5) STOCK API (SUMMARY + CARD)**

---

### **GET /stock**

**Query:**

- `product_id`
- `warehouse_id`

**Response**

```
[
  {
    "product_id": 1,
    "warehouse_id": 2,
    "remain1": 50,
    "remain2": 3,
    "total_remain": 86,
    "updated_at": "2025-01-21T10:20:00"
  }
]

```

---

### **GET /stock/card**

**Query:**

`product_id`, `warehouse_id`

```
[
  {
    "date": "2025-01-20",
    "type": "IN",
    "qty": 10,
    "unit": "ลัง",
    "ref_type": "PO",
    "ref_id": 12,
    "note": "Received"
  }
]

```

---

# 🟥 **6) STOCK MOVEMENT — IN/OUT/BREAK/PACK/TRANSFER/ADJUST**

ตาราง Stock Movement ใช้ endpoint ตัวเดียว

แต่ type ต่างกัน

## 🎉 Object for POST /stock-movements

```
{
  "product_id": 1,
  "warehouse_id": 2,
  "type": "IN | OUT | BREAK | PACK | TF-IN | TF-OUT | ADJUST",
  "qty": 5,
  "unit": "ชิ้น",
  "ref_type": "PO | DO | GRN | SO | ADJ | TF",
  "ref_id": 12,
  "note": "Some notes"
}

```

---

# 🟨 **7) STOCK IN**

### **POST /stock/in**

```
{
  "product_id": 1,
  "warehouse_id": 2,
  "qty": 10,
  "unit": "ลัง",
  "ref_type": "PO",
  "ref_id": 121,
  "note": "Received"
}

```

---

# 🟧 **8) STOCK OUT**

### **POST /stock/out**

```
{
  "product_id": 1,
  "warehouse_id": 2,
  "qty": 3,
  "unit": "ชิ้น",
  "ref_type": "DO",
  "ref_id": 130,
  "note": "Delivered"
}

```

---

# 🟫 **9) BREAK DOWN (แตกลัง)**

### **POST /stock/break**

```
{
  "product_id": 1,
  "warehouse_id": 2,
  "qty_unit2": 1,
  "note": "Open box for retail"
}

```

---

# 🟪 **10) PACK UP (รวมชิ้นเป็นลัง)**

### **POST /stock/pack**

```
{
  "product_id": 1,
  "warehouse_id": 2,
  "qty_unit2": 2,
  "note": "Pack for delivery"
}

```

---

# 🟩 **11) TRANSFER (TF-IN + TF-OUT)**

### **POST /stock/transfer**

```
{
  "product_id": 1,
  "from_warehouse_id": 1,
  "to_warehouse_id": 2,
  "qty": 5,
  "unit": "ลัง",
  "ref_type": "TF",
  "ref_id": 202501,
  "note": "Transfer to main warehouse"
}

```

ระบบ backend จะสร้าง:

```
TF-OUT (from warehouse)
TF-IN  (to warehouse)

```

---

# 🟧 **12) STOCK ADJUSTMENT**

### **POST /stock/adjust**

```
{
  "product_id": 1,
  "warehouse_id": 2,
  "new_remain1": 40,
  "new_remain2": 3,
  "reason": "Stock Count Correction",
  "note": ""
}

```

---

# 📄 **13) DOCUMENT APIs**

---

## **SO — Sales Orders**

### GET /sales-orders

### POST /sales-orders

```
{
  "so_number": "SO-20250101-001",
  "customer": "John Co.",
  "date": "2025-01-21",
  "note": ""
}

```

---

## **DO — Delivery Orders**

```
POST /delivery-orders
{
  "do_number": "DO-20250101-001",
  "customer": "John Co.",
  "warehouse_id": 1,
  "date": "2025-01-21",
  "note": ""
}

```

---

## **PO — Purchase Orders**

```
POST /purchase-orders
{
  "po_number": "PO-20250101-001",
  "supplier": "ABC Supply",
  "date": "2025-01-21",
  "note": ""
}

```

---

## **GRN — Goods Receipts**

```
POST /goods-receipts
{
  "grn_number": "GRN-20250101-001",
  "supplier": "ABC Supply",
  "warehouse_id": 1,
  "date": "2025-01-21",
  "note": ""
}

```

---

# 🎉 **สรุป Endpoint ทั้งหมด**

## ✔ Products

/products

/products/{id}

/product-prices

/customer-groups

## ✔ Warehouse

/warehouses

## ✔ Stock

/stock

/stock/card

## ✔ Stock Movements

/stock-movements

/stock/in

/stock/out

/stock/break

/stock/pack

/stock/transfer

/stock/adjust

## ✔ Documents

/sales-orders

/delivery-orders

/purchase-orders

/goods-receipts

/transfers

/stock-adjustments