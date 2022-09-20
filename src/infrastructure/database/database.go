package database

import (
	"database/sql"
	"log"

	"github.com/mattn/go-sqlite3"
	data "github.com/moohbr/WebMonitor/src/data"
)

// Database is the database struct
type Database struct {
	*sql.DB
}

// NewDatabase creates a new database
func NewDatabase() *Database {
	sqlite3.Version()
	db, err := sql.Open("sqlite3", "file:database.db?cache=shared&mode=rwc&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS servers (name TEXT PRIMARY KEY, ip TEXT, url TEXT, avarageResponseTime TEXT , lastUpdate TEXT , lastCheck TEXT , lastStatus TEXT, monitor BOOLEAN)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (name TEXT PRIMARY KEY, password TEXT, email TEXT, admin BOOLEAN, lastLogin TEXT, lastNotif TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	return &Database{db}
}

func OpenDatabase() *Database {
	db, err := sql.Open("sqlite3", "file:database.db?cache=shared&mode=rwc&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return &Database{db}
}

// AddServer adds a server to the database
func (db *Database) AddServer(s data.Server) {
	_, err := db.Exec("INSERT INTO servers VALUES (?, ?, ?, ?, ?, ?, ?, ?)", s.Name, s.IP, s.URL, s.AvarageResponseTime, s.LastUpdate.Unix(), s.LastCheck.Unix(), s.LastStatus, s.Monitor)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[SYSTEM] Server %s added  to database!", s.Name)
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
	lastupdate := s.LastUpdate.String()
	lastcheck := s.LastCheck.String()
	err := db.QueryRow("SELECT * FROM servers WHERE name=?", name).Scan(&s.Name, &s.IP, &s.URL, &s.AvarageResponseTime,
		&lastupdate, &lastcheck, &s.LastStatus, &s.Monitor)

	if err != nil {
		log.Fatal(err)
	}
	return s
}

// GetUser gets a user from the database
func (db *Database) GetUser(name string) data.User {
	var u data.User
	lastlogin := u.LastLogin.String()
	lastnotif := u.LastNotif.String()
	err := db.QueryRow("SELECT * FROM users WHERE name=?", name).Scan(&u.Name, &u.Password, &u.Email, &u.Admin, &lastlogin, &lastnotif)
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
		lastupdate := s.LastUpdate.String()
		lastcheck := s.LastCheck.String()
		err = rows.Scan(&s.Name, &s.IP, &s.URL, &s.AvarageResponseTime, &lastupdate, &lastcheck, &s.LastStatus, &s.Monitor)

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
		lastlogin := u.LastLogin.String()
		lastnotif := u.LastNotif.String()
		err = rows.Scan(&u.Name, &u.Password, &u.Email, &u.Admin, &lastlogin, &lastnotif)
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
