package main

/*
#cgo CFLAGS: -I${SRCDIR}/../native/packaged/include -I${SRCDIR}/../native/packaged/include/TrustWalletCore
#cgo LDFLAGS: -L${SRCDIR}/../native/packaged/lib -lTrustWalletCore -lwallet_core_rs -lprotobuf -lTrezorCrypto -lstdc++ -lm
*/
import "C"

import (
	"fmt"

	"github.com/Cramiumlabs/wallet-core/wrapper/go-wrapper/chain_abstraction"
	"github.com/Cramiumlabs/wallet-core/wrapper/go-wrapper/chain_abstraction/chains"
)

// Global registry instance
var globalRegistry *chain_abstraction.ChainRegistry

// init function to register all chains when package is imported
func init() {
	globalRegistry = chain_abstraction.NewChainRegistry()
	chains.RegisterAll(globalRegistry)
}

func main() {
	// All chains are now registered and ready to use!
	availableChains := globalRegistry.ListChains()
	fmt.Printf("Registered %d chains\n", len(availableChains))
}
