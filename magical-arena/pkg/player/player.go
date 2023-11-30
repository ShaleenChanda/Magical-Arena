package player


// Player represents a player in the game. It has attributes for health, strength and attack.
type Player struct {
	name string

	health int

	strength int
	
	attack int
}


// NewPlayer creates and initializes a new Player instance with the specified attributes.
//
// Parameters:
//   - name: The name of the player.
//   - health: The health attribute of the player.
//   - strength: The strength attribute of the player.
//   - attack: The attack attribute of the player.
//
// Returns:
//   - *Player: A pointer to the newly created Player instance.
//
// Example:
//   player := NewPlayer("Name", 100, 10, 5)
//   fmt.Printf("%s has %d health, %d strength, and %d attack\n", player.Name, player.Health, player.Strength, player.Attack)
//
// Note: The example assumes a Player struct with exported fields (Name, Health, Strength, Attack).
func NewPlayer(name string, health, strength, attack int) *Player {
	return &Player{name, health, strength, attack}
}

// GetPlayerBaseAttributes returns the fundamental attributes of a player, including name, health, strength, and attack.
//
// Parameters:
//   - p: A pointer to the Player whose basic attributes are to be retrieved.
//
// Returns:
//   - string: Name of the player.
//   - int: Health of the player.
//   - int: Strength of the player.
//   - int: Attack value of the player.
//
// Example:
//   name, health, strength, attack := GetPlayerBaseAttributes(player)
//   fmt.Printf("%s has %d health, %d strength, and %d attack\n", name, health, strength, attack)
//
// Note: This function is designed for retrieving the core attributes of a player.
func GetPlayerBaseAttributes(p *Player) (string, int, int, int) {
	return p.name, p.health, p.strength, p.attack
}

