#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source $root/.config/install.config.sh

cp -r ./shell/completeWord.sh $yourPath

source $root/.config/complete-word-autokey-rofi/completeWord.config.sh

if [[ ! -d $completeWordPath ]]; then
  mkdir $completeWordPath
fi
cp -r ./shell/complete-word/* $completeWordPath

cp -r ./autokey/* ~/.config/autokey/data/scripts/

configPath=~/.config/complete-word-autokey-rofi/
if [[ ! -d $configPath ]]; then
  mkdir $configPath
fi
cp -r $root/.config/complete-word-autokey-rofi/* $configPath

echo "------------install complete------------"