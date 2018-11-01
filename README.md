# tokenwrap

Demonstrates an issue where response wrapping a root token does not work

> $ export VAULT_TOKEN=45eeedf2-0e99-4ca9-0fc6-90def95f1de0
> $ export VAULT_ADDR=http://$(hostname -f):8200
> $ go install
> $ $GOPATH/bin/tokenwrap 
> DeriveRootToken op=POST path=auth/token/create wrapTTL=1m
> Created token secret: &{RequestID: LeaseID: LeaseDuration:0 Renewable:false Data:map[] Warnings:[] Auth:<nil> WrapInfo:0xc00005b1a0}
> Wrapped token: 
> Created token secret: &{RequestID:a8f862a8-4049-da85-4c81-76a3fcdb7ba7 LeaseID: LeaseDuration:0 Renewable:false Data:map[] Warnings:[] Auth:0xc0000ecb00 WrapInfo:<nil>}
> Non wrapped token: 1b7700bd-1b15-4576-2d4a-fabb0e3cf4c8
