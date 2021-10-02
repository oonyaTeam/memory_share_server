CREATE TABLE IF NOT EXISTS author_seen_memory(
    id serial PRIMARY KEY,
    memory_id int,
    author_id int,
    foreign key (memory_id) references memories(id),
    foreign key (author_id) references authors(id)
);
