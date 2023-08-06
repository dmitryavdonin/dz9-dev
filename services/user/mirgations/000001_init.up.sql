CREATE TABLE "user" (
    "id" serial primary key,
    "username" varchar not null,
    "balance" integer not null,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
)