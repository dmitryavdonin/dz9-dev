CREATE TABLE "order" (
    "id" serial primary key,
    "product_id" integer not null,
    "quantity" integer not null,
    "price" integer not null,
    "status" varchar not null,
    "delivery_address" varchar not null,
    "delivery_date" timestamp not null,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
)