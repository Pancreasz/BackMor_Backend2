-- +goose Up
-- +goose StatementBegin
INSERT INTO USERS (
   EMAIL,
   PASSWORD_HASH,
   DISPLAY_NAME,
   AVATAR_URL,
   BIO
) VALUES ( 'karn@example.com',
           'hashed_password_1',
           'Suksom',
           NULL,
           'gaymer' ),( 'Puwit@example.com',
                        'hashed_password_2',
                        'pu_gay',
                        'https://example.com/avatar.jpg',
                        'I love gaming and coding!' ),( 'plabid6969420@example.com',
                                                        'hashed_password_3',
                                                        'Plabid',
                                                        'https://example.com/avatar.png',
                                                        'คนขี้บิดแห่งชาติ' );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM USERS
WHERE EMAIL IN ( 'karn@example.com',
                 'Puwit@example.com' );
-- +goose StatementEnd