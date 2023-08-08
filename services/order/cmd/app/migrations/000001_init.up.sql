CREATE TABLE "order" (
    "id" serial primary key,
    "book_id" integer not null,
    "quantity" integer not null,
    "status" varchar,
    "reason" varchar,
    "delivery_address" varchar,
    "delivery_date" timestamp,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
)