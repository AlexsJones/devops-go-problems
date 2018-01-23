package lock_test

import (
	"testing"

	"github.com/SeedJobs/devops-go-problems/lock/problem"
)

func TestDefaultLock(t *testing.T) {
	l := problem.NewLock(problem.NewEventLog())
	l.Initialise()
	if l.GetDisplay() != "UNLOCKED" {
		t.Fatal("Incorrect default display")
	}
}

func TestBasicUsage(t *testing.T) {
	l := problem.NewLock(problem.NewEventLog())
	l.Initialise()
	l.EnteredKey('1')
	if l.GetDisplay() != "1" {
		t.Log("Expected display is 1")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Doesn't display the correct key")
	}
	l.EnteredKey('1')
	if l.GetDisplay() != "11" {
		t.Log("Expected display is 11")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Doesn't display the correct key")
	}
	l.EnteredKey('h')
	if l.GetDisplay() != "11" {
		t.Log("Expected display is 11")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Doesn't display the correct key")
	}
	l.EnteredKey('s')
	if l.GetDisplay() != "UNLOCKED" {
		t.Fatal("Incorrect default display")
	}
}

func TestLocking(t *testing.T) {
	l := problem.NewLock(problem.NewEventLog())
	l.Initialise()
	for i := '0'; i < '5'; i++ {
		l.EnteredKey(i)
	}
	if l.GetDisplay() != "0123" {
		t.Log("Expected display is 0123")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Doesn't display the correct key")
	}
	l.EnteredKey('h')
	if l.GetDisplay() != "LOCKED" {
		t.Fatal("Incorrect default display")
	}
}

func TestTiming(t *testing.T) {
	l := problem.NewLock(problem.NewEventLog())
	l.Initialise()
	l.EnteredKey('1')
	for i := 0; i < 9; i++ {
		l.Tick()
	}
	if l.GetDisplay() != "1" {
		t.Log("Expected the display to be 1")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Incorrect default display")
	}
	l.Tick()
	if l.GetDisplay() != "UNLOCKED" {
		t.Log("Expected display to be UNLOCKED")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Incorrect default display")
	}
	for i := '0'; i < '4'; i++ {
		l.EnteredKey(i)
	}
	if err := l.EnteredKey('h'); err != nil {
		t.Fatal(err)
	}
	if l.GetDisplay() != "LOCKED" {
		t.Log("Expected display to be LOCKED")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Incorrect default display")
	}
	if err := l.EnteredKey('1'); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 9; i++ {
		l.Tick()
	}
	if l.GetDisplay() != "1" {
		t.Log("Expected the display to be 1")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Incorrect default display")
	}
	l.Tick()
	if l.GetDisplay() != "LOCKED" {
		t.Log("Expected display to be LOCKED")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Incorrect default display")
	}
}

func TestBadInput(t *testing.T) {
	l := problem.NewLock(problem.NewEventLog())
	l.Initialise()
	if err := l.EnteredKey('z'); err == nil {
		t.Fatal("An invalid key has been pressed")
	}
	for i := '0'; i < '4'; i++ {
		l.EnteredKey(i)
	}
	if err := l.EnteredKey('h'); err != nil {
		t.Fatal(err)
	}
	if l.GetDisplay() != "LOCKED" {
		t.Log("Expected display to be LOCKED")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Incorrect default display")
	}
	l.Initialise()
	if l.GetDisplay() != "UNLOCKED" {
		t.Log("Expected display to be UNLOCKED")
		t.Log("Actual display is", l.GetDisplay())
		t.Fatal("Incorrect default display")
	}
}
