#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source $root/.local/completeWord.config.sh
export chooseSh=$completeWordPath/src/main/words/choose.sh
export addSh=$completeWordPath/src/main/words/add.sh
$completeWordPath/src/main/complete.sh "$@"