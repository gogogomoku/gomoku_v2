package ai

import (
	"fmt"
	"testing"

	"github.com/gogogomoku/gomoku_v2/arcade/match"
)

func TestAi_Something(t *testing.T) {
	match := match.CreateMatch(false, false, 0)
	fmt.Println("this is a test")
	fmt.Println(match)
}
