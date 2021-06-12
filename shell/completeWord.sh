#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source $root/.local/completeWord.config.sh
$completeWordPath/src/main/complete.sh $1