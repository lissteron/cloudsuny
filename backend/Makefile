include tron.mk

run:
	go run -ldflags "$(LDFLAGS)" \
	--tags "sqlite_app_armor sqlite_foreign_keys sqlite_vacuum_incr sqlite_fts5 sqlite_json sqlite_secure_delete" \
	cmd/cloudsuny/main.go --$(filter-out $@,$(MAKECMDGOALS))

%:
	@true
