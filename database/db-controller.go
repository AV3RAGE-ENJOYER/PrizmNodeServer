package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Sql *sql.DB
}

func NewDb(dbFile string) (*Database, error) {
	sqlDB, err := sql.Open("sqlite3", dbFile)

	if err != nil {
		return nil, err
	}

	database := Database{
		Sql: sqlDB,
	}

	return &database, nil
}

func (db *Database) GetApiKeys() ([]ApiKey, error) {
	rows, err := db.Sql.Query("SELECT * FROM `auth`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []ApiKey

	for rows.Next() {
		key := ApiKey{}

		err := rows.Scan(
			&key.ApiKey,
			&key.AllowedIPs,
			&key.Name)

		if err != nil {
			return nil, err
		}

		results = append(results, key)
	}

	return results, nil
}

func (db *Database) GetClientById(id int) (*Client, error) {
	row := db.Sql.QueryRow("SELECT * FROM `clients` WHERE `id` = ?", id)

	client := Client{}

	err := row.Scan(
		&client.Id,
		&client.PublicKey,
		&client.WireguardConfig,
		&client.ExpiryDate)

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (db *Database) GetClients() ([]Client, error) {
	rows, err := db.Sql.Query("SELECT * FROM `clients`")

	if err != nil {
		return nil, err
	}

	var clients []Client

	for rows.Next() {
		client := Client{}

		rows.Scan(
			&client.Id,
			&client.PublicKey,
			&client.WireguardConfig,
			&client.ExpiryDate,
		)

		clients = append(clients, client)
	}

	return clients, nil
}

func (db *Database) AddClient(id int, expiryDate int) (sql.Result, error) {
	r, err := db.Sql.Exec(
		"INSERT INTO `clients` VALUES (?, ?, ?, ?)",
		id, "test", "test", expiryDate)

	return r, err
}

func (db *Database) UpdateExpiryDate(id, expiryDate int) (sql.Result, error) {
	r, err := db.Sql.Exec(
		"UPDATE `clients` set `expiry_date` = ? WHERE `id` = ?",
		expiryDate, id,
	)

	return r, err
}

func (db *Database) Close() error {
	defer db.Sql.Close()

	return nil
}
