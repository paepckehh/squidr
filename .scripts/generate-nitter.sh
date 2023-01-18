#!/bin/sh
NAME="nitter"
T="../config-$NAME.go"
URL="https://github.com/zedeus/$NAME.wiki" # alt: xnaas.github.io/nitter-instancees
REPO="$NAME.wiki"
CHECKOUT="$REPO"
INFILE="$CHECKOUT/Instances.md"
echo "package squidr" > $T
echo >> $T
echo "// var _$NAME contains a generated list of all currently active $NAME instances" >> $T
echo "// *** DO NOT EDIT *** generated on $(date +%Y-%m-%d) via paepcke.de/squidr/.scripts/generate-$NAME.sh" >> $T
echo "// source: $URL" >> $T
echo "var _$NAME = []string{" >> $T
# git.clone $URL $REPO
# git.update $REPO
git.checkout $REPO
cat $INFILE | xurls | sort -u | sed '/git/d' | sed '/ssllabs/d' | sed '/instances/d' | while read line; do
	case $line in
	https://nitter*) DOMAIN="$(echo $line | sed 's/https:\/\///g' | sed 's/\///g')" && echo "\"$DOMAIN\"," >> $T ;;
	esac
done
echo >> $T
echo "}" >> $T
cat $T | uniq | sponge $T
rm -rf $CHECKOUT
gofumpt -d -w -extra $T
