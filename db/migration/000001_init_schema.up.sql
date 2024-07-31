CREATE TABLE "device" (
  "id" bigserial PRIMARY KEY,
  "device_name" varchar NOT NULL,
  "room_id" bigserial,
  "user_id" bigserial
);

CREATE TABLE "room" (
  "id" bigserial PRIMARY KEY,
  "room_name" varchar NOT NULL
);

CREATE TABLE "reading" (
  "id" bigserial PRIMARY KEY,
  "reading_wattage" decimal NOT NULL,
  "reading_hour" timestamptz NOT NULL DEFAULT (now()),
  "device_id" bigserial
);

CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "user_email" varchar NOT NULL,
  "user_password" varchar NOT NULL
);

CREATE INDEX ON "device" ("user_id");

CREATE INDEX ON "device" ("room_id");

CREATE INDEX ON "reading" ("device_id");

ALTER TABLE "device" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("id");

ALTER TABLE "device" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "reading" ADD FOREIGN KEY ("device_id") REFERENCES "device" ("id");
