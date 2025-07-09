-- +goose Up
-- +goose StatementBegin
CREATE TABLE vouchers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    crew_name TEXT NOT NULL,
    crew_id TEXT NOT NULL,
    flight_number TEXT NOT NULL,
    flight_date TEXT NOT NULL,
    aircraft_type TEXT NOT NULL,
    seat1 TEXT NOT NULL,
    seat2 TEXT NOT NULL,
    seat3 TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT (datetime('now', 'localtime'))
);

CREATE UNIQUE INDEX idx_voucher_flight_date ON vouchers (flight_number, flight_date);
CREATE INDEX idx_aircraft_type_flight_date ON vouchers(flight_date, aircraft_type)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_aircraft_type_flight_date;
DROP INDEX IF EXISTS idx_voucher_flight_date;
DROP TABLE IF EXISTS vouchers;
-- +goose StatementEnd