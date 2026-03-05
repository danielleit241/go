create extension if not exists "uuid-ossp";

create table if not exists users (
    id uuid primary key default uuid_generate_v4(),
    name varchar(50) not null,
    age int,
    email varchar(100) not null unique,
    password varchar(60) not null,
    status int not null default 1 check (status in (0, 1)),
    role int not null default 0 check (role in (0, 1, 2)),
    created_at_utc timestamptz not null default now(),
    updated_at_utc timestamptz not null default now(),
    deleted_at_utc timestamptz default null
);

COMMENT ON COLUMN users.status IS '0: inactive, 1: active';
COMMENT ON COLUMN users.role IS '0: user, 1: moderator, 2: admin';
COMMENT ON COLUMN users.deleted_at_utc IS 'Soft delete timestamp, null if not deleted';

CREATE INDEX IF NOT EXISTS idx_users_status ON users (status);
CREATE INDEX IF NOT EXISTS idx_users_email_status ON users (email, status);
CREATE INDEX IF NOT EXISTS idx_users_role ON users (role);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at_utc ON users (deleted_at_utc);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at_utc = now();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER set_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
