CREATE TABLE card_sets
(
    id       BIGSERIAL PRIMARY KEY,
    title    VARCHAR        NOT NULL,
    slug     VARCHAR UNIQUE NOT NULL UNIQUE,
    owner_id INT            NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE cards
(
    id          BIGSERIAL PRIMARY KEY,
    question    VARCHAR NOT NULL,
    answer      VARCHAR NOT NULL,
    position    int     NOT NULL,
    card_set_id int     NOT NULL,

    FOREIGN KEY (card_set_id) REFERENCES card_sets (id) ON DELETE CASCADE,
    CONSTRAINT unique_position_per_set UNIQUE (card_set_id, position)
);
