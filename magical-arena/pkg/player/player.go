package player


// Player represents a player in the game. It has attributes for health, strength and attack.
type Player struct {
	health int
	strength int
	attack int
}


// NewPlayer creates and initializes a new Player instance with the given attributes.
// The function takes values for health, strength, and attack.
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
func NewPlayer(health, strength, attack int) *Player {
	return &Player{health, strength, attack}
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
func GetPlayerBaseAttributes(p *Player) (int, int, int) {
	return p.health, p.strength, p.attack
}

