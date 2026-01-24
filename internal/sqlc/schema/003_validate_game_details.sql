-- Migration: Validate player's team participated in game
-- This ensures that a player can only have game_details entries for games 
-- where their team (home or away) actually played.

-- Create a function to check if a player's team played in a specific game
CREATE OR REPLACE FUNCTION validate_player_team_in_game(p_player_id BIGINT, p_game_id BIGINT)
RETURNS BOOLEAN AS $$
DECLARE
    player_team_id BIGINT;
    game_home_team_id BIGINT;
    game_away_team_id BIGINT;
BEGIN
    SELECT team_id INTO player_team_id
    FROM players
    WHERE id = p_player_id;
    
    -- If player has no team, return false
    IF player_team_id IS NULL THEN
        RETURN FALSE;
    END IF;
    
    SELECT home_team_id, away_team_id INTO game_home_team_id, game_away_team_id
    FROM games
    WHERE id = p_game_id;
    
    -- Check if player's team matches either home or away team
    RETURN (player_team_id = game_home_team_id OR player_team_id = game_away_team_id);
END;
$$ LANGUAGE plpgsql STABLE;

ALTER TABLE game_details 
ADD CONSTRAINT player_team_must_play_in_game 
CHECK (validate_player_team_in_game(player_id, game_id));
