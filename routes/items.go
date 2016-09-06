package routes

type ItemMin struct {
	Id   int    `json:"id"`
	Name string `json:"string"`
}

type Item struct {
	Id       int     `json:"id"`
	Name     string  `json:"string"`
	Toxicity float32 `json:"toxicity"`
}

var allItems []Item
var mapItems map[int]Item

func resetItems() {
	allItems = []Item{
		{1, "Acidic Tonic", 10.5},
		{2, "Brew of Exalted Cats", 11.75},
		{3, "Brew of the Innocent Gibberer", 5.28},
		{4, "Disrupting Elixir", 6.7},
		{5, "Drakes' Elixir", 1.3},
		{6, "Draught of Dryness and Virtue", 20.3},
		{7, "Draught of the Witchery of Absorbs Heroism", 27.25},
		{8, "Drink of Conjure Law", 33.33},
		{9, "Elves' Elixir", 14.35},
		{10, "Immovable Tonic", 29.75},
		{11, "Just Gnome's Philter", 47.45},
		{12, "Just Returner's Tonic of Serpent Seduction", 94.46},
		{13, "Philter of the Travellers", 4.20},
		{14, "Potion of Studious Leaders", 17.15},
		{15, "Saintly Fish's Elixir of Drake Killing", 0.65},
		{16, "Scintillating Draught", 88.15},
		{17, "Studious Travellers' Draught", 71.45},
		{18, "Ultimate Canines' Potion", 51.0},
		{19, "Vicious Barbarian's Potion", 45.3},
		{20, "Deadly Purple Philter of Archmagi Calling", 3.5},
	}

	mapItems = make(map[int]Item)
	for _, item := range allItems {
		mapItems[item.Id] = item
	}
}
