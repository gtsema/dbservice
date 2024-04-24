#!/bin/sh

./dbservice &

while [ ! -f "./dbservice.db" ]
do
  sleep 1
done

./goose up
