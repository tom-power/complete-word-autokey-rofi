#!/bin/bash
rofi -dmenu \
-p "" \
-i -normalize-match \
-normal-window \
-matching fuzzy \
-sort -sorting-method levenshtein \
-filter $1