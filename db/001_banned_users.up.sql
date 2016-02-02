create table banned_users (user_id varchar, channel_id varchar, banned_by varchar, banned_on timestamp with time zone, primary key(user_id, channel_id));
