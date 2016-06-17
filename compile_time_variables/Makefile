
IS=${if ${shell sh -c 'case `go version` in go1.[01234].*) ;; *) echo yes; ;; esac'},=, }

BUILD_TIME=${shell date +%FT%T%z}

main: main.go
	go build -ldflags "-X main.BuildTime${IS}${BUILD_TIME}" main.go

# voor een pakket, i.p.v. main.BuildTime${IS}$${BUILD_TIME} gebruik:
#     github.com/pebbe/pakket.BuildTime${IS}${BUILD_TIME}
