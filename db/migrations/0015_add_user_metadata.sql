-- +migrate Up
ALTER TABLE authd_user ADD COLUMN "metadata" text;
