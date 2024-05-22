# Installation target defaults to /usr/local/bin if not defined! 
# Please set DESTDIR or PREFIX to modify!

all: 
	sh certstore/update.sh
	sh .build.sh

mini:
	sh certstore/update.sh
	sh .build.sh

test:
	if [ -e APP/test/test.sh ];then sh APP/test/test.sh ; fi

install:

clean: 
	rm certstore/*.pem
	touch certstore/rootCA.pem certstore/external_trust.pem 
