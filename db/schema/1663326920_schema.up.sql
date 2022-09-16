BEGIN;
CREATE TABLE IF NOT EXISTS user_profile(
    id                      bigserial                                              PRIMARY KEY,
    name                    VARCHAR(255)                                        NOT NULL,
    surname                 VARCHAR(255)                                        NOT NULL,
    born                    TIMESTAMP                                           NOT NULL,
    phone_number            VARCHAR(255) UNIQUE                                 NOT NULL,
    email                   VARCHAR(255) UNIQUE                                 NOT NULL,
    gender                  TEXT CHECK(gender IN ('Male','Female'))             NOT NULL,
    address                 TEXT                                                NOT NULL,
    password                TEXT                                                NOT NULL,
    created_at              TIMESTAMP                                           NOT NULL
);
COMMIT;