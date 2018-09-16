
.PHONY: validate
validate: build/kayako
	build/kayako doc

build/kayako:
	@mkdir -p build 
	@go build -o build/kayako github.com/sisatech/vorteil-documentation/kayako 

.PHONY: sync 
sync:
	build/kayako doc backup --sync
