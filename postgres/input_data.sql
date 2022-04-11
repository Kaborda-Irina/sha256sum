CREATE TABLE hashFiles
(
    id    SERIAL PRIMARY KEY,
    fileName varchar(100) NOT NULL,
    fullFilePath varchar(100) NOT NULL,
    hashSum  varchar(100) NOT NULL
);