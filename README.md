# finch

## Setup

```
git clone https://github.com/lightster/finch.git
docker-compose up
```

In another terminal, manually setup the schema:
```
docker-compose exec postgres psql -U postgres finch
```

Paste in the follow sql into psql:
```sql
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
```

## Run some manual tests

```
http GET http://127.0.0.1:8080/song/803034c3-b3d1-45b8-98c4-cb385aab8c31
http PUT http://127.0.0.1:8080/song/803034c3-b3d1-45b8-98c4-cb385aab8c31 name="New song" genre="Rock" artist="Sam" length:=123 file_path="s3://some/path/to/song" ranking:=2
http PUT http://127.0.0.1:8080/song/803034c3-b3d1-45b8-98c4-cb385aab8c31 name="Sam's Awesome Rock Song"
http POST http://127.0.0.1:8080/song/ name="New pop song" genre="Pop" artist="Matt" length:=137 file_path="s3://some/path/to/pop" ranking:=3
```

# Scope I did not cover in my implementation but needed to

 - Testing: I did not write any automated tests :( I focused on getting something working, thinking I would have time to write tests later. This is probably my biggest regret of this assessment.
 - Authentication: the API developed is currently complete unsecured. At the very least I would have liked to implement basic auth token authentication but I did not get to it
 - Validation: currently I trust the user's input too much. They can pass values of any length or numbers outside of the spec'd range (e.g. ranking can be 10).
 - Song streaming: I implemented a `GET /song/{song_id}` endpoint but not the `GET /stream/{song_id}/` that would return the song's audio binary. If I were to write this, I would have had my Stub return an `io.ReaderCloser` and used `io.Copy(w, readerCloser)` to copy the contents from the stub directly to the `http.ResponseWriter` to eliminate loading a large file into memory before shipping it to the browser.  

# Other considerations

 - I assumed creation and update payloads accept JSON. This was not something that was in the spec. Since I did not ask during the first 30 minutes, I went with JSON because it is basic to implement and often used for APIs.
 - The JSON decode will catch type errors—sometimes too stringently (e.g. `"5"` for a ranking is not accepted)
 - We talked about and I thought about normalizing the database schema to move things like genre and artist out of the `songs` table. Each song can also have multiple ratings—a rating per user, at least—which could benefit from being in another table. Since there are not endpoints for managing artists, genres, users, etc. I chose to keep the database flat for now. I would change this schema for most applications but it could be useful to have it flat in a case such as having an analytics DB.
 - From an environment perspective, I would generally not put the database credentials/URL in docker-compose. It belongs in an environment variable that can be injected.
 - I have no Go experience in a work setting, so my Go could use a lot of feedback. I chose Go because out of Go, Node, Ruby, and Python, I felt most comfortable choosing Go. I know JavaScript really well, but I do not have experience writing web apps and accessing DBs using Node. (I realize in the spec, PHP was acceptable, but I had prepped myself for Go.)
