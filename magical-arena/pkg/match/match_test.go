package match


import (
	"testing"
	"magical-arena/pkg/player"
)


// TestGetDeterminStartingPlayer tests the GetDeterminStartingPlayer function,
// which determines the starting player for a match based on the health attributes
// of the players.
//
// Test scenarios:
//   1. Create a match with PlayerA's health 100 and PlayerB's health 50. Check that
//      PlayerB is the starting player.
//   2. Create a match with PlayerA's health 50 and PlayerB's health 100. Check that
//      PlayerA is the starting player.
//   3. Create a match with PlayerA's health 50 and PlayerB's health 100. Check that
//      PlayerA is the starting player.
//
// Note: The tests cover different scenarios to ensure that the starting player
// determination is based on health attributes as expected.
func TestGetDeterminStartingPlayer(t *testing.T){
	//TEST 1: create a match with playerA health 100, playerB health 50
	//		check that playerB is the starting player
	playerA := player.NewPlayer("PlayerA", 100, 10, 5)
	playerB := player.NewPlayer("PlayerB", 50, 5, 2)
	match := NewMatch(playerA, playerB)
	startingPlayer := GetDeterminStartingPlayer(match)
	startingPlayerName, _, _, _ := player.GetPlayerBaseAttributes(startingPlayer)
	if startingPlayerName != "PlayerB" {
		t.Errorf("Expected startingPlayerName to be PlayerB, got %s", startingPlayerName)
	}

	//TEST 2: create a match with playerA health 50, playerB health 100
	//		check that playerB is the starting player
	playerA = player.NewPlayer("PlayerA", 100, 5, 2)
	playerB = player.NewPlayer("PlayerB", 100, 10, 5)
	match = NewMatch(playerA, playerB)
	startingPlayer = GetDeterminStartingPlayer(match)
	startingPlayerName, _, _, _ = player.GetPlayerBaseAttributes(startingPlayer)
	if startingPlayerName != "PlayerA" {
		t.Errorf("Expected startingPlayerName to be PlayerA, got %s", startingPlayerName)
	}

	//TEST 3: create a match with playerA health 50, playerB health 100
	// 		check that playerA is the starting player
	playerA = player.NewPlayer("PlayerA", 50, 5, 2)
	playerB = player.NewPlayer("PlayerB", 100, 10, 5)
	match = NewMatch(playerA, playerB)
	startingPlayer = GetDeterminStartingPlayer(match)
	startingPlayerName, _, _, _ = player.GetPlayerBaseAttributes(startingPlayer)
	if startingPlayerName != "PlayerA" {
		t.Errorf("Expected startingPlayerName to be PlayerA, got %s", startingPlayerName)
	}
}


// TestGetIsMatchOver tests the GetIsMatchOver function, which determines whether
// a match is over based on the health attributes of the players.
//
// Test scenarios:
//   1. Create a match with PlayerA's health 100 and PlayerB's health 50. Check that
//      the match is not over.
//   2. Create a match with PlayerA's health 0 and PlayerB's health 50. Check that
//      the match is over.
//   3. Create a match with PlayerA's health 50 and PlayerB's health 0. Check that
//      the match is over.
//   4. Create a match with PlayerA's health 0 and PlayerB's health 0. Check that
//      the match is over.
func TestGetIsMatchOver(t *testing.T){
	//TEST 1: create a match with playerA health 100, playerB health 50
	// 		GetIsMatchOver should return false as both players are alive (i.e. health > 0)
	healthA := 100
	healthB := 50
	matchOver := GetIsMatchOver(healthA, healthB)
	if matchOver != false {
		t.Errorf("Expected matchOver to be false, got %t", matchOver)
	}

	//TEST 2: create a match with playerA health 0, playerB health 50
	// 		GetIsMatchOver should return true as playerA is dead (i.e. health <= 0)
	healthA = 0
	healthB = 50
	matchOver = GetIsMatchOver(healthA, healthB)
	if matchOver != true {
		t.Errorf("Expected matchOver to be true, got %t", matchOver)
	}

	//TEST 3: create a match with playerA health 50, playerB health 0
	// 		GetIsMatchOver should return true as playerB is dead (i.e. health <= 0)
	healthA = 50
	healthB = 0
	matchOver = GetIsMatchOver(healthA, healthB)
	if matchOver != true {
		t.Errorf("Expected matchOver to be true, got %t", matchOver)
	}

	//TEST 4: create a match with playerA health 0, playerB health 0
	// 		GetIsMatchOver should return true as both players are dead (i.e. health <= 0)
	healthA = 0
	healthB = 0
	matchOver = GetIsMatchOver(healthA, healthB)
	if matchOver != true {
		t.Errorf("Expected matchOver to be true, got %t", matchOver)
	}
}


func TestGetConductRound(t *testing.T){
	// TEST 1: creating a current player with name "testA". 
	// playerA is the current player, playerB is the opponent
	// playerA attributes: health 100, strength 10, attack 10
	// playerB attributes: health 50, strength 5, attack 2
	//expected health of PlayerB after round = 50 - max(0, 10*4 - 5*4) = 30
	playerA := player.NewPlayer("testA", 100, 10, 10)
	//playerB := player.NewPlayer("PlayerB", 50, 5, 2)
	currentPlayer := playerA
	roundResult, healthA, healthB := conductRound(currentPlayer, "testA", 100, 10, 10, "PlayerB", 50, 5, 2)
	if healthB != 30 {
		t.Errorf("Expected healthB to be 30, got %d", healthB)
	}
	if roundResult != "testA attacked PlayerB for 20 damage" {
		t.Errorf("Expected roundResult to be 'testA attacked PlayerB for 20 damage', got %s", roundResult)
	}
	if healthA != 100 {
		t.Errorf("Expected healthA to be 100, got %d", healthA)
	}

	// TEST 2: creating a current player with name "testA". 
	// playerA is the current player, playerB is the opponent
	// playerA attributes: health 100, strength 10, attack 4
	// playerB attributes: health 50, strength 5, attack 2
	//expected health of PlayerB after round = 50 - max(0, 4*4 - 5*4) = 50
	playerA = player.NewPlayer("testA", 100, 10, 4)
	//playerB := player.NewPlayer("PlayerB", 50, 5, 2)
	currentPlayer = playerA
	roundResult, healthA, healthB = conductRound(currentPlayer, "testA", 100, 10, 4, "PlayerB", 50, 5, 2)
	if healthB != 50 {
		t.Errorf("Expected healthB to be 50, got %d", healthB)
	}
	if roundResult != "testA attacked PlayerB for 0 damage" {
		t.Errorf("Expected roundResult to be 'testA attacked PlayerB for 0 damage', got %s", roundResult)
	}
	if healthA != 100 {
		t.Errorf("Expected healthA to be 100, got %d", healthA)
	}

	// TEST 3: creating a current player with name "testB".
	// playerB is the current player, playerA is the opponent
	// playerA attributes:  health 50, strength 5, attack 2
	// playerB attributes: health 100, strength 10, attack 10
	// expected health of PlayerA after round = 50 - max(0, 10*4 - 5*4) = 30
	playerB := player.NewPlayer("testB", 100, 10, 10)
	currentPlayer = playerB
	roundResult, healthA, healthB = conductRound(currentPlayer, "PlayerA", 50, 5, 2, "testB", 100, 10, 10)
	if healthA != 30 {
		t.Errorf("Expected healthA to be 30, got %d", healthA)
	}
	if roundResult != "testB attacked PlayerA for 20 damage" {
		t.Errorf("Expected roundResult to be 'testB attacked PlayerA for 20 damage', got %s", roundResult)
	}
	if healthB != 100 {
		t.Errorf("Expected healthB to be 100, got %d", healthB)
	}

	// TEST 4: creating a current player with name "testB".
	// playerB is the current player, playerA is the opponent
	// playerA attributes:  health 50, strength 5, attack 2
	// playerB attributes: health 100, strength 10, attack 4
	// expected health of PlayerA after round = 50 - max(0, 4*4 - 5*4) = 50
	playerB = player.NewPlayer("testB", 100, 10, 4)
	currentPlayer = playerB
	roundResult, healthA, healthB = conductRound(currentPlayer, "PlayerA", 50, 5, 2, "testB", 100, 10, 4)
	if healthA != 50 {
		t.Errorf("Expected healthA to be 50, got %d", healthA)
	}
	if roundResult != "testB attacked PlayerA for 0 damage" {
		t.Errorf("Expected roundResult to be 'testB attacked PlayerA for 0 damage', got %s", roundResult)
	}
	if healthB != 100 {
		t.Errorf("Expected healthB to be 100, got %d", healthB)
	}
}