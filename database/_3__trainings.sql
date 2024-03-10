CREATE TABLE trainings
(
    id            BIGSERIAL PRIMARY KEY,
    user_id       INT,
    card_set_id   INT,
    training_type VARCHAR,
    status        VARCHAR,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    finished_at   TIMESTAMP NULL,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (card_set_id) REFERENCES card_sets (id) ON DELETE CASCADE
);

CREATE TABLE training_tasks_results
(
    id          BIGSERIAL PRIMARY KEY,
    training_id INT,
    card_id     INT,
    answer      VARCHAR NULL,
    is_correct  BOOLEAN NULL ,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (training_id) REFERENCES trainings (id) ON DELETE CASCADE,
    FOREIGN KEY (card_id) REFERENCES cards (id) ON DELETE CASCADE,
);

