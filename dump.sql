--
-- File generated with SQLiteStudio v3.4.4 on Вс фев 4 01:43:06 2024
--
-- Text encoding used: System
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: auth
CREATE TABLE IF NOT EXISTS auth (api_key TEXT, allowed_ips TEXT, name TEXT);

-- Table: clients
CREATE TABLE IF NOT EXISTS clients (id INTEGER, public_key TEXT, wireguard_config TEXT, expiry_date INTEGER);

COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
