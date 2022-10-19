CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY,
    owner_id UUID NOT NULL REFERENCES users(id),
    opponent_id UUID REFERENCES users(id),
    current_player_id UUID REFERENCES users(id),
    step_count INTEGER NOT NULL DEFAULT 0,
    winner_id UUID REFERENCES users(id),
    field json DEFAULT '[["", "", ""], ["", "", ""], ["", "", ""]]',
    current_state smallint NOT NULL DEFAULT 0
);
