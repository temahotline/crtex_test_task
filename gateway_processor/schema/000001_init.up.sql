CREATE TABLE users (
    id          SERIAL  PRIMARY KEY,
    first_name  TEXT    NOT NULL,
    last_name   TEXT    NOT NULL,
    balance     INT     NOT NULL CHECK (balance >= 0)
);

CREATE TABLE transactions (
    id                      SERIAL      PRIMARY KEY,
    user_id                 INT         NOT NULL,
    amount                  INT         NOT NULL,
    balance_before          INT         NOT NULL,
    transaction_type        TEXT        NOT NULL    CHECK (transaction_type IN ('deposit', 'withdrawal')),
    created_at              TIMESTAMP   NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id)   REFERENCES  users (id)
);

CREATE INDEX transactions_user_id_idx ON transactions (user_id);
