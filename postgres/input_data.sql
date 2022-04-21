CREATE TABLE hashFiles
(
    id    BIGSERIAL PRIMARY KEY,
    fileName VARCHAR NOT NULL,
    fullFilePath TEXT NOT NULL ,
    algorithm VARCHAR NOT NULL,
    hashSum  VARCHAR NOT NULL,
    deleted BOOLEAN DEFAULT false

);
CREATE UNIQUE INDEX path_alg_idx ON hashFiles (fullFilePath,algorithm);