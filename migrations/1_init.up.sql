CREATE TABLE IF NOT EXISTS users
(
    id           INTEGER PRIMARY KEY,
    email        TEXT    NOT NULL UNIQUE,
    pass_hash    BYTEA    NOT NULL,
    first_name   VARCHAR(100) NOT NULL,
    name VARCHAR(40) NOT NULL,
    last_name VARCHAR(130) NULL,
    phone VARCHAR(20) NULL,
    sex VARCHAR(3) NULL
);
CREATE INDEX IF NOT EXISTS idx_email ON users (email);

CREATE TABLE IF NOT EXISTS apps
(
    id     INTEGER PRIMARY KEY,
    name   TEXT NOT NULL UNIQUE,
    secret TEXT NOT NULL UNIQUE
);