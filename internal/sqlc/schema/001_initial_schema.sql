-- Set the current time zone to UTC for consistency
SET TIME ZONE 'UTC';

---------------------------------------------------
-- 1. USERS Table (Authentication/Site Accounts)
---------------------------------------------------
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    phone_number TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('normal', 'admin')),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

---------------------------------------------------
-- 2. TEAMS Table (League Competitors)
---------------------------------------------------
CREATE TABLE teams (
    id BIGSERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    wins INT NOT NULL DEFAULT 0,
    losses INT NOT NULL DEFAULT 0,
    draws INT NOT NULL DEFAULT 0,
    points_for INT NOT NULL DEFAULT 0,
    points_against INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

---------------------------------------------------
-- 3. PLAYERS Table (League Participants)
---------------------------------------------------
CREATE TABLE players (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE, -- 1:1 relationship
    team_id BIGINT REFERENCES teams(id) ON DELETE SET NULL,                -- Many:1 relationship (nullable)
    registration_fee_due DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    is_fully_registered BOOLEAN NOT NULL DEFAULT FALSE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    jersey_number INT, -- Nullable
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

---------------------------------------------------
-- 4. PAYMENTS Table (Transaction Records)
---------------------------------------------------
CREATE TABLE payments (
    id BIGSERIAL PRIMARY KEY,
    player_id BIGINT NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    stripe_id TEXT UNIQUE NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('completed', 'pending', 'failed')),
    payment_date TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

-- Index for faster lookup of a player's payments
CREATE INDEX ON payments (player_id);

---------------------------------------------------
-- 5. GAMES Table (Scheduled Matchups)
---------------------------------------------------
CREATE TABLE games (
    id BIGSERIAL PRIMARY KEY,
    home_team_id BIGINT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    away_team_id BIGINT NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    game_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT teams_cannot_be_same CHECK (home_team_id <> away_team_id)
);

---------------------------------------------------
-- 6. GAME_RESULTS Table (Scores and Outcome)
---------------------------------------------------
CREATE TABLE game_results (
    id BIGSERIAL PRIMARY KEY,
    game_id BIGINT UNIQUE NOT NULL REFERENCES games(id) ON DELETE CASCADE, -- 1:1 relationship with Game
    home_score INT NOT NULL DEFAULT 0,
    away_score INT NOT NULL DEFAULT 0,
    winning_team_id BIGINT REFERENCES teams(id) ON DELETE SET NULL, 
    recorded_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
