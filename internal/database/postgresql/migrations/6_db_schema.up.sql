BEGIN;

CREATE TABLE IF NOT EXISTS songs
(
  id                   BIGSERIAL PRIMARY KEY,
  "group"              VARCHAR(255),
  name                 VARCHAR(255),
  release_date         VARCHAR(255),
  text                 VARCHAR(65535),
  link                 VARCHAR(255)
);

COMMIT;