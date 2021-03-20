#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/achievements
curl localhost:9090/achievements/a2181017-5c53-422b-b6bc-036b27c04fc8
curl localhost:9090/achievements -XPOST -d '{"name":"addName", "description":"addDescription", "condition":"addCondition"}'
curl localhost:9090/achievements -XPUT -d '{"id":0, "name":"newName", "description":"newDescription", "condition":"newCondition"}'
curl localhost:9090/achievements/a2181017-5c53-422b-b6bc-036b27c04fc8 -XDELETE
