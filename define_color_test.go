package main

import (
	"testing"
)

func TestDefineColorWithInterval0to5(t *testing.T) {
	got := defineColor(0, 5, 0)
	if got != "\033[35m" {
		t.Errorf(`Expected: \033[35m, but got: %s`, got)
	}
}

func TestDefineColorWithInterval1to6(t *testing.T) {
	got := defineColor(1, 6, 4)
	if got != "\033[32m" {
		t.Errorf(`Expected: \033[32m, but got: %s`, got)
	}
}

func TestDefineColorWithInterval5to15(t *testing.T) {
	got := defineColor(5, 15, 13)
	if got != "\033[33m" {
		t.Errorf(`Expected: \033[33m, but got: %s`, got)
	}
}

func TestDefineColorWithInterval0to4(t *testing.T) {
	got := defineColor(0, 4, 7)
	if got != "\033[31m" {
		t.Errorf(`Expected: \033[31m, but got: %s`, got)
	}
}

func TestDefineColorWithInterval_4to9(t *testing.T) {
	got := defineColor(-4, 9, -6)
	if got != "\033[35m" {
		t.Errorf(`Expected: \033[35m, but got: %s`, got)
	}
}
