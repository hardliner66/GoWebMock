#!/bin/sh
goxc -max-processors=$(grep -c ^processor /proc/cpuinfo) -include=README*,script/*,static/*,autoexec.json,keys/gen_test_key.sh,keys/server.crt,keys/server.key
