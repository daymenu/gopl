package main

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	Deposit(100)
	t.Errorf("want 100, but get %d", Balance())
}
