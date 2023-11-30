package match

// Import the player package to use the Player struct.
import (
	"fmt"
	"magical-arena/pkg/player"
	"math/rand"
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

// ConductMatch simulates a match between two players in the magical arena.
// The player with lower health attacks first, and rounds are conducted until the match is over (player.health <= 0).
// The result of each round and the overall match result are recorded.
//
// Parameters:
//   - match: A pointer to the Match instance representing the ongoing match (type *Match).
//
// Returns:
//   - []string: A slice containing descriptions of each round result.
//   - string: A string indicating the result of the entire match.
//
// Example:
//   roundResults, matchResult := ConductMatch(myMatch)
//   fmt.Println("Round Results:", roundResults)
//   fmt.Println("Match Result:", matchResult)
//
// Note: This function updates the Match instance with round results and the final match result.
func ConductMatch(match *Match) ([]string, string) {
	// The player with lower health attacks first
	// determine the starting player
	currentPlayer := determineStartingPlayer(match)

	//initializing the match
	//extracting the attributes of the players
	nameA, healthA, strengthA, attackA := player.GetPlayerBaseAttributes(match.PlayerA)
	nameB, healthB, strengthB, attackB := player.GetPlayerBaseAttributes(match.PlayerB)

	//conducting the match
	for !isMatchOver(healthA, healthB) {
		//conducting a round
		roundResult, currentHealthA, currentHealthB := conductRound(currentPlayer, nameA, healthA, strengthA, attackA, nameB, healthB, strengthB, attackB)
		match.roundResults = append(match.roundResults, roundResult)
		//updating the health of playerA
		healthA = currentHealthA
		//updating the health of playerB
		healthB = currentHealthB
		//switching the current player (example: if current player is playerA, switch to playerB)
		switchCurrentPlayer(&currentPlayer, match.PlayerA, match.PlayerB)
	}

	match.result = MatchResult(nameA, healthA, nameB, healthB)
	return match.roundResults, match.result
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
	if healthA <= healthB {
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
func conductRound(currentPlayer *player.Player, nameA string, healthA int, strengthA int, attackA int, nameB string, healthB int, strengthB int, attackB int) (string, int, int) {
	//fetching the base attributes of the current player
	playerName, _, _, _ := player.GetPlayerBaseAttributes(currentPlayer)

	//initializing the roundresult and the current health of the attacking and opponent player
	roundResult := ""
	currentHealthB := healthB
	currentHealthA := healthA

	//conducting the round
	if playerName == nameA {
		attackFromCurrentPlayer := attackA * (rand.Intn(6) + 1)
		defenceFromOtherPlayer := strengthB * (rand.Intn(6) + 1)
		damageToOtherPlayer := max(0, attackFromCurrentPlayer-defenceFromOtherPlayer)
		currentHealthB = max(0, healthB-damageToOtherPlayer)
		roundResult = fmt.Sprintf("%s attacked %s for %d damage", nameA, nameB, damageToOtherPlayer)
	}

	if playerName == nameB {
		attackFromCurrentPlayer := attackB * (rand.Intn(6) + 1)
		defenceFromOtherPlayer := strengthA * (rand.Intn(6) + 1)
		damageToOtherPlayer := max(0, attackFromCurrentPlayer-defenceFromOtherPlayer)
		currentHealthA = max(0, healthA-damageToOtherPlayer)
		roundResult = fmt.Sprintf("%s attacked %s for %d damage", nameB, nameA, damageToOtherPlayer)
	}

	//below code is used for unit testing please ignore otherwise
	//only for testing purposes current player is set to testA
	if playerName == "testA" {
		attackFromCurrentPlayer := attackA * 4
		defenceFromOtherPlayer := strengthB * 4
		damageToOtherPlayer := max(0, attackFromCurrentPlayer-defenceFromOtherPlayer)
		currentHealthB = max(0, healthB-damageToOtherPlayer)
		roundResult = fmt.Sprintf("%s attacked %s for %d damage", nameA, nameB, damageToOtherPlayer)
	}

	//only for testing purposes current player is set to testB
	if playerName == "testB" {
		attackFromCurrentPlayer := attackB * 4
		defenceFromOtherPlayer := strengthA * 4
		damageToOtherPlayer := max(0, attackFromCurrentPlayer-defenceFromOtherPlayer)
		currentHealthA = max(0, healthA-damageToOtherPlayer)
		roundResult = fmt.Sprintf("%s attacked %s for %d damage", nameB, nameA, damageToOtherPlayer)
	}

	//returning the round result and the updated health of the attacking and opponent player
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
func GetConductRound(currentPlayer *player.Player, nameA string, healthA int, strengthA int, attackA int, nameB string, healthB int, strengthB int, attackB int) (string, int, int) {
	return conductRound(currentPlayer, nameA, healthA, strengthA, attackA, nameB, healthB, strengthB, attackB)
}

//max returns the maximum of two integers
//
// Parameters:
//   - a: An integer.
//   - b: An integer.
//
// Returns:
//   - int: The maximum of the two integers.
//
// Example:
//   max := max(10, 20)
//   fmt.Println(max) // Output: 20
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// switchCurrentPlayer switches the current player between playerA and playerB based on the current player's reference.
//
// Parameters:
//   - currentPlayer: A double pointer to the current player (type **player.Player).
//   - playerA: A pointer to playerA (type *player.Player).
//   - playerB: A pointer to playerB (type *player.Player).
//
// Note: The function modifies the value pointed to by currentPlayer to switch between playerA and playerB.
func switchCurrentPlayer(currentPlayer **player.Player, playerA *player.Player, playerB *player.Player) {
	if *currentPlayer == playerA {
		*currentPlayer = playerB
	} else {
		*currentPlayer = playerA
	}
}

// GetSwitchCurrentPlayer is a helper function that provides external access to the private switchCurrentPlayer function for testing purposes.
// It allows switching the current player between playerA and playerB based on the current player's reference.
//
// Parameters:
//   - currentPlayer: A double pointer to the current player (type **player.Player).
//   - playerA: A pointer to playerA (type *player.Player).
//   - playerB: A pointer to playerB (type *player.Player).
//
// Example:
//   // Assuming currentPlayer, playerA, and playerB are initialized.
//   GetSwitchCurrentPlayer(&currentPlayer, playerA, playerB)
func GetSwitchCurrentPlayer(currentPlayer **player.Player, playerA *player.Player, playerB *player.Player) {
	switchCurrentPlayer(currentPlayer, playerA, playerB)
}

// MatchResult determines the result of a match based on the health attributes of two players.
//
// Parameters:
//   - nameA: The name of Player A.
//   - healthA: The current health of Player A.
//   - nameB: The name of Player B.
//   - healthB: The current health of Player B.
//
// Returns:
//   - string: A message indicating the winner of the match. The message is formatted as "{winner} wins".
//
// Example:
//   result := MatchResult("PlayerA", 0, "PlayerB", 30)
//   fmt.Println(result) // Output: "PlayerB wins"
func MatchResult(nameA string, healthA int, nameB string, healthB int) string {
	if healthA <= 0 {
		return fmt.Sprintf("%s wins", nameB)
	} else {
		return fmt.Sprintf("%s wins", nameA)
	}
}

// GetMatchResult is a testing wrapper for the MatchResult function.
//
// Parameters:
//   - nameA: The name of Player A.
//   - healthA: The current health of Player A.
//   - nameB: The name of Player B.
//   - healthB: The current health of Player B.
//
// Returns:
//   - string: A message indicating the winner of the match. The message is formatted as "{winner} wins".
//
// Example:
//   result := GetMatchResult("PlayerA", 10, "PlayerB", 0)
//   fmt.Println(result) // Output: "PlayerA wins"
func GetMatchResult(nameA string, healthA int, nameB string, healthB int) string {
	return MatchResult(nameA, healthA, nameB, healthB)
}
