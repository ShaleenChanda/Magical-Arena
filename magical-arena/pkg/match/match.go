package match


// Import the player package to use the Player struct.
import (
	"magical-arena/pkg/player"
)


// Match represents a match between two players in the Magical Arena.
type Match struct {
    // PlayerA is a pointer to the first player in the match.
    PlayerA *player.Player

    // PlayerB is a pointer to the second player in the match.
    PlayerB *player.Player

    // RoundResults stores the results of each round in the match.
    roundResults []string

    // Result indicates the overall result of the match (e.g., "PlayerA wins", "Draw", etc.).
    result string
}


// NewMatch creates and initializes a new Match instance with the provided players.
//
// Parameters:
//   - playerA: A pointer to the first player in the match.
//   - playerB: A pointer to the second player in the match.
//
// Returns:
//   - *Match: A pointer to the newly created Match instance.
//
// Example:
//   match := NewMatch(player1, player2)
func NewMatch(playerA, playerB *player.Player) *Match {
	return &Match{playerA, playerB, []string{}, ""}
}


func ConductMatch(match *Match) {
	
}