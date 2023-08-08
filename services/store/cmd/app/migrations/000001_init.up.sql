CREATE TABLE "store_book" (
    "id" serial primary key,
    "book_id" integer not null,
    "in_stock" integer,    
    "created_at" timestamp not null,
    "modified_at" timestamp not null
);

CREATE TABLE "store_order" (
    "id" serial primary key,
    "order_id" integer not null,
    "book_id" integer not null,
    "quantity" integer,
    "status" varchar,
    "reason" varchar,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
);