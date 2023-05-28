// schema.sql
drop table if exists todos;
CREATE  TABLE  todos (
    id SERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    status VARCHAR(50),
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);