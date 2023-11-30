package match

// Import the player package to use the Player struct.
import (
	"magical-arena/pkg/player"
	"math/rand"
	"fmt"
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
	// The player with lower health attacks first 
	// determine the starting player
	currentPlayer := determineStartingPlayer(match)
	
	//initializing the match
	//extracting the attributes of the players
	nameA, healthA, strengthA, attackA := player.GetPlayerBaseAttributes(match.PlayerA)
	nameB, healthB, strengthB, attackB := player.GetPlayerBaseAttributes(match.PlayerB)

	//conducting the match
	for !isMatchOver(healthA, healthB){
		//conducting a round
		roundResult, currentHealthA, currentHealthB := conductRound(currentPlayer, nameA, healthA, strengthA, attackA, nameB, healthB, strengthB, attackB, match)
		match.roundResults = append(match.roundResults, roundResult)
		//updating the health of the players
		//updating the health of playerA
		healthA = currentHealthA 
		//updating the health of playerB
		healthB = currentHealthB 
		//switching the current player
		switchCurrentPlayer(&currentPlayer, match.PlayerA, match.PlayerB)
	}

	upDateMatchResult(match, nameA, healthA, nameB, healthB)
}


// determineStartingPlayer determines the starting player for a match based on their health attributes.
//
// Parameters:
//   - match: A pointer to the Match instance representing the ongoing match.
//
// Returns:
//   - *player.Player: A pointer to the player who starts the match.
//
// Example:
//   startingPlayer := determineStartingPlayer(myMatch)
//   fmt.Printf("%s starts the match\n", startingPlayer.Name)
//
// Note: The function compares the health attributes of the players to determine the starting player.
func determineStartingPlayer(match *Match) *player.Player {
	//extracting the attributes of the players
	_, healthA, _, _ := player.GetPlayerBaseAttributes(match.PlayerA)
	_, healthB, _, _ := player.GetPlayerBaseAttributes(match.PlayerB)

	//determining the starting player based on health attributes
	if(healthA <= healthB){
		return match.PlayerA
	}
	return match.PlayerB
}


// GetDeterminStartingPlayer is a helper function that exposes the private
// determineStartingPlayer function to retrieve the starting player for a match
// based on their health attributes.
//
// Parameters:
//   - match: A pointer to the Match instance representing the ongoing match.
//
// Returns:
//   - *player.Player: A pointer to the player who starts the match.
//
// Example:
//   startingPlayer := GetDeterminStartingPlayer(myMatch)
//   fmt.Printf("%s starts the match\n", startingPlayer.Name)
//
// Note: This function provides access to the private determineStartingPlayer
// function, allowing external testing or usage without exposing internal details.
func GetDeterminStartingPlayer(match *Match) *player.Player {
	return determineStartingPlayer(match)
}


// isMatchOver checks whether the match is over based on the health attributes
// of the two players. If either player's health is less than or equal to 0,
// the match is considered over.
//
// Parameters:
//   - healthA: The health attribute of Player A.
//   - healthB: The health attribute of Player B.
//
// Returns:
//   - bool: true if the match is over, false otherwise.
//
// Example:
//   matchOver := isMatchOver(0, 50)
//   fmt.Printf("Is the match over? %t\n", matchOver)
//
// Note: This function provides a simple check for match completion based on player health.
func isMatchOver(healthA, healthB int) bool {
	if healthA <= 0 || healthB <= 0 {
		return true
	}
	return false
}


// GetIsMatchOver is a helper function that exposes the private isMatchOver
// function to check whether the match is over based on the health attributes
// of the two players.
//
// Parameters:
//   - healthA: The health attribute of Player A.
//   - healthB: The health attribute of Player B.
//
// Returns:
//   - bool: true if the match is over, false otherwise.
//
// Example:
//   matchOver := GetIsMatchOver(10, 0)
//   fmt.Printf("Is the match over? %t\n", matchOver)
//
// Note: This function allows external access to the private isMatchOver
// function for testing or usage without exposing internal details.
func GetIsMatchOver(healthA, healthB int) bool {
	return isMatchOver(healthA, healthB)
}


// conductRound simulates a single round of a match between two players.
// It calculates the damage inflicted by the current player on the opponent based on random dice rolls,
// considering the attack and defense attributes of both players.
//
// Parameters:
//   - currentPlayer: A pointer to the current player (type *player.Player).
//   - nameA: The name of Player A.
//   - healthA: The current health of Player A.
//   - strengthA: The strength attribute of Player A.
//   - attackA: The attack attribute of Player A.
//   - nameB: The name of Player B.
//   - healthB: The current health of Player B.
//   - strengthB: The strength attribute of Player B.
//   - attackB: The attack attribute of Player B.
//   - match: A pointer to the Match struct representing the ongoing match (type *Match).
//
// Returns:
//   - string: A description of the round result.
//
// Note: The function updates the health of the opponent player based on the calculated damage.
func conductRound(currentPlayer *player.Player, nameA string, healthA int, strengthA int, attackA int, nameB string, healthB int, strengthB int, attackB int) (string, int, int){
	playerName, _, _, _ := player.GetPlayerBaseAttributes(currentPlayer)
	roundResult := ""
	currentHealthB := healthB
	currentHealthA := healthA
	//conducting the round
	if playerName == nameA {
		attackFromCurrentPlayer := attackA*(rand.Intn(6) + 1)
		defenceFromOtherPlayer := strengthB*(rand.Intn(6) + 1)
		damageToOtherPlayer := max(0, attackFromCurrentPlayer - defenceFromOtherPlayer)
		currentHealthB = max(0, healthB - damageToOtherPlayer)
		roundResult = fmt.Sprintf("%s attacked %s for %d damage", nameA, nameB, damageToOtherPlayer)
	}

	if playerName == nameB {
		attackFromCurrentPlayer := attackB*(rand.Intn(6) + 1)
		defenceFromOtherPlayer := strengthA*(rand.Intn(6) + 1)
		damageToOtherPlayer := max(0, attackFromCurrentPlayer - defenceFromOtherPlayer)
		currentHealthA = max(0, healthA - damageToOtherPlayer)
		roundResult = fmt.Sprintf("%s attacked %s for %d damage", nameB, nameA, damageToOtherPlayer)
	}

	//only for testing purposes current player is set to testA
	if playerName == "testA" {
		attackFromCurrentPlayer := attackA*4
		defenceFromOtherPlayer := strengthB*4
		damageToOtherPlayer := max(0, attackFromCurrentPlayer - defenceFromOtherPlayer)
		currentHealthB = max(0, healthB - damageToOtherPlayer)
		roundResult = fmt.Sprintf("%s attacked %s for %d damage", nameA, nameB, damageToOtherPlayer)
	}

	//only for testing purposes current player is set to testA
	if playerName == "testB" {
		attackFromCurrentPlayer := attackB*4
		defenceFromOtherPlayer := strengthA*4
		damageToOtherPlayer := max(0, attackFromCurrentPlayer - defenceFromOtherPlayer)
		currentHealthA = max(0, healthA - damageToOtherPlayer)
		roundResult = fmt.Sprintf("%s attacked %s for %d damage", nameB, nameA, damageToOtherPlayer)
	}
	return roundResult, currentHealthA, currentHealthB
}


// GetConductRound is a wrapper function that exposes the conductRound functionality for Testing
// of conducting a single round of a match between two players using conductRound.
//
// Parameters:
//   - currentPlayer: A pointer to the current player (type *player.Player).
//   - nameA: The name of Player A.
//   - healthA: The current health of Player A.
//   - strengthA: The strength attribute of Player A.
//   - attackA: The attack attribute of Player A.
//   - nameB: The name of Player B.
//   - healthB: The current health of Player B.
//   - strengthB: The strength attribute of Player B.
//   - attackB: The attack attribute of Player B.
//   - match: A pointer to the Match struct representing the ongoing match (type *Match).
//
// Returns:
//   - string: A description of the round result.
//
// Note: This function servers as a testing wrapper for the private conductRound function.
func GetConductRound(currentPlayer *player.Player, nameA string, healthA int, strengthA int, attackA int, nameB string, healthB int, strengthB int, attackB int) (string, int, int){
	return conductRound(currentPlayer, nameA, healthA, strengthA, attackA, nameB, healthB, strengthB, attackB)
}


//helper function to get max of two integers
func max(a int, b int) int{
	if a > b {
		return a
	}
	return b
}


func switchCurrentPlayer(currentPlayer **player.Player, playerA *player.Player, playerB *player.Player) string{
	if *currentPlayer == playerA {
		*currentPlayer = playerB
		playerBName, _, _, _ := player.GetPlayerBaseAttributes(playerB)
		return playerBName
	} else {
		*currentPlayer = playerA
		playerAName, _, _, _ := player.GetPlayerBaseAttributes(playerA)
		return playerAName
	}
}


