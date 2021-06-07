#!/bin/bash

source "lib/assert.sh"
keepCaps="src/main/words/keepCaps.sh"

assert "$keepCaps Com complete" "Complete"
assert "$keepCaps Com autocomplete" "Autocomplete"
assert "$keepCaps com complete" "complete"
assert "$keepCaps com autocomplete" "autocomplete"
assert_end keepTheCapitals