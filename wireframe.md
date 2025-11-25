# **1) Product Form — เพิ่มสินค้า**

```
+-----------------------------------------------------+
|                   PRODUCT FORM                      |
+-----------------------------------------------------+

[ Product Name           ]   [ Short Name         ]

[ Brand                 ]    [ Model             ]

[ Size                  ]    [ Group             ]

-------------------------------------------------------
 Units
-------------------------------------------------------
[ Unit1 (Base unit)   ]  e.g. ชิ้น
[ Unit2 (Pack unit)   ]  e.g. ลัง
[ Ratio (Unit2 -> Unit1) ]  e.g. 1 ลัง = 12 ชิ้น

-------------------------------------------------------
 Cost & Notes
-------------------------------------------------------
[ Cost                 ]
[ Message (optional)   ]
[ Note                 ]

( Save )   ( Cancel )

```

# **2) Product Price Form — ตั้งราคา**

```
+----------------------------------------------------+
|                 PRODUCT PRICE FORM                 |
+----------------------------------------------------+

[ Product                   v ]

------------------------------------------------------
 Price per Customer Group
------------------------------------------------------

[ Customer Group   v ]   [ Unit (unit1/unit2)  v ]
[ Price                           ]

( + Add Price Row )

( Save )   ( Cancel )
```

# **3) Warehouse Form**

```
+----------------------------------------------------+
|                 WAREHOUSE FORM                     |
+----------------------------------------------------+

[ Warehouse Name         ]
[ Location               ]
[ Note                   ]

( Save )   ( Cancel )

```

# **4) Stock Summary Page**

```
+-------------------------------------------------------------+
|                      STOCK SUMMARY                          |
+-------------------------------------------------------------+

Filter:
[ Product     v ]  [ Warehouse   v ]   [ Search ]

----------------------------------------------------------------
| Product | Warehouse | Remain (unit1) | Remain (unit2) | Total |
----------------------------------------------------------------
| ...     | ...       | ...            | ...             | ...   |
| ...     | ...       | ...            | ...             | ...   |
----------------------------------------------------------------

```

# **5) Stock Card (ประวัติ Movement)**

```
+-------------------------------------------------------------+
|                        STOCK CARD                           |
+-------------------------------------------------------------+

Filter:
[ Product   v ]   [ Warehouse   v ]    [ Date Range     ]

----------------------------------------------------------------------
| Date | Type | Qty | Unit | Ref Doc | Note | Created By |
----------------------------------------------------------------------
| ...  | IN   | 10  | ลัง  | PO-001  | ...  | admin      |
| ...  | OUT  | 3   | ชิ้น | DO-005  | ...  | staff      |
| ...  | BREAK| 1   | ลัง  | BD-002  | ...  | staff      |
----------------------------------------------------------------------

```

# **6) Stock IN Form**

```
+-------------------------------------------------------------+
|                        STOCK IN FORM                        |
+-------------------------------------------------------------+

[ Product             v ]
[ Warehouse           v ]
[ Quantity            ]    [ Unit   v ]  (ชิ้น/ลัง)
[ Ref Doc             ]    e.g. PO-2025-001
[ Note                ]

( Save )   ( Cancel )

```

# **7) Stock OUT Form**

```
+-------------------------------------------------------------+
|                       STOCK OUT FORM                        |
+-------------------------------------------------------------+

[ Product             v ]
[ Warehouse           v ]
[ Quantity            ]    [ Unit   v ]
[ Ref Doc             ]    e.g. DO-2025-010
[ Note                ]

( Save )   ( Cancel )

```

# **8) Break Down Form (แตกลังเป็นชิ้น)**

```
+----------------------------------------------------------------+
|                      BREAK DOWN FORM                            |
+----------------------------------------------------------------+

[ Product                      v ]
[ Warehouse                    v ]
[ จำนวนลังที่ต้องการแตก (Unit2) ]   e.g. 1

คุณจะได้เพิ่ม = (Ratio x จำนวนลัง)

[ Note ]

( Execute Break Down )
```

# **9) Pack Up Form (รวมชิ้นเป็นลัง)**

```
+----------------------------------------------------------------+
|                        PACK UP FORM                            |
+----------------------------------------------------------------+

[ Product                    v ]
[ Warehouse                  v ]
[ จำนวนลังที่ต้องการสร้าง  ]  e.g. 2

ระบบจะคำนวณ: ต้องใช้ remain1 = 2 * ratio

[ Note ]

( Execute Pack Up )

```

# **10) Transfer Form (โอนคลัง)**

```
+-------------------------------------------------------------+
|                        TRANSFER FORM                        |
+-------------------------------------------------------------+

[ Product               v ]

From:
[ From Warehouse        v ]

To:
[ To Warehouse          v ]

[ Quantity              ]   [ Unit  v ]

[ Ref Doc (TF)          ]
[ Note                  ]

( Execute Transfer )
```

# **11) Stock Adjustment Form**

```
+-------------------------------------------------------------+
|                    STOCK ADJUSTMENT FORM                    |
+-------------------------------------------------------------+

[ Product                      v ]
[ Warehouse                    v ]

Current:
Remain1 = xx
Remain2 = yy

New:
[ New Remain1      ]
[ New Remain2      ]

[ Reason            ]
[ Note              ]

( Adjust Stock )

```

# **12) Document Forms (SO / DO / PO / GRN)**

## 📄 Sales Order (SO)

```
SO Number
Date
Customer
Note
( Item Table… )

```

## 📄 Delivery Order (DO)

```
DO Number
Date
Customer
Warehouse
Note
( Item Table… )

```

## 📄 Purchase Order (PO)

```
PO Number
Date
Supplier
Note
( Item Table… )

```

## 📄 Goods Receipt (GRN)

```
GRN Number
Date
Supplier
Warehouse
Note
( Item Table… )

```