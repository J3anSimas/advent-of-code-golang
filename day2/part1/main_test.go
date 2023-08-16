package main

import "testing"

func TestGetPlayScore(t *testing.T) {
	want := 4
	score := getPlayScore(RockEnemy, MyRock)
	if score != want {
		t.Errorf("getPlayScore(RockEnemy, MyRock) = %d, want %d", score, want)
	}
	want = 8
	score = getPlayScore(RockEnemy, MyPaper)
	if score != want {
		t.Errorf("getPlayScore(RockEnemy, MyPaper) = %d, want %d", score, want)
	}
	want = 3
	score = getPlayScore(RockEnemy, MyScissors)
	if score != want {
		t.Errorf("getPlayScore(RockEnemy, MyScissors) = %d, want %d", score, want)
	}
	want = 1
	score = getPlayScore(PaperEnemy, MyRock)
	if score != want {
		t.Errorf("getPlayScore(PaperEnemy, MyRock) = %d, want %d", score, want)
	}
	want = 5
	score = getPlayScore(PaperEnemy, MyPaper)
	if score != want {
		t.Errorf("getPlayScore(PaperEnemy, MyPaper) = %d, want %d", score, want)
	}
	want = 9
	score = getPlayScore(PaperEnemy, MyScissors)
	if score != want {
		t.Errorf("getPlayScore(PaperEnemy, MyScissors) = %d, want %d", score, want)
	}
	want = 7
	score = getPlayScore(ScissorsEnemy, MyRock)
	if score != want {
		t.Errorf("getPlayScore(ScissorsEnemy, MyRock) = %d, want %d", score, want)
	}
	want = 2
	score = getPlayScore(ScissorsEnemy, MyPaper)
	if score != want {
		t.Errorf("getPlayScore(ScissorsEnemy, MyPaper) = %d, want %d", score, want)
	}
	want = 6
	score = getPlayScore(ScissorsEnemy, MyScissors)
	if score != want {
		t.Errorf("getPlayScore(ScissorsEnemy, MyScissors) = %d, want %d", score, want)
	}

}
