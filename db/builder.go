package db

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strings"
)

type Client struct { //to be  into main.go
	conn *sql.DB
}

func MustCreate(url, username, password string) *Client {
	conn, err := sql.Open("pgx", completeURL(url, username, password))
	if err != nil {
		log.Fatalf("cannot create database connection: %v", err)
	}
	return &Client{
		conn: conn,
	}
}

func completeURL(databaseURL, username, password string) string {
	databaseURL = databaseURL[0:strings.Index(databaseURL, "//")+2] + "%s@" + databaseURL[strings.Index(databaseURL, "//")+2:]
	if password != "" {
		if strings.Contains(databaseURL, "?") {
			databaseURL += "&password=%s"
		} else {
			databaseURL += "?password=%s"
		}

		return fmt.Sprintf(databaseURL, username, url.QueryEscape(password))
	}
	return fmt.Sprintf(databaseURL, username)
}
