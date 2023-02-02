#!/bin/sh
NAME="searx"
T="../config-$NAME-map.go"
URL="https://searx.space/data/instances.json"
# REPO="$NAME"
# CHECKOUT="$REPO"
INFILE="$DOC/snip/searx.zstd"
echo "package squidr" > $T
echo >> $T
echo "// var _$NAME contains a generated list of all currently active $NAME instances" >> $T
echo "// *** DO NOT EDIT *** generated on $(date +%Y-%m-%d) via paepcke.de/squidr/.scripts/generate-$NAME.sh" >> $T
echo "// source: $URL" >> $T
echo "var _map$NAME = map[string]bool{" >> $T
# git.clone $URL $REPO
# git.update $REPO
# curl -o $INFILE $URL
zstdcat $INFILE | sed 's/,/\n/g' | xurls | sort -u | while read line; do
	case $line in
	https://sear*) DOMAIN="$(echo $line | sed 's/https:\/\///g' | cut -d '/' -f 1)" && echo "\"$DOMAIN\": true," >> $T ;;
	esac
done
echo >> $T
echo "}" >> $T
cat $T | uniq | sponge $T
rm -rf $CHECKOUT
gofumpt -d -w -extra $T
