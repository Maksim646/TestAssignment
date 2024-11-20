BEGIN;

CREATE TABLE IF NOT EXISTS songs
(
  id BIGINT,
  song_name VARCHAR(255),
  song_text VARCHAR(255)
);

COMMIT;