set SOURCES_HOST=localhost
set SOURCES_PORT=5432
set SOURCES_USER=postgres
set SOURCES_PASSWORD=Kk918523
set SOURCES_DATABASE=test
set SOURCES_SSLMODE=disable
set POSTGRESQL_URL=postgres://%SOURCES_USER%:%SOURCES_PASSWORD%@%SOURCES_HOST%:%SOURCES_PORT%/%SOURCES_DATABASE%?sslmode=%SOURCES_SSLMODE%
migrate -database %POSTGRESQL_URL% -path ./migrations up