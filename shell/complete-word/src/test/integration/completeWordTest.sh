#!/bin/bash
source "lib/assert.sh"

root="$(dirname $(dirname "$test"))"

export completeWordPath=$root
export chooseSh=$test/fake/choose.sh
complete=$root/src/main/complete.sh

assert "$complete hello" "helloone"
assert_end canCompleteAWord