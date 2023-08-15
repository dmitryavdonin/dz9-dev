CREATE TABLE "delivery" (
    "id" serial primary key,
    "user_id" integer not null,
    "order_id" integer not null,
    "delivery_address" varchar,
    "delivery_date" timestamp,
    "status" varchar,
    "reason" varchar,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
)