#!/usr/bin/env bash

src=~/coding/li/Algorithm-Problems/leetcode/

subdirs=$(ls $src)
for i in ${subdirs[@]}; do
	if [[ $i =~ L[0-9]*\..* ]]; then
		name=${i%.*}
		dst=${src}${name}
		mkdir -p $dst
		mv ${src}${i} $dst
	fi
done
