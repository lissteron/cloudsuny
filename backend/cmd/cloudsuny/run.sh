#!/bin/sh

rm app.exe 2> /dev/null
go build -v -o app.exe --tags "sqlite_app_armor sqlite_foreign_keys sqlite_vacuum_incr sqlite_fts5 sqlite_json sqlite_secure_delete" main.go 
./app.exe
