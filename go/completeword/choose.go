package completeword

import (
	"os"
	"os/exec"
	"strings"
)

const rofi =
    "rofi " +
    "-dmenu p '' i -normalize-match " +
    "normal-window no-fixed-num-lines " +
    "width 30 " +
    "matching fuzzy " +
    "sort -sorting-method levenshtein " +
    "-filter "

func choose(selection string, words string) (string, error) {
    var cmd *exec.Cmd
    cmd = exec.Command("sh", "-c", rofi + selection)
    cmd.Stderr = os.Stderr
    cmd.Stdin = strings.NewReader(words)
    out, err := cmd.Output()
    result := strings.TrimRight(string(out), "\n")
    return result, err
}

