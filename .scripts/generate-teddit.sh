#!/bin/sh
NAME="teddit"
T="../config-$NAME.go"
URL="https://codeberg.org/$NAME/$NAME"
REPO="$NAME"
CHECKOUT="$NAME"
INFILE="$CHECKOUT/instances.json"
echo "package squidr" > $T
echo >> $T
echo "// var _$NAME contains a generated list of all currently active $NAME instances" >> $T
echo "// *** DO NOT EDIT *** generated on $(date +%Y-%m-%d) via paepcke.de/squidr/.scripts/generate-$NAME.sh" >> $T
echo "// source: $URL" >> $T
echo "var _$NAME = []string{" >> $T
# git.clone $URL $REPO
# git.update $REPO
git.checkout $REPO
cat $INFILE | sed 's/,/\n/g' | xurls | sort -u | while read line; do
	case $line in
	https://teddit*) DOMAIN="$(echo $line | sed 's/https:\/\///g' | sed 's/\///g')" && echo "\"$DOMAIN\"," >> $T ;;
	esac
done
echo >> $T
echo "}" >> $T
cat $T | uniq | sponge $T
rm -rf $CHECKOUT
gofumpt -d -w -extra $T
