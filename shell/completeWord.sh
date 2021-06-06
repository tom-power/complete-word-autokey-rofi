#!/bin/bash
export root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
$root/complete-word/complete.sh $1