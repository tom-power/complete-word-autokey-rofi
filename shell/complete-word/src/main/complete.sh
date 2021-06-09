#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
selection=$1
completion=$($root/words/all.sh | $root/words/choose.sh $selection)
$root/words/keepCaps.sh $selection $completion