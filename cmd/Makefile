# Makefile

# Define variables
GO_CFLAGS="-I$(shell pwd)/../native/packaged/include"
CGO_LDFLAGS="-L$(shell pwd)/../native/packaged/lib -lTrustWalletCore -lwallet_core_rs -lprotobuf -lTrezorCrypto -lstdc++ -lm"

# Target to run the Go application
run:
	GO_CFLAGS=$(GO_CFLAGS) CGO_LDFLAGS=$(CGO_LDFLAGS) go run main.go
