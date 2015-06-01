#!/bin/sh
goxc -max-processors=$(grep -c ^processor /proc/cpuinfo) -include=README*,script/*,static/*,autoexec.json
