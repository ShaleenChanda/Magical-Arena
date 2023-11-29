package main


// Player represents a player in the game. It has attributes for health, strength, attack, wins, and losses.
type Player struct {
	health int
	strength int
	attack int
	wins int
	losses int
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
	return &Player{health, strength, attack, 0, 0}
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


// GetPlayerMatchStats returns the match-related statistics of a player, including the number of wins and losses.
//
// Parameters:
//   - p: A pointer to the Player whose match-related stats are to be retrieved.
//
// Returns:
//   - int: Number of wins for the player.
//   - int: Number of losses for the player.
//
// Example:
//   wins, losses := GetPlayerMatchStats(player)
//	 fmt.Printf("Player has %d wins and %d losses\n", wins, losses)
//
// Note: This function is designed for retrieving the match-related statistics of a player.
func GetPlayerMatchStats(p *Player) (int, int) {
	return p.wins, p.losses
}


// IncrementWins increments the wins attribute of a player and returns the updated wins value.
//
// Parameters:
//   - p: A pointer to the Player whose wins attribute is to be incremented.
//
// Returns:
//   - int: The updated number of wins for the player.
func IncrementWins(p *Player) int{
	p.wins++;
	return p.wins
}


// IncrementLosses increments the losses attribute of a player and returns the updated losses value.
//
// Parameters:
//   - p: A pointer to the Player whose losses attribute is to be incremented.
//
// Returns:
//   - int: The updated number of losses for the player.
func IncrementLosses(p *Player) int{
	p.losses++;
	return p.losses
}
