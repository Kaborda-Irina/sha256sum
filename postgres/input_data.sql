CREATE TABLE hashFiles
(
    id    BIGSERIAL PRIMARY KEY,
    fileName VARCHAR NOT NULL,
    fullFilePath TEXT NOT NULL ,
    algorithm VARCHAR NOT NULL,
    hashSum  VARCHAR NOT NULL
);

CREATE FUNCTION checkHashSum (VARCHAR, TEXT, VARCHAR, VARCHAR) RETURNS void
AS $$
BEGIN
    IF EXISTS (SELECT * FROM hashFiles
                   WHERE fullFilePath=$2
                     AND algorithm=$4)
    THEN
        UPDATE hashFiles SET hashSum=$3 WHERE fullFilePath=$2 AND algorithm=$4;
    ELSE
        INSERT INTO hashFiles (fileName,fullFilePath,hashSum, algorithm)
        VALUES ($1, $2, $3, $4);
    END IF;
END;
$$ LANGUAGE plpgsql;
