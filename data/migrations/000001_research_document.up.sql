CREATE SCHEMA main;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE main.Document
(
    DocumentId UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    Subject         VARCHAR,
    CreatedAt     TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE main.SubDocument
(
    SubDocumentId UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    Title         VARCHAR,
    CreatedAt     TIMESTAMP DEFAULT current_timestamp,
    Content VARCHAR,
    DocumentId UUID,
    CONSTRAINT FkSubDocument_DocumentId
        FOREIGN KEY (DocumentId) REFERENCES main.Document(DocumentId)
)