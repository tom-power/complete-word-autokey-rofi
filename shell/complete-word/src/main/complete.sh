#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
$root/words/keepCaps.sh $1 $($root/words/all.sh | $root/words/choose.sh $1)