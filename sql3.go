package main

import (
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
)

type TestItem struct {
    Username  string
    Tweetid    int64
}


func InitDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)
    check(err)
    if db == nil { panic("db nil") }
    return db
}

func CreateTable(db *sql.DB) {
    // create table if not exists
    sql_table := `
    CREATE TABLE IF NOT EXISTS tweets(
        username text primary key,
        tweetid int64
    );
    `

    _, err := db.Exec(sql_table)
    check(err)
}

func StoreItem(db *sql.DB, items []TestItem) {
    sql_additem := `
    INSERT OR REPLACE INTO tweets(
        Username,
        Tweetid
    ) values(?, ?)
    `

    stmt, err := db.Prepare(sql_additem)
    check(err)
    defer stmt.Close()

    for _, item := range items {
        _, err2 := stmt.Exec(item.Username, item.Tweetid)
        check(err2)
    }
}

func ReadItem(db *sql.DB, username string) (tweetid int64) {
    sql_readall := `
    SELECT Tweetid FROM tweets WHERE Username = ?
    `

    stmt, err := db.Prepare(sql_readall)
    check(err)
    defer stmt.Close()
    err = stmt.QueryRow(username).Scan(&tweetid)
    check(err)
    return tweetid
}
