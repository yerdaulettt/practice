create table if not exists user_friends (
    user_id integer references users(id) on delete cascade,
    friend_id integer references users(id) on delete cascade,
    primary key (user_id, friend_id),
    constraint self_friend_deny check(user_id <> friend_id)
);