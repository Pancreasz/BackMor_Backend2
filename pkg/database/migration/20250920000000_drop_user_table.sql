-- +goose Up
-- +goose StatementBegin
DROP TABLE USER_TABLE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE TABLE USER_TABLE (
   ID                SERIAL PRIMARY KEY,
   USERNAME          VARCHAR(50) UNIQUE NOT NULL,
   NAME              VARCHAR(255) NOT NULL,
   SEX               VARCHAR(255) NOT NULL,
   AGE               INT NOT NULL,
   HASH_PASS         VARCHAR(255) NOT NULL,
   EMAIL             VARCHAR(255) UNIQUE NOT NULL,
   IMAGE_PATH        VARCHAR(500),
   CREATED_TIMESTAMP TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO USER_TABLE (
   USERNAME,
   NAME,
   SEX,
   AGE,
   HASH_PASS,
   EMAIL
) VALUES ( 'karn',
           'suksom',
           'gay',
           99,
           'oasis',
           'liam@noel.com' );
INSERT INTO USER_TABLE (
   USERNAME,
   NAME,
   SEX,
   AGE,
   HASH_PASS,
   EMAIL
) VALUES ( 'pla3',
           'pla2',
           'fish',
           88,
           'blur',
           'song2@gmail.com' );
-- +goose StatementEnd