#!/bin/bash
rofi -dmenu \
-p "" \
-i -normalize-match \
-normal-window \
-no-fixed-num-lines \
-width 30 \
-matching fuzzy \
-sort -sorting-method levenshtein \
-filter $1