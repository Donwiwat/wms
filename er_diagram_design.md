
erDiagram

    PRODUCTS {
        int id PK
        string name
        string short_name
        string brand
        string model
        string size
        string group
        string unit1
        string unit2
        float ratio
        float cost
        string message
        string note
    }

    CUSTOMER_GROUPS {
        int id PK
        string name
        string description
    }

    PRODUCT_PRICES {
        int id PK
        int product_id FK
        string unit
        int customer_group_id FK
        float price
    }

    WAREHOUSES {
        int id PK
        string name
        string location
        string note
    }

    STOCK {
        int id PK
        int product_id FK
        int warehouse_id FK
        int remain1
        int remain2
        int total_remain
        datetime updated_at
    }

    STOCK_MOVEMENTS {
        int id PK
        int product_id FK
        int warehouse_id FK
        string type         
        float qty
        string unit
        int ref_id FK
        string ref_type     
        datetime date
        string note
        string created_by
    }

    SALES_ORDERS {
        int id PK
        string so_number
        datetime date
        string customer
        string note
    }

    DELIVERY_ORDERS {
        int id PK
        string do_number
        datetime date
        string customer
        int warehouse_id FK
        string note
    }

    PURCHASE_ORDERS {
        int id PK
        string po_number
        datetime date
        string supplier
        string note
    }

    GOODS_RECEIPTS {
        int id PK
        string grn_number
        datetime date
        string supplier
        int warehouse_id FK
        string note
    }

    TRANSFERS {
        int id PK
        string tf_number
        int from_warehouse_id FK
        int to_warehouse_id FK
        datetime date
        string note
    }

    STOCK_ADJUSTMENTS {
        int id PK
        string adj_number
        datetime date
        int warehouse_id FK
        string reason
        string note
    }

    %% RELATIONSHIPS

    PRODUCTS ||--o{ PRODUCT_PRICES : "has prices"
    CUSTOMER_GROUPS ||--o{ PRODUCT_PRICES : "price level"

    PRODUCTS ||--o{ STOCK : "has stock"
    WAREHOUSES ||--o{ STOCK : "contains"

    PRODUCTS ||--o{ STOCK_MOVEMENTS : "moved"
    WAREHOUSES ||--o{ STOCK_MOVEMENTS : "movement"

    SALES_ORDERS ||--o{ STOCK_MOVEMENTS : "SO ref"
    DELIVERY_ORDERS ||--o{ STOCK_MOVEMENTS : "DO ref"
    PURCHASE_ORDERS ||--o{ STOCK_MOVEMENTS : "PO ref"
    GOODS_RECEIPTS ||--o{ STOCK_MOVEMENTS : "GRN ref"
    STOCK_ADJUSTMENTS ||--o{ STOCK_MOVEMENTS : "Adjustment ref"
    TRANSFERS ||--o{ STOCK_MOVEMENTS : "Transfer ref"