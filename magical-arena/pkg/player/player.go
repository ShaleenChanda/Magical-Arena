package player

import "fmt"

// Player represents a player in the game. It has attributes for health, strength, attack, wins, losses and draws.
type Player struct {
	health int
	strength int
	attack int
	wins int
	losses int
	draws int
}


// NewPlayer creates and initializes a new Player instance with the given attributes.
// The function takes values for health, strength, and attack, and initializes wins and losses to 0.
//
// Parameters:
//   - health: The health attribute of the player.
//   - strength: The strength attribute of the player.
//   - attack: The attack attribute of the player.
//
// Returns:
//   - *Player: A pointer to the newly created Player instance.
//
// Example:
//   player := NewPlayer(100, 10, 5)
//   fmt.Printf("Player has %d health, %d strength, and %d attack\n", player.health, player.strength, player.attack)
//
// Note: The wins and losses attributes are initialized to 0 by default.
func NewPlayer(health, strength, attack int) *Player {
	return &Player{health, strength, attack, 0, 0, 0}
}


// GetPlayerAttributes returns the fundamental attributes of a player, including health, strength, and attack.
//
// Parameters:
//   - p: A pointer to the Player whose basic attributes are to be retrieved.
//
// Returns:
//   - int: Health of the player.
//   - int: Strength of the player.
//   - int: Attack value of the player.
//
// Example:
//   health, strength, attack := GetPlayerAttributes(player)
//   fmt.Printf("Player has %d health, %d strength, and %d attack\n", health, strength, attack)
//
// Note: This function is designed for retrieving the core attributes of a player.
func GetPlayerAttributes(p *Player) (int, int, int) {
	return p.health, p.strength, p.attack
}


// GetPlayerMatchStats returns the match-related statistics of a player, including the number of wins, losses, draws, and total matches played.
//
// Parameters:
//   - p: A pointer to the Player whose match-related stats are to be retrieved.
//
// Returns:
//   - int: Number of wins for the player.
//   - int: Number of losses for the player.
//   - int: Number of draws for the player.
//   - int: Total number of matches played by the player.
//
// Example:
//   wins, losses, draws, totalMatchesPlayed := GetPlayerMatchStats(player)
//   fmt.Printf("Player has %d wins, %d losses, %d draws, and has played %d matches\n", wins, losses, draws, totalMatchesPlayed)
//
// Note: This function is designed for retrieving the match-related statistics of a player.
func GetPlayerMatchStats(p *Player) (int, int, int, int) {
    totalMatchesPlayed := p.wins + p.losses + p.draws
    return p.wins, p.losses, p.draws, totalMatchesPlayed
}


// IncrementStat increments the specified attribute (e.g., wins, losses, draws) of a player and returns the updated value.
//
// Parameters:
//   - p: A pointer to the Player whose attribute is to be incremented.
//   - stat: The attribute (wins, losses, draws) to be incremented.
//
// Returns:
//   - int: The updated value of the specified attribute for the player.
//   - error: An error value if the attribute name is invalid.
//
// Example:
//   wins, err := IncrementStat(player, "wins")
//   losses, err := IncrementStat(player, "losses")
//   draws, err := IncrementStat(player, "draws")
func IncrementStat(p *Player, stat string) (int, error) {
    switch stat {
    case "wins":
        p.wins++
        return p.wins, nil
    case "losses":
        p.losses++
        return p.losses, nil
    case "draws":
        p.draws++
        return p.draws, nil
    default:
        return 0, fmt.Errorf("Invalid stat: %s", stat)
    }
}