BEGIN;

CREATE TABLE IF NOT EXISTS songs
(
  id                   BIGSERIAL PRIMARY KEY,
  "group"              VARCHAR(255),
  song                 VARCHAR(255)
);

COMMIT;