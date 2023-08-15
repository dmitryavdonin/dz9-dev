CREATE TABLE "book" (
    "id" serial primary key,
    "price" integer not null,
    "title" varchar not null,
    "author" varchar not null,
    "created_at" timestamp not null,
    "modified_at" timestamp not null
)