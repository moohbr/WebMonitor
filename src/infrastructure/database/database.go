package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	data "github.com/moohbr/WebMonitor/src/data"
)

// Database is the database struct
type Database struct {
	*sql.DB
}

// NewDatabase creates a new database
func NewDatabase() *Database {
	db, err := sql.Open("sqlite3", "file:database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	return &Database{db}
}

// InitDatabase creates the tables
func (db *Database) InitDatabase() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS servers (name TEXT PRIMARY KEY, ip TEXT, url TEXT, avarageResponseTime INTEGER, lastUpdate INTEGER, lastCheck INTEGER, lastStatus INTEGER, monitor INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (name TEXT PRIMARY KEY, password TEXT, email TEXT, admin INTEGER, lastLogin INTEGER, lastNotif INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) Connect() {
	err := db.DB.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("[SYSTEM] Already have a database installed!")
	}
}

// AddServer adds a server to the database
func (db *Database) AddServer(s data.Server) {
	_, err := db.Exec("INSERT INTO servers VALUES (?, ?, ?, ?, ?, ?, ?, ?)", s.Name, s.IP, s.URL, s.AvarageResponseTime, s.LastUpdate.Unix(), s.LastCheck.Unix(), s.LastStatus, s.Monitor)
	if err != nil {
		log.Fatal(err)
	}
}

// AddUser adds a user to the database
func (db *Database) AddUser(u data.User) {
	_, err := db.Exec("INSERT INTO users VALUES (?, ?, ?, ?, ?, ?)", u.Name, u.Password, u.Email, u.Admin, u.LastLogin.Unix(), u.LastNotif.Unix())
	if err != nil {
		log.Fatal(err)
	}
}

// GetServer gets a server from the database
func (db *Database) GetServer(name string) data.Server {
	var s data.Server
	err := db.QueryRow("SELECT * FROM servers WHERE name=?", name).Scan(&s.Name, &s.IP, &s.URL, &s.AvarageResponseTime, &s.LastUpdate, &s.LastCheck, &s.LastStatus, &s.Monitor)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

// GetUser gets a user from the database
func (db *Database) GetUser(name string) data.User {
	var u data.User
	err := db.QueryRow("SELECT * FROM users WHERE name=?", name).Scan(&u.Name, &u.Password, &u.Email, &u.Admin, &u.LastLogin, &u.LastNotif)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

// GetServers gets all servers from the database
func (db *Database) GetServers() []data.Server {
	rows, err := db.Query("SELECT * FROM servers")
	if err != nil {
		log.Fatal(err)
	}
	var servers []data.Server
	for rows.Next() {
		var s data.Server
		err = rows.Scan(&s.Name, &s.IP, &s.URL, &s.AvarageResponseTime, &s.LastUpdate, &s.LastCheck, &s.LastStatus, &s.Monitor)
		if err != nil {
			log.Fatal(err)
		}
		servers = append(servers, s)
	}
	return servers
}

// GetUsers gets all users from the database
func (db *Database) GetUsers() []data.User {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	var users []data.User
	for rows.Next() {
		var u data.User
		err = rows.Scan(&u.Name, &u.Password, &u.Email, &u.Admin, &u.LastLogin, &u.LastNotif)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	return users
}

// UpdateServer updates a server in the database
func (db *Database) UpdateServer(s data.Server) {
	_, err := db.Exec("UPDATE servers SET ip=?, url=?, avarageResponseTime=?, lastUpdate=?, lastCheck=?, lastStatus=?, monitor=? WHERE name=?", s.IP, s.URL, s.AvarageResponseTime, s.LastUpdate.Unix(), s.LastCheck.Unix(), s.LastStatus, s.Monitor, s.Name)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateUser updates a user in the database
func (db *Database) UpdateUser(u data.User) {
	_, err := db.Exec("UPDATE users SET password=?, email=?, admin=?, lastLogin=?, lastNotif=? WHERE name=?", u.Password, u.Email, u.Admin, u.LastLogin.Unix(), u.LastNotif.Unix(), u.Name)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteServer deletes a server from the database
func (db *Database) DeleteServer(name string) {
	_, err := db.Exec("DELETE FROM servers WHERE name=?", name)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteUser deletes a user from the database
func (db *Database) DeleteUser(name string) {
	_, err := db.Exec("DELETE FROM users WHERE name=?", name)
	if err != nil {
		log.Fatal(err)
	}
}

// Close closes the database
func (db *Database) Close() {
	db.DB.Close()
}
