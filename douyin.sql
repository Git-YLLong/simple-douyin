USE douyin;

-- DROP TABLE IF EXISTS 'user';
CREATE TABLE IF NOT EXISTS user (
    id BIGINT primary key,
    username VARCHAR(25),
    password VARCHAR(25)
);