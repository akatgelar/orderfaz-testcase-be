/*
 Navicat Premium Dump SQL

 Source Server         : PostgreSQL Local
 Source Server Type    : PostgreSQL
 Source Server Version : 140018 (140018)
 Source Host           : localhost:5432
 Source Catalog        : orderfaz_testcase
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140018 (140018)
 File Encoding         : 65001

 Date: 31/07/2025 16:11:29
*/


-- ----------------------------
-- Table structure for auth
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth";
CREATE TABLE "public"."auth" (
  "id" uuid NOT NULL,
  "msisdn" varchar(255) COLLATE "pg_catalog"."default",
  "username" varchar(255) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6),
  "created_by" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."auth" OWNER TO "postgres";

-- ----------------------------
-- Records of auth
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for logistics
-- ----------------------------
DROP TABLE IF EXISTS "public"."logistics";
CREATE TABLE "public"."logistics" (
  "id" uuid NOT NULL,
  "logistic_name" varchar(255) COLLATE "pg_catalog"."default",
  "amount" float4,
  "destination_name" varchar(255) COLLATE "pg_catalog"."default",
  "origin_name" varchar(255) COLLATE "pg_catalog"."default",
  "duration" varchar(255) COLLATE "pg_catalog"."default",
  "is_active" bool,
  "created_at" timestamptz(6),
  "created_by" varchar(255) COLLATE "pg_catalog"."default",
  "updated_at" timestamp(6),
  "updated_by" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."logistics" OWNER TO "postgres";

-- ----------------------------
-- Records of logistics
-- ----------------------------
BEGIN;
INSERT INTO "public"."logistics" ("id", "logistic_name", "amount", "destination_name", "origin_name", "duration", "is_active", "created_at", "created_by", "updated_at", "updated_by") VALUES ('a123ae73-460d-4a6a-8ac8-c0f44ebdf87d', 'JNE', 10000, 'bandung', 'jakarta', '1-2', 'f', '2025-07-31 07:08:40.281031+00', '', '2025-07-31 07:08:40.283457', '');
INSERT INTO "public"."logistics" ("id", "logistic_name", "amount", "destination_name", "origin_name", "duration", "is_active", "created_at", "created_by", "updated_at", "updated_by") VALUES ('7ac531cc-68da-4865-a6db-4524910589af', 'JNT', 11000, 'bandung', 'jakarta', '1-2', 'f', '2025-07-31 07:22:57.127459+00', '', '2025-07-31 07:22:57.130794', '');
INSERT INTO "public"."logistics" ("id", "logistic_name", "amount", "destination_name", "origin_name", "duration", "is_active", "created_at", "created_by", "updated_at", "updated_by") VALUES ('f0bb1ed3-eff0-4617-aeca-ce10788056b1', 'JNT', 21000, 'bandung', 'surabaya', '1-2', 'f', '2025-07-31 07:23:17.162912+00', '', '2025-07-31 07:23:17.164025', '');
COMMIT;

-- ----------------------------
-- Primary Key structure for table auth
-- ----------------------------
ALTER TABLE "public"."auth" ADD CONSTRAINT "auth_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table logistics
-- ----------------------------
ALTER TABLE "public"."logistics" ADD CONSTRAINT "logistics_pkey" PRIMARY KEY ("id");
