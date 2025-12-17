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
INSERT INTO teams (name, wins, losses, draws, points_for, points_against, created_at, updated_at) VALUES
('Thunder Strikers', 5, 2, 1, 145, 98, '2024-01-01 12:00:00', '2024-12-01 18:00:00'),
('Lightning Bolts', 6, 1, 1, 168, 87, '2024-01-01 12:00:00', '2024-12-01 18:00:00'),
('Phoenix Rising', 4, 3, 1, 132, 115, '2024-01-01 12:00:00', '2024-12-01 18:00:00'),
('Dragon Warriors', 3, 4, 1, 110, 128, '2024-01-01 12:00:00', '2024-12-01 18:00:00'),
('Avalanche', 2, 5, 1, 95, 145, '2024-01-01 12:00:00', '2024-12-01 18:00:00'),
('Wildcats', 4, 3, 1, 128, 120, '2024-01-01 12:00:00', '2024-12-01 18:00:00');

---------------------------------------------------
-- 3. PLAYERS (League Participants)
---------------------------------------------------
-- Admin is not a player (just site admin)
-- Users 2-16 are players, some fully registered, some pending
INSERT INTO players (user_id, team_id, registration_fee_due, is_fully_registered, is_active, jersey_number, created_at, updated_at) VALUES
(2, 1, 0.00, TRUE, TRUE, 7, '2024-01-05 09:45:00', '2024-01-15 10:00:00'),
(3, 1, 0.00, TRUE, TRUE, 12, '2024-01-05 10:30:00', '2024-01-15 10:00:00'),
(4, 1, 50.00, FALSE, TRUE, NULL, '2024-01-06 11:30:00', '2024-01-06 11:30:00'),
(5, 2, 0.00, TRUE, TRUE, 23, '2024-01-06 15:00:00', '2024-01-15 10:00:00'),
(6, 2, 0.00, TRUE, TRUE, 5, '2024-01-07 08:15:00', '2024-01-15 10:00:00'),
(7, 2, 25.00, FALSE, TRUE, NULL, '2024-01-07 09:45:00', '2024-01-07 09:45:00'),
(8, 3, 0.00, TRUE, TRUE, 18, '2024-01-08 10:15:00', '2024-01-15 10:00:00'),
(9, 3, 0.00, TRUE, TRUE, 9, '2024-01-08 11:30:00', '2024-01-15 10:00:00'),
(10, 4, 0.00, TRUE, TRUE, 14, '2024-01-09 13:15:00', '2024-01-15 10:00:00'),
(11, 4, 75.00, FALSE, TRUE, NULL, '2024-01-09 14:45:00', '2024-01-09 14:45:00'),
(12, 5, 0.00, TRUE, TRUE, 3, '2024-01-10 09:15:00', '2024-01-15 10:00:00'),
(13, 5, 0.00, TRUE, TRUE, 21, '2024-01-10 11:00:00', '2024-01-15 10:00:00'),
(14, 6, 0.00, TRUE, TRUE, 10, '2024-01-11 08:45:00', '2024-01-15 10:00:00'),
(15, 6, 100.00, FALSE, TRUE, NULL, '2024-01-11 11:15:00', '2024-01-11 11:15:00'),
(16, NULL, 150.00, FALSE, TRUE, NULL, '2024-01-12 09:30:00', '2024-01-12 09:30:00');

---------------------------------------------------
-- 4. PAYMENTS (Transaction Records)
---------------------------------------------------
-- Payments for fully registered players
INSERT INTO payments (player_id, stripe_id, amount, status, payment_date) VALUES
-- Player 1 (user 2) - John Smith
(1, 'pi_3ABC123DEF456GHI789', 150.00, 'completed', '2024-01-10 14:30:00'),
-- Player 2 (user 3) - Sarah Jones
(2, 'pi_3ABC123DEF456GHI790', 150.00, 'completed', '2024-01-11 09:15:00'),
-- Player 3 (user 4) - Mike Wilson (partial payment)
(3, 'pi_3ABC123DEF456GHI791', 100.00, 'completed', '2024-01-12 11:00:00'),
-- Player 4 (user 5) - Emily Brown
(4, 'pi_3ABC123DEF456GHI792', 150.00, 'completed', '2024-01-13 10:45:00'),
-- Player 5 (user 6) - David Lee
(5, 'pi_3ABC123DEF456GHI793', 150.00, 'completed', '2024-01-14 15:20:00'),
-- Player 6 (user 7) - Jessica Davis (partial payment)
(6, 'pi_3ABC123DEF456GHI794', 125.00, 'completed', '2024-01-15 08:30:00'),
-- Player 7 (user 8) - Chris Martin
(7, 'pi_3ABC123DEF456GHI795', 150.00, 'completed', '2024-01-16 13:00:00'),
-- Player 8 (user 9) - Amanda Taylor
(8, 'pi_3ABC123DEF456GHI796', 150.00, 'completed', '2024-01-17 10:00:00'),
-- Player 9 (user 10) - James Anderson
(9, 'pi_3ABC123DEF456GHI797', 150.00, 'completed', '2024-01-18 11:30:00'),
-- Player 10 (user 11) - Lisa Thomas (partial payment)
(10, 'pi_3ABC123DEF456GHI798', 75.00, 'completed', '2024-01-19 09:45:00'),
-- Player 11 (user 12) - Robert Garcia
(11, 'pi_3ABC123DEF456GHI799', 150.00, 'completed', '2024-01-20 14:15:00'),
-- Player 12 (user 13) - Michelle Rodriguez
(12, 'pi_3ABC123DEF456GHI800', 150.00, 'completed', '2024-01-21 10:30:00'),
-- Player 13 (user 14) - Kevin Martinez
(13, 'pi_3ABC123DEF456GHI801', 150.00, 'completed', '2024-01-22 12:00:00'),
-- Player 14 (user 15) - Nicole Hernandez (partial payment)
(14, 'pi_3ABC123DEF456GHI802', 50.00, 'completed', '2024-01-23 08:45:00'),
-- Some pending/failed payments
(3, 'pi_3ABC123DEF456GHI803', 50.00, 'pending', '2024-12-01 10:00:00'),
(6, 'pi_3ABC123DEF456GHI804', 25.00, 'pending', '2024-12-02 14:30:00'),
(15, 'pi_3ABC123DEF456GHI805', 100.00, 'failed', '2024-11-28 16:20:00');

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

-- Update sequence values to continue from the last inserted ID
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));
SELECT setval('teams_id_seq', (SELECT MAX(id) FROM teams));
SELECT setval('players_id_seq', (SELECT MAX(id) FROM players));
SELECT setval('payments_id_seq', (SELECT MAX(id) FROM payments));
SELECT setval('games_id_seq', (SELECT MAX(id) FROM games));

-- Verify data insertion
SELECT 'Users: ' || COUNT(*) FROM users;
SELECT 'Teams: ' || COUNT(*) FROM teams;
SELECT 'Players: ' || COUNT(*) FROM players;
SELECT 'Payments: ' || COUNT(*) FROM payments;
SELECT 'Games: ' || COUNT(*) FROM games;
