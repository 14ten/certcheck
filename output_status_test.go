package main

import "testing"

func TestStatus(t *testing.T) {
	cases := []struct {
		name string
		r    Result
		warn int
		crit int
		want string
	}{
		{"ok", Result{DaysLeft: 90}, 30, 7, "OK"},
		{"warn lower bound", Result{DaysLeft: 30}, 30, 7, "WARN"},
		{"crit lower bound", Result{DaysLeft: 7}, 30, 7, "CRIT"},
		{"crit beats warn", Result{DaysLeft: 5}, 30, 7, "CRIT"},
		{"error wins", Result{Error: "boom", DaysLeft: 100}, 30, 7, "ERR"},
		{"already expired", Result{DaysLeft: -3}, 30, 7, "CRIT"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := status(c.r, c.warn, c.crit); got != c.want {
				t.Errorf("got %s want %s", got, c.want)
			}
		})
	}
}
