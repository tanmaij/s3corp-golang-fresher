CREATE SCHEMA main;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE main."user"
(
    Username varchar primary key,
    Password varchar NOT NULL,
    Email varchar UNIQUE NOT NULL,
    Name varchar NOT NULL
);

CREATE TABLE main.Document
(
    DocumentId UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    Subject         VARCHAR NOT NULL,
    CreatedAt     TIMESTAMP DEFAULT current_timestamp,
    Username VARCHAR,
    CONSTRAINT FkSubDocument_DocumentId
        FOREIGN KEY (Username) REFERENCES main."user"(Username)  ON DELETE CASCADE
);

CREATE TABLE main.DocumentItem
(
    DocumentItemId UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    Title         VARCHAR NOT NULL,
    CreatedAt     TIMESTAMP DEFAULT current_timestamp,
    Content       VARCHAR NOT NULL,
    DocumentId UUID,
    CONSTRAINT FkSubDocument_DocumentId
        FOREIGN KEY (DocumentId) REFERENCES main.Document(DocumentId) ON DELETE CASCADE
)