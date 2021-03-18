#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/achievements
curl localhost:9090/achievements/1
curl localhost:9090/achievements -XPOST -d '{"name":"addName", "description":"addDescription", "condition":"addCondition"}'
curl localhost:9090/achievements -XPUT -d '{"id":2, "name":"newName", "description":"addDescription", "condition":"addCondition"}'
curl localhost:9090/achievements/1 -XDELETE
