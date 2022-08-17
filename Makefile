test:
	go test ./tools/benchmark
build:
	go build main.go
	mv ./main ./bin/wmt