package completeword

import (
	"os/exec"
	"strings"
)

type ChooseWord func([]string, string) (string, error)

var ChooseWordUsingRofi ChooseWord = func(words []string, selection string) (string, error) {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", rofi+selection)
	cmd.Stdin = strings.NewReader(strings.Join(words, "\n"))
	out, err := cmd.Output()
	return string(out), err
}

const rofi = "rofi " +
	"-dmenu p '' i " +
	"-normal-window -no-fixed-num-lines " +
	"-normalize-match -matching fuzzy " +
	"-sort -sorting-method levenshtein " +
	"-filter "
