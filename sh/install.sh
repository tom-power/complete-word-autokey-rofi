#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# autokey
cp -r $root/autokey/* ~/.config/autokey/data/scripts/complete-word/

# config
mkdir -p ~/.config/complete-word-autokey-rofi/words

echo "finished"
