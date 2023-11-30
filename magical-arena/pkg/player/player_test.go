package player

import (
	"fmt"
	"os"
	"testing"
)


func TestNewPlayer(t *testing.T) {
	//testing NewPlayer and GetPlayerBaseAttributes as a single unit
	//TEST 1: create a player name shaleen with health 100, strength 10, attack 5, and check the assignment of values is correct
	player := NewPlayer("shaleen", 100, 10, 5)
	name, health, strength, attack := GetPlayerBaseAttributes(player)
	if name != "shaleen" {
		t.Errorf("Expected player.name to be shaleen, got %s", name)
	}
	if health != 100 {
		t.Errorf("Expected player.health to be 100, got %d", health)
	}
	if strength != 10 {
		t.Errorf("Expected player.strength to be 10, got %d", strength)
	}
	if attack != 5 {
		t.Errorf("Expected player.attack to be 5, got %d", attack)
	}

	//TEST 2: create a playerA with health 50, strength 5, attack 2
	//		create a playerB with health 100, strength 10, attack 5, and check the assignment of values is correct
	playerA := NewPlayer("PlayerA", 50, 5, 2)
	playerB := NewPlayer("PlayerB",100, 10, 5)
	nameA, healthA, strengthA, attackA := GetPlayerBaseAttributes(playerA)
	nameB, healthB, strengthB, attackB := GetPlayerBaseAttributes(playerB)
	//check playerA
	if nameA != "PlayerA" {
		t.Errorf("Expected playerA.name to be PlayerA, got %s", nameA)
	}
	if healthA != 50 {
		t.Errorf("Expected playerA.health to be 50, got %d", healthA)
	}
	if strengthA != 5 {
		t.Errorf("Expected playerA.strength to be 5, got %d", strengthA)
	}
	if attackA != 2 {
		t.Errorf("Expected playerA.attack to be 2, got %d", attackA)
	}
	//check playerB
	if nameB != "PlayerB" {
		t.Errorf("Expected playerB.name to be PlayerB, got %s", nameB)
	}
	if healthB != 100 {
		t.Errorf("Expected playerB.health to be 100, got %d", healthB)
	}
	if strengthB != 10 {
		t.Errorf("Expected playerB.strength to be 10, got %d", strengthB)
	}
	if attackB != 5 {
		t.Errorf("Expected playerB.attack to be 5, got %d", attackB)
	}

	//TEST 3: create a playerA with health 0, strength 0, attack 0
	//create a playerB with health 0, strength 0, attack 0, and check that they are equal. 
	playerA = NewPlayer("0", 0, 0, 0)
	playerB = NewPlayer("0",0, 0, 0)
	nameA, healthA, strengthA, attackA = GetPlayerBaseAttributes(playerA)
	nameB, healthB, strengthB, attackB = GetPlayerBaseAttributes(playerB)
	if(nameA != nameB || healthA != healthB || strengthA != strengthB || attackA != attackB) {
		t.Errorf("Expected playerA and playerB to be equal, got %d %d %d and %d %d %d", healthA, strengthA, attackA, healthB, strengthB, attackB)
	}
}

// TestMain runs the main testing suite.
func TestMain(m *testing.M) {
	fmt.Println("Testing player package...")
	Result:=m.Run()
	fmt.Println("Testing complete.")
	os.Exit(Result)
}
