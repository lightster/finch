CREATE TYPE genre AS ENUM ('Rock', 'Pop', 'Rap', 'R&B');

CREATE TABLE songs (
    song_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(255) NOT NULL,
    genre genre NOT NULL,
    artist varchar(255) NOT NULL,
    length integer NOT NULL,
    file_path varchar(2056) NOT NULL,
    ranking smallint NOT NULL
);

INSERT INTO songs
(song_id, name, genre, artist, length, file_path, ranking)
VALUES (
    '803034c3-b3d1-45b8-98c4-cb385aab8c31',
    'New song',
    'Rock',
    'Sam',
    123,
    's3://some/path/to/song',
    5
);
