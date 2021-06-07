#!/bin/bash
source "lib/assert.sh"

test="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

unit="${test}/unit"
echo "-------------------------unit-------------------------------------"
echo "-------------------------keepCaps----------------------------------"
source "$unit/keepCapsTest.sh"
