CREATE TABLE "payment" (
    "id" serial primary key,
    "user_id" integer not null,
    "order_id" integer not null,
    "money" integer not null,
    "status" varchar,
    "reason" varchar,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
)