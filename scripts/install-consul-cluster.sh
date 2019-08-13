#!/bin/bash

# https://segmentfault.com/a/1190000008471221

consul agent -server -bootstrap-expect 3 -data-dir /tmp/consul -node=server001 -bind=10.0.0.10
consul agent -server  -data-dir /tmp/consul -node=server002 -bind=10.0.0.33 -join 10.0.0.10
consul agent -server  -data-dir /tmp/consul -node=server003 -bind=10.0.0.30 -join 10.0.0.10

