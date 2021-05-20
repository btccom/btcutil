module github.com/btcsuite/btcutil

go 1.14

require (
	github.com/aead/siphash v1.0.1
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/davecgh/go-spew v1.1.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/kkdai/bstream v0.0.0-20161212061736-f391b8402d23
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

replace github.com/btcsuite/btcd => github.com/btccom/btcd v0.0.0-20210520085239-5b8aab39eae7
