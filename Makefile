.PHONY: test bench

test:
	go test -v -count=1 ./...

bench:
	{ echo "Benchmark Iterations Time ns/op Memory B/op Allocs allocs/op"; go test -bench=. -benchmem -count=1 ./... | grep "^Benchmark"; } | column -t
