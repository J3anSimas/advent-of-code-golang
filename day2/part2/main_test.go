package main

import "testing"

func TestGetPlayScore(t *testing.T) {
	want := 3
	score := getPlayScore(RockEnemy, Lose)
	if score != want {
		t.Errorf("getPlayScore(RockEnemy, Lose) = %d, want %d", score, want)
	}
	want = 4
	score = getPlayScore(RockEnemy, Draw)
	if score != want {
		t.Errorf("getPlayScore(RockEnemy, Draw) = %d, want %d", score, want)
	}
	want = 8
	score = getPlayScore(RockEnemy, Win)
	if score != want {
		t.Errorf("getPlayScore(RockEnemy, Win) = %d, want %d", score, want)
	}
	want = 1
	score = getPlayScore(PaperEnemy, Lose)
	if score != want {
		t.Errorf("getPlayScore(PaperEnemy, Lose) = %d, want %d", score, want)
	}
	want = 5
	score = getPlayScore(PaperEnemy, Draw)
	if score != want {
		t.Errorf("getPlayScore(PaperEnemy, Draw) = %d, want %d", score, want)
	}
	want = 9
	score = getPlayScore(PaperEnemy, Win)
	if score != want {
		t.Errorf("getPlayScore(PaperEnemy, Win) = %d, want %d", score, want)
	}
	want = 2
	score = getPlayScore(ScissorsEnemy, Lose)
	if score != want {
		t.Errorf("getPlayScore(ScissorsEnemy, Lose) = %d, want %d", score, want)
	}
	want = 6
	score = getPlayScore(ScissorsEnemy, Draw)
	if score != want {
		t.Errorf("getPlayScore(ScissorsEnemy, Draw) = %d, want %d", score, want)
	}
	want = 7
	score = getPlayScore(ScissorsEnemy, Win)
	if score != want {
		t.Errorf("getPlayScore(ScissorsEnemy, Win) = %d, want %d", score, want)
	}

}
