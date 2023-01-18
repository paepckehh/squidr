#!/bin/sh
ls -I *.sh | while read line; do
	case $line in
	generate.sh) ;;
	generate-*) sh $line ;;
	esac
done
