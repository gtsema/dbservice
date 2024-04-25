#!/bin/sh

if [ -f "./dbservice.db" ]; then
  ./dbservice
else
  ./dbservice &

  while [ ! -f "./dbservice.db" ]
  do
    sleep 1
  done

  ./goose up
fi
