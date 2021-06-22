#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# autokey
cp -r $root/autokey/* ~/.config/autokey/data/scripts/complete-word/

# config
configPath=~/.config/complete-word-autokey-rofi/
if [[ ! -d $configPath ]]; then
  mkdir $configPath
fi
cp -r $root/words $configPath

echo "finished"