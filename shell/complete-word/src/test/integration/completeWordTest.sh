#!/bin/bash
source "lib/assert.sh"
root="$(dirname $(dirname "$test"))"
export completeWordPath=$root
export chooseSh=$test/fake/choose.sh
complete=$root/src/main/complete.sh

assert "$complete --complete hello" "helloone"
assert "$complete -c hello" "helloone"
assert "$complete --complete Hello" "Helloone"
assert "$complete -c Hello" "Helloone"
assert_end canCompleteAWord

source $root/src/main/vars.sh

assert "$complete blahBlah" "$options"
assert_end canTellYouTheOptions