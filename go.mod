module github.com/btcsuite/btcutil

go 1.14

require (
	github.com/aead/siphash v1.0.1
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/davecgh/go-spew v1.1.1
	github.com/kkdai/bstream v1.0.0
	github.com/kr/text v0.2.0 // indirect
	github.com/onsi/gomega v1.12.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/btcsuite/btcd => github.com/btccom/btcd v0.0.0-20210520085239-5b8aab39eae7
