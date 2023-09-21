BEGIN;
DROP TYPE IF EXISTS temperature_style CASCADE;
DROP TABLE IF EXISTS coffee_drink CASCADE;

CREATE TYPE "temperature_style" AS ENUM (
    'HOT',
    'COLD'
);
CREATE TABLE "coffee_drink" (
    "id" uuid PRIMARY KEY,
    "name" varchar(32) UNIQUE NOT NULL,
    "origin" varchar(32),
    "description" text NOT NULL,
    "temperature_style" temperature_style NOT NULL
);
COMMIT;
