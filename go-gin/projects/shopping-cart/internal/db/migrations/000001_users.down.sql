drop trigger if exists set_users_updated_at on users;

drop function if exists update_updated_at_column();

drop index if exists idx_users_status;
drop index if exists idx_users_email_status;
drop index if exists idx_users_role;
drop index if exists idx_users_deleted_at_utc;

drop table if exists users;