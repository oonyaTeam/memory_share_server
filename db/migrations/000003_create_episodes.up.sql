CREATE TABLE IF NOT EXISTS episodes(
    id serial PRIMARY KEY,
    episode text,
    longitude double precision,
    latitude double precision,
    memory_id int,
    foreign key (memory_id) references memories(id)
);
