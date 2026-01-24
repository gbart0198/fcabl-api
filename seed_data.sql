-- FCABL Seed Data
-- Load this file with: psql -U <username> -d <database> -f seed_data.sql
-- Or via Docker: docker exec -i <container_name> psql -U <username> -d <database> < seed_data.sql

SET TIME ZONE 'UTC';

---------------------------------------------------
-- 1. USERS (Authentication/Site Accounts)
---------------------------------------------------
-- Password hash is bcrypt hash of "password123" for all test users
INSERT INTO users (email, phone_number, password_hash, first_name, last_name, role, created_at, updated_at) VALUES
('admin@fcabl.com', '555-0100', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Admin', 'User', 'admin', '2024-01-01 10:00:00', '2024-01-01 10:00:00'),
('john.smith@email.com', '555-0101', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'John', 'Smith', 'normal', '2024-01-05 09:30:00', '2024-01-05 09:30:00'),
('sarah.jones@email.com', '555-0102', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Sarah', 'Jones', 'normal', '2024-01-05 10:15:00', '2024-01-05 10:15:00'),
('mike.wilson@email.com', '555-0103', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Mike', 'Wilson', 'normal', '2024-01-06 11:20:00', '2024-01-06 11:20:00'),
('emily.brown@email.com', '555-0104', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Emily', 'Brown', 'normal', '2024-01-06 14:45:00', '2024-01-06 14:45:00'),
('david.lee@email.com', '555-0105', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'David', 'Lee', 'normal', '2024-01-07 08:00:00', '2024-01-07 08:00:00'),
('jessica.davis@email.com', '555-0106', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Jessica', 'Davis', 'normal', '2024-01-07 09:30:00', '2024-01-07 09:30:00'),
('chris.martin@email.com', '555-0107', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Chris', 'Martin', 'normal', '2024-01-08 10:00:00', '2024-01-08 10:00:00'),
('amanda.taylor@email.com', '555-0108', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Amanda', 'Taylor', 'normal', '2024-01-08 11:15:00', '2024-01-08 11:15:00'),
('james.anderson@email.com', '555-0109', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'James', 'Anderson', 'normal', '2024-01-09 13:00:00', '2024-01-09 13:00:00'),
('lisa.thomas@email.com', '555-0110', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Lisa', 'Thomas', 'normal', '2024-01-09 14:30:00', '2024-01-09 14:30:00'),
('robert.garcia@email.com', '555-0111', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Robert', 'Garcia', 'normal', '2024-01-10 09:00:00', '2024-01-10 09:00:00'),
('michelle.rodriguez@email.com', '555-0112', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Michelle', 'Rodriguez', 'normal', '2024-01-10 10:45:00', '2024-01-10 10:45:00'),
('kevin.martinez@email.com', '555-0113', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Kevin', 'Martinez', 'normal', '2024-01-11 08:30:00', '2024-01-11 08:30:00'),
('nicole.hernandez@email.com', '555-0114', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Nicole', 'Hernandez', 'normal', '2024-01-11 11:00:00', '2024-01-11 11:00:00'),
('daniel.moore@email.com', '555-0115', '$2a$10$rQ4hZKJ3YnZZHlDMKJ6j2eXqZ4vN7nwQHLdOZ3LxJxVfKJ3Y5Z6YO', 'Daniel', 'Moore', 'normal', '2024-01-12 09:15:00', '2024-01-12 09:15:00');

---------------------------------------------------
-- 2. TEAMS (League Competitors)
---------------------------------------------------
INSERT INTO teams (name, created_at, updated_at) VALUES
('Thunder Strikers', '2024-01-01 12:00:00', '2024-12-01 18:00:00'),
('Lightning Bolts', '2024-01-01 12:00:00', '2024-12-01 18:00:00'),
('Phoenix Rising', '2024-01-01 12:00:00', '2024-12-01 18:00:00'),
('Dragon Warriors', '2024-01-01 12:00:00', '2024-12-01 18:00:00'),
('Avalanche', '2024-01-01 12:00:00', '2024-12-01 18:00:00'),
('Wildcats', '2024-01-01 12:00:00', '2024-12-01 18:00:00');

---------------------------------------------------
-- 3. PLAYERS (League Participants)
---------------------------------------------------
INSERT INTO players (user_id, team_id, fee_remainder, jersey_number, created_at, updated_at) VALUES
(2, 1, 2500, 7, '2024-01-05 09:45:00', '2024-01-15 10:00:00'),
(3, 1, 2500, 12, '2024-01-05 10:30:00', '2024-01-15 10:00:00'),
(4, 1, 2500, NULL, '2024-01-06 11:30:00', '2024-01-06 11:30:00'),
(5, 2, 2500, 23, '2024-01-06 15:00:00', '2024-01-15 10:00:00'),
(6, 2, 2500, 5, '2024-01-07 08:15:00', '2024-01-15 10:00:00'),
(7, 2, 2500, NULL, '2024-01-07 09:45:00', '2024-01-07 09:45:00'),
(8, 3, 2500, 18, '2024-01-08 10:15:00', '2024-01-15 10:00:00'),
(9, 3, 2500, 9, '2024-01-08 11:30:00', '2024-01-15 10:00:00'),
(10, 4, 2500, 14, '2024-01-09 13:15:00', '2024-01-15 10:00:00'),
(11, 4, 2500, NULL, '2024-01-09 14:45:00', '2024-01-09 14:45:00'),
(12, 5, 2500, 3, '2024-01-10 09:15:00', '2024-01-15 10:00:00'),
(13, 5, 2500, 21, '2024-01-10 11:00:00', '2024-01-15 10:00:00'),
(14, 6, 2500, 10, '2024-01-11 08:45:00', '2024-01-15 10:00:00'),
(15, 6, 2500, NULL, '2024-01-11 11:15:00', '2024-01-11 11:15:00'),
(16, NULL, 2500, NULL, '2024-01-12 09:30:00', '2024-01-12 09:30:00');

---------------------------------------------------
-- 4. PAYMENTS (Transaction Records)
---------------------------------------------------
-- Payments for fully registered players
INSERT INTO payments (player_id, transaction_id, amount, status, payment_date) VALUES
(1, 'pi_3ABC123DEF456GHI789', 15000, 'completed', '2024-01-10 14:30:00'),
(2, 'pi_3ABC123DEF456GHI790', 15000, 'completed', '2024-01-11 09:15:00'),
(3, 'pi_3ABC123DEF456GHI791', 10000, 'completed', '2024-01-12 11:00:00'),
(4, 'pi_3ABC123DEF456GHI792', 15000, 'completed', '2024-01-13 10:45:00'),
(5, 'pi_3ABC123DEF456GHI793', 15000, 'completed', '2024-01-14 15:20:00'),
(6, 'pi_3ABC123DEF456GHI794', 12500, 'completed', '2024-01-15 08:30:00'),
(7, 'pi_3ABC123DEF456GHI795', 15000, 'completed', '2024-01-16 13:00:00'),
(8, 'pi_3ABC123DEF456GHI796', 15000, 'completed', '2024-01-17 10:00:00'),
(9, 'pi_3ABC123DEF456GHI797', 15000, 'completed', '2024-01-18 11:30:00'),
(10, 'pi_3ABC123DEF456GHI798', 7500, 'completed', '2024-01-19 09:45:00'),
(11, 'pi_3ABC123DEF456GHI799', 15000, 'completed', '2024-01-20 14:15:00'),
(12, 'pi_3ABC123DEF456GHI800', 15000, 'completed', '2024-01-21 10:30:00'),
(13, 'pi_3ABC123DEF456GHI801', 15000, 'completed', '2024-01-22 12:00:00'),
(14, 'pi_3ABC123DEF456GHI802', 5000, 'completed', '2024-01-23 08:45:00'),
(3, 'pi_3ABC123DEF456GHI803', 5000, 'pending', '2024-12-01 10:00:00'),
(6, 'pi_3ABC123DEF456GHI804', 2500, 'pending', '2024-12-02 14:30:00'),
(15, 'pi_3ABC123DEF456GHI805', 10000, 'failed', '2024-11-28 16:20:00');

---------------------------------------------------
-- 5. GAMES (Scheduled Matchups with Scores)
---------------------------------------------------
-- Completed games (with scores and status)
INSERT INTO games (home_team_id, away_team_id, home_score, away_score, game_time, status, created_at, updated_at) VALUES
(1, 2, 18, 21, '2024-02-10 14:00:00', 'completed', '2024-01-15 10:00:00', '2024-02-10 16:00:00'),
(3, 4, 14, 12, '2024-02-10 16:00:00', 'completed', '2024-01-15 10:00:00', '2024-02-10 18:00:00'),
(5, 6, 10, 15, '2024-02-10 18:00:00', 'completed', '2024-01-15 10:00:00', '2024-02-10 20:00:00'),
(2, 3, 24, 18, '2024-02-17 14:00:00', 'completed', '2024-01-15 10:00:00', '2024-02-17 16:00:00'),
(4, 5, 14, 14, '2024-02-17 16:00:00', 'completed', '2024-01-15 10:00:00', '2024-02-17 18:00:00'),
(6, 1, 16, 19, '2024-02-17 18:00:00', 'completed', '2024-01-15 10:00:00', '2024-02-17 20:00:00'),
(1, 3, 21, 17, '2024-02-24 14:00:00', 'completed', '2024-01-15 10:00:00', '2024-02-24 16:00:00'),
(2, 4, 22, 15, '2024-02-24 16:00:00', 'completed', '2024-01-15 10:00:00', '2024-02-24 18:00:00'),
(5, 1, 12, 20, '2024-03-02 14:00:00', 'completed', '2024-01-15 10:00:00', '2024-03-02 16:00:00'),
(3, 6, 19, 19, '2024-03-02 16:00:00', 'completed', '2024-01-15 10:00:00', '2024-03-02 18:00:00'),
(4, 1, 13, 18, '2024-03-09 14:00:00', 'completed', '2024-01-15 10:00:00', '2024-03-09 16:00:00'),
(2, 5, 25, 18, '2024-03-09 16:00:00', 'completed', '2024-01-15 10:00:00', '2024-03-09 18:00:00'),
-- Upcoming games (scheduled, no scores yet)
(1, 4, 0, 0, '2024-12-20 18:00:00', 'scheduled', '2024-01-15 10:00:00', '2024-01-15 10:00:00'),
(3, 5, 0, 0, '2024-12-20 20:00:00', 'scheduled', '2024-01-15 10:00:00', '2024-01-15 10:00:00'),
(2, 6, 0, 0, '2024-12-27 18:00:00', 'scheduled', '2024-01-15 10:00:00', '2024-01-15 10:00:00'),
(1, 5, 0, 0, '2025-01-03 18:00:00', 'scheduled', '2024-01-15 10:00:00', '2024-01-15 10:00:00');

---------------------------------------------------
-- 6. GAME_DETAILS (Player Scores Per Game)
---------------------------------------------------
-- Game 1: Thunder Strikers (1) 18 vs Lightning Bolts (2) 21
INSERT INTO game_details (game_id, player_id, score) VALUES
(1, 1, 6),  -- John Smith (Thunder Strikers)
(1, 2, 7),  -- Sarah Jones (Thunder Strikers)
(1, 3, 5),  -- Mike Wilson (Thunder Strikers)
(1, 4, 8),  -- Emily Brown (Lightning Bolts)
(1, 5, 7),  -- David Lee (Lightning Bolts)
(1, 6, 6);  -- Jessica Davis (Lightning Bolts)

-- Game 2: Phoenix Rising (3) 14 vs Dragon Warriors (4) 12
INSERT INTO game_details (game_id, player_id, score) VALUES
(2, 7, 8),  -- Chris Martin (Phoenix Rising)
(2, 8, 6),  -- Amanda Taylor (Phoenix Rising)
(2, 9, 6),  -- James Anderson (Dragon Warriors)
(2, 10, 6); -- Lisa Thomas (Dragon Warriors)

-- Game 3: Avalanche (5) 10 vs Wildcats (6) 15
INSERT INTO game_details (game_id, player_id, score) VALUES
(3, 11, 5),  -- Robert Garcia (Avalanche)
(3, 12, 5),  -- Michelle Rodriguez (Avalanche)
(3, 13, 8),  -- Kevin Martinez (Wildcats)
(3, 14, 7);  -- Nicole Hernandez (Wildcats)

-- Game 4: Lightning Bolts (2) 24 vs Phoenix Rising (3) 18
INSERT INTO game_details (game_id, player_id, score) VALUES
(4, 4, 9),   -- Emily Brown (Lightning Bolts)
(4, 5, 8),   -- David Lee (Lightning Bolts)
(4, 6, 7),   -- Jessica Davis (Lightning Bolts)
(4, 7, 10),  -- Chris Martin (Phoenix Rising)
(4, 8, 8);   -- Amanda Taylor (Phoenix Rising)

-- Game 5: Dragon Warriors (4) 14 vs Avalanche (5) 14
INSERT INTO game_details (game_id, player_id, score) VALUES
(5, 9, 7),   -- James Anderson (Dragon Warriors)
(5, 10, 7),  -- Lisa Thomas (Dragon Warriors)
(5, 11, 7),  -- Robert Garcia (Avalanche)
(5, 12, 7);  -- Michelle Rodriguez (Avalanche)

-- Game 6: Wildcats (6) 16 vs Thunder Strikers (1) 19
INSERT INTO game_details (game_id, player_id, score) VALUES
(6, 13, 8),  -- Kevin Martinez (Wildcats)
(6, 14, 8),  -- Nicole Hernandez (Wildcats)
(6, 1, 7),   -- John Smith (Thunder Strikers)
(6, 2, 6),   -- Sarah Jones (Thunder Strikers)
(6, 3, 6);   -- Mike Wilson (Thunder Strikers)

-- Game 7: Thunder Strikers (1) 21 vs Phoenix Rising (3) 17
INSERT INTO game_details (game_id, player_id, score) VALUES
(7, 1, 8),   -- John Smith (Thunder Strikers)
(7, 2, 7),   -- Sarah Jones (Thunder Strikers)
(7, 3, 6),   -- Mike Wilson (Thunder Strikers)
(7, 7, 9),   -- Chris Martin (Phoenix Rising)
(7, 8, 8);   -- Amanda Taylor (Phoenix Rising)

-- Game 8: Lightning Bolts (2) 22 vs Dragon Warriors (4) 15
INSERT INTO game_details (game_id, player_id, score) VALUES
(8, 4, 8),   -- Emily Brown (Lightning Bolts)
(8, 5, 7),   -- David Lee (Lightning Bolts)
(8, 6, 7),   -- Jessica Davis (Lightning Bolts)
(8, 9, 8),   -- James Anderson (Dragon Warriors)
(8, 10, 7);  -- Lisa Thomas (Dragon Warriors)

-- Game 9: Avalanche (5) 12 vs Thunder Strikers (1) 20
INSERT INTO game_details (game_id, player_id, score) VALUES
(9, 11, 6),  -- Robert Garcia (Avalanche)
(9, 12, 6),  -- Michelle Rodriguez (Avalanche)
(9, 1, 7),   -- John Smith (Thunder Strikers)
(9, 2, 7),   -- Sarah Jones (Thunder Strikers)
(9, 3, 6);   -- Mike Wilson (Thunder Strikers)

-- Game 10: Phoenix Rising (3) 19 vs Wildcats (6) 19
INSERT INTO game_details (game_id, player_id, score) VALUES
(10, 7, 10),  -- Chris Martin (Phoenix Rising)
(10, 8, 9),   -- Amanda Taylor (Phoenix Rising)
(10, 13, 10), -- Kevin Martinez (Wildcats)
(10, 14, 9);  -- Nicole Hernandez (Wildcats)

-- Game 11: Dragon Warriors (4) 13 vs Thunder Strikers (1) 18
INSERT INTO game_details (game_id, player_id, score) VALUES
(11, 9, 7),  -- James Anderson (Dragon Warriors)
(11, 10, 6), -- Lisa Thomas (Dragon Warriors)
(11, 1, 6),  -- John Smith (Thunder Strikers)
(11, 2, 6),  -- Sarah Jones (Thunder Strikers)
(11, 3, 6);  -- Mike Wilson (Thunder Strikers)

-- Game 12: Lightning Bolts (2) 25 vs Avalanche (5) 18
INSERT INTO game_details (game_id, player_id, score) VALUES
(12, 4, 9),  -- Emily Brown (Lightning Bolts)
(12, 5, 8),  -- David Lee (Lightning Bolts)
(12, 6, 8),  -- Jessica Davis (Lightning Bolts)
(12, 11, 9), -- Robert Garcia (Avalanche)
(12, 12, 9); -- Michelle Rodriguez (Avalanche)

-- Update sequence values to continue from the last inserted ID
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));
SELECT setval('teams_id_seq', (SELECT MAX(id) FROM teams));
SELECT setval('players_id_seq', (SELECT MAX(id) FROM players));
SELECT setval('payments_id_seq', (SELECT MAX(id) FROM payments));
SELECT setval('games_id_seq', (SELECT MAX(id) FROM games));
SELECT setval('game_details_id_seq', (SELECT MAX(id) FROM game_details));

-- Verify data insertion
SELECT 'Users: ' || COUNT(*) FROM users;
SELECT 'Teams: ' || COUNT(*) FROM teams;
SELECT 'Players: ' || COUNT(*) FROM players;
SELECT 'Payments: ' || COUNT(*) FROM payments;
SELECT 'Games: ' || COUNT(*) FROM games;
SELECT 'Game Details: ' || COUNT(*) FROM game_details;
