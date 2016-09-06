package routes

import (
	"testing"
)

func TestMixPotionItems(t *testing.T) {
	x := mixPotionItems([]Item{
		{1, "", -50.0},
		{2, "", 1.0},
		{3, "", -1.0},
	}...)
	if x < 0 {
		t.Error("toxicity value cannot be less than 0")
	}

	x = mixPotionItems([]Item{
		{1, "", 0},
		{2, "", 0},
		{3, "", 0},
	}...)
	if x != 0 {
		t.Error("toxicity value should be equal to 0")
	}
}
