#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source $root/vars.sh

for arg in "$@"; do
  shift
  case "$arg" in
    "--complete")
        set -- "$@" "-c"
        ;;
    "--add")
        set -- "$@" "-a"
        ;;
    *)
        set -- "$@" "$arg"
        ;;
  esac
done

while getopts "c:a:" arg; do
    case "${arg}" in
        c)
            selection=${OPTARG}
            completion=$($root/words/all.sh | $chooseSh $selection)
            $root/words/keepCaps.sh $selection $completion
            ;;
        a)
            selection=${OPTARG}
            $addSh $($chooseSh $selection)
            ;;
        *)
            echo $options
            exit 0;
            ;;
    esac
done
shift $((OPTIND-1))
if (( $OPTIND == 1 )); then
   echo $options
fi
