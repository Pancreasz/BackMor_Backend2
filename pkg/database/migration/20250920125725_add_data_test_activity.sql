-- +goose Up
-- +goose StatementBegin
INSERT INTO ACTIVITIES (
   CREATOR_ID,
   TITLE,
   DESCRIPTION,
   START_TIME,
   END_TIME,
   MAX_PARTICIPANTS,
   VISIBILITY,
   LATITUDE,
   LONGITUDE,
   LOCATION
) VALUES ( (
   SELECT ID
   FROM USERS
   WHERE EMAIL = 'karn@example.com'
),
           'Weekend Football Match',
           'Friendly 5v5 football game at the park.',
           '2025-09-27 15:00:00+00',
           '2025-09-27 17:00:00+00',
           10,
           'public',
           13.7563,
           100.5018,
           'Bangkok Old Town' ),( (
   SELECT ID
   FROM USERS
   WHERE EMAIL = 'puwit@example.com'
),
                                  'Board Game Night',
                                  'Bring your favorite games and snacks!',
                                  '2025-09-28 18:00:00+00',
                                  NULL,
                                  8,
                                  'friends',
                                  13.7367,
                                  100.5231,
                                  'Bangkok Old Town' ),( (
   SELECT ID
   FROM USERS
   WHERE EMAIL = 'plabid6969420@example.com'
),
                                                         'Morning Cycling Group',
                                                         'Casual ride along the river, all levels welcome.',
                                                         '2025-09-29 06:30:00+00',
                                                         '2025-09-29 08:00:00+00',
                                                         15,
                                                         'public',
                                                         13.7300,
                                                         100.5410,
                                                         'Bangkok Old Town' ),( (
   SELECT ID
   FROM USERS
   WHERE EMAIL = 'karn@example.com'
),
                                                                                'Sunset Photography Walk',
                                                                                'Casual walk through the old town to capture golden hour shots.'
                                                                                ,
                                                                                '2025-09-21 17:30:00+00',
                                                                                '2025-09-21 19:00:00+00',
                                                                                12,
                                                                                'public',
                                                                                13.7563,
                                                                                100.5018,
                                                                                'Bangkok Old Town' );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM ACTIVITIES
WHERE TITLE IN ( 'Weekend Football Match',
                 'Board Game Night',
                 'Morning Cycling Group',
                 'Sunset Photography Walk' );
-- +goose StatementEnd