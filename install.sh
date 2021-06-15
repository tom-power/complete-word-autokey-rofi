#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# autokey
cp -r ./autokey/* ~/.config/autokey/data/scripts/

# shell scripts
source $root/install.config.sh
cp -r ./shell/completeWord.sh $yourPath
source $root/.config/complete-word-autokey-rofi/install.config.sh
if [[ ! -d $completeWordPath ]]; then
  mkdir $completeWordPath
fi
cp -r ./shell/complete-word/* $completeWordPath

# config
configPath=~/.config/complete-word-autokey-rofi/
if [[ ! -d $configPath ]]; then
  mkdir $configPath
fi
cp -r $root/.config/complete-word-autokey-rofi/* $configPath

echo "finished"