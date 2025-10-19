package main

type Match struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	BeginAt    string     `json:"begin_at"`
	Status     string     `json:"status"`
	League     League     `json:"league"`
	Serie      Serie      `json:"serie"`
	Tournament Tournament `json:"tournament"`
	Opponents  []Opponent `json:"opponents"`
}

type League struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Serie struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Slug     string `json:"slug"`
	Year     int    `json:"year"`
}

type Tournament struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	BeginAt string `json:"begin_at"`
	EndAt   string `json:"end_at"`
}

type Opponent struct {
	Opponent OpponentTeam `json:"opponent"`
}

type OpponentTeam struct {
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
}
