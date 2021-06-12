#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source $root/.local/install.config.sh

cp -r ./shell/.local* $yourPath
cp -r ./shell/completeWord.sh $yourPath

source $root/shell/.local/completeWord.config.sh

cp -r ./shell/complete-word* $completeWordPath

cp -r ./autokey/* ~/.config/autokey/data/scripts/

echo "------------install complete------------"