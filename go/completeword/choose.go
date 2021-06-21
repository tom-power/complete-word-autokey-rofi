package completeword

import (
	"os"
	"os/exec"
	"strings"
)

type ChooseWord func([]string, string) (string, error)

const rofi = "rofi " +
	"-dmenu p '' i -normalize-match " +
	"normal-window no-fixed-num-lines " +
	"width 30 " +
	"matching fuzzy " +
	"sort -sorting-method levenshtein " +
	"-filter "

var ChooseWordUsingRofi = func(words []string, selection string) (string, error) {
	var cmd *exec.Cmd
	cmd = exec.Command("sh", "-c", rofi+selection)
	cmd.Stderr = os.Stderr
	cmd.Stdin = strings.NewReader(strings.Join(words, "\n"))
	out, err := cmd.Output()
	result := strings.TrimRight(string(out), "\n")
	return result, err
}
