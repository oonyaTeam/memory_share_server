-- TODO:いい感じにトランザクション張りたいねえ
insert into authors(uuid) values ('firstUuid'), ('secondUuid'), ('thirdUuid');
insert into memories(memory, image, longitude, latitude, author_id, angle) values ('first memory', 'img1', 12.3, 32.1, 1, 67.7), ('second memory', 'img2', 23.4, 54.6, 1, 89.2);
insert into episodes(episode, longitude, latitude, memory_id) values ('firstEpisode', 43.5, 65.7, 1), ('secondEpisode', 32.4, 54.6, 1), ('thirdEpisode', 554.6, 23.23, 2);
insert into author_seen_memory(memory_id, author_id) values (1, 1),(1, 2), (2, 3);
