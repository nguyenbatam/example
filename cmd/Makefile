# Makefile (Linux only)

# Paths (relative to this Makefile)
ROOT       := $(shell pwd)/..
INCLUDEDIR := $(ROOT)/native/packaged/include
LIBDIR     := $(ROOT)/native/packaged/lib

# CGO flags for Linux
CGO_CFLAGS  := -I$(INCLUDEDIR)
CGO_LDFLAGS := -L$(LIBDIR) -lTrustWalletCore -lwallet_core_rs -lprotobuf -lTrezorCrypto -lstdc++ -lm

# Run the Go app
run:
	CGO_CFLAGS="$(CGO_CFLAGS)" CGO_LDFLAGS="$(CGO_LDFLAGS)" go run main.go

# Build the Go app
build:
	CGO_CFLAGS="$(CGO_CFLAGS)" CGO_LDFLAGS="$(CGO_LDFLAGS)" go build -o app main.go

