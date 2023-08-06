CREATE TABLE "store_product" (
    "id" serial primary key,
    "in_stock" integer not null,
    "price" integer not null,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
);

CREATE TABLE "store_order" (
    "id" serial primary key,
    "order_id" integer not null,
    "product_id" integer not null,
    "quantity" integer not null,
    "status" varchar not null,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
);