#!/bin/bash
firstOfFirst=${1:0:1}
firstOfSecond=${2:0:1}
restOfSecond=${2:1}
if [[ "$firstOfFirst" =~ [A-Z] ]]; then
  if [[ "${firstOfFirst,,}" = "${firstOfSecond,,}" ]]; then
    echo $firstOfFirst$restOfSecond
  else
    echo ${firstOfSecond^^}$restOfSecond
  fi
else
  echo $2
fi

