
PKG = PKG_CONFIG_PATH=/my/opt/zeromq-4.1.3/lib/pkgconfig
CGO = CGO_LDFLAGS="-L/my/opt/libsodium-1.0.3/lib -lsodium -lstdc++ -lrt"

all: version ver

version: version.go
	$(PKG) $(CGO) go build -v --ldflags '-extldflags "-static"' -a version.go
	strip version

# geen gcc met optie -s in combinatie met -static, anders kijg je bij run: unexpected reloc type in static binary
ver: ver.c
	gcc -Wall -static `$(PKG) pkg-config --cflags libzmq` -o ver ver.c \
		`$(PKG) pkg-config --libs libzmq` -L/my/opt/libsodium-1.0.3/lib -lsodium -lstdc++ -lrt -pthread
	strip ver
