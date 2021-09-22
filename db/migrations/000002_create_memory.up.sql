CREATE TABLE IF NOT EXISTS memories(
    id serial PRIMARY KEY,
    memory text,
    image text,
    longitude double precision,
    latitude double precision,
    author_id int,
    angle double precision,
    foreign key (author_id) references authors(id)
);
