-- Create "tasks" table
CREATE TABLE "public"."tasks" (
  "id" bigserial NOT NULL,
  "title" text NULL,
  "completed" boolean NULL DEFAULT false,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
