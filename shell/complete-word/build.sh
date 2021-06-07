#!/bin/bash
export root="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# add libs
cd ${root}/lib &&
wget https://raw.github.com/lehmannro/assert.sh/master/assert.sh -O assert.sh &&
chmod + assert.sh &&

# run tests
cd ${root} && src/test/allTest.sh