
HERE=$(abspath .)

all: main1 main2

clean:
	rm -fr main1 main2 go/pkg

main1: main.c go/pkg/linux_amd64/p.h
	gcc -s -Wall -Igo/pkg/linux_amd64 -o main1 main.c go/pkg/linux_amd64/m.a -lpthread

main2: main.c go/pkg/linux_amd64_shared/p.h
	gcc -s -Wall -Igo/pkg/linux_amd64_shared -o main2 main.c go/pkg/linux_amd64_shared/m.a

go/pkg/linux_amd64/p.h: go/src/m/m.go go/src/p/p.go
	GOPATH=$(HERE)/go go install -buildmode=c-archive m
	rm go/pkg/linux_amd64/p.a

go/pkg/linux_amd64_shared/p.h: go/src/m/m.go go/src/p/p.go
	GOPATH=$(HERE)/go go install -buildmode=c-shared m
	rm go/pkg/linux_amd64_shared/p.a
