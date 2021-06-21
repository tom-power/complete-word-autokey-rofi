#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# autokey
cp -r ./autokey/* ~/.config/autokey/data/scripts/

# binary
source $root/install.config.sh
cp -r ./go/build/completeWord $yourPath

# config
configPath=~/.config/complete-word-autokey-rofi/
if [[ ! -d $configPath ]]; then
  mkdir $configPath
fi
cp -r $root/.config/complete-word-autokey-rofi/* $configPath

echo "finished"