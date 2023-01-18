#!/bin/sh
# copy here your individual public key certificates (proxy and external trusts)
ln -vf $BSD_KEY/ca/rootCA.pem $BSD_DEV/squidr/certstore/rootCA.pem
ln -vf $BSD_KEY/ca/external_trust.pem $BSD_DEV/squidr/certstore/external_trust.pem
