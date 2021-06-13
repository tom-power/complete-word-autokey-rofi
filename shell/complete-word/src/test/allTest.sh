#!/bin/bash
test="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

unit="${test}/unit"
echo "-------------------------unit-------------------------------------"
echo "-------------------------keepCaps----------------------------------"
source "$unit/keepCapsTest.sh"
unit="${test}/unit"
printf "\n"
echo "-------------------------integration------------------------------"
echo "-------------------------completeWord--------------------------------"
source "${test}/integration/completeWordTest.sh"
