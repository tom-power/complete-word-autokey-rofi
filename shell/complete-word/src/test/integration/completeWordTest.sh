#!/bin/bash
source "lib/assert.sh"
root="$(dirname $(dirname "$test"))"
export completeWordPath=$root
common=$test/integration/common
export chooseSh=$common/chooseFirst.sh
complete=$root/src/main/complete.sh

assert "$complete --complete hello" "helloone"
assert "$complete -c hello" "helloone"
assert "$complete --complete Hello" "Helloone"
assert "$complete -c Hello" "Helloone"
assert_end canCompleteAWord

export chooseSh=$common/echoArg.sh
export addSh=$common/echoArg.sh
assert "$complete --add hello" "hello"
assert_end canAddAWord

source $root/src/main/vars.sh
assert "$complete blahBlah" "$options"
assert "$complete -b" "$options"
assert "$complete --blahBlah" "$options"
assert_end canTellYouTheOptions