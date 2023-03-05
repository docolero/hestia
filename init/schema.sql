CREATE TABLE todolists (
    id BIGSERIAL PRIMARY KEY,
    name text not null,
    created_at DATETIME
);

CREATE TABLE todo (
    id BIGSERIAL PRIMARY KEY,
    title text not null,
    description text,
    priority int,
    state text,
    created_at DATETIME,
    finished_at DATETIME,
    responsible text
)