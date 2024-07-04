gen:
	c-for-go kuzu.yml

clean:
	rm -f kuzu/cgo_helpers.go kuzu/cgo_helpers.h kuzu/cgo_helpers.c
	rm -f kuzu/const.go kuzu/doc.go kuzu/types.go
	rm -f kuzu/kuzu.go

test: gen
	cd kuzu && go build

