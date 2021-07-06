#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# binary
if [[ -f $root/install.conf.sh ]]; then
  source $root/install.conf.sh
  cd $root/go && ./build.sh && cp build/completeWord $binPath
fi

# autokey
cp -r $root/autokey/* ~/.config/autokey/data/scripts/complete-word/

# words
configPath=~/.config/complete-word-autokey-rofi/
if [[ ! -d $configPath ]]; then
  mkdir $configPath
fi
cp -r $root/words $configPath

echo "finished"