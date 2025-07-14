.PHONY: build build-gojq

build-gojq: GOJQVERSION=$(shell cd gojq-lib; go list -m -f '{{ .Version }}' github.com/itchyny/gojq)
build-gojq:
	cd gojq-lib; GOOS=js GOARCH=wasm go build -ldflags="-X main.gojqVersion=$(GOJQVERSION)" -o ../front/script/gojs.wasm ./main.go
	cd gojq-lib; GOOS=js GOARCH=wasm go-licenses save . --save_path ../tmp/licenses --ignore gojqlib --force
	cd tmp/licenses; rm -f "../../front/LICENSES" ; echo -n "" > ../LICENSE ; find . -type f -exec sh -c 'echo $$0 | sed "s/^\.\///g" >> ../LICENSE ; cat $$0 >> ../LICENSE ; echo "" >> ../LICENSE' {} \; ;\
	cd ../../additional-licenses;                                             find . -type f -exec sh -c 'echo $$0 | sed "s/^\.\///g" >> ../tmp/LICENSE ; cat $$0 >> ../tmp/LICENSE ; echo "" >> ../tmp/LICENSE' {} \; ; mv ../tmp/LICENSE ../front/LICENSES.txt

docker:
	docker build -t ghcr.io/entrivax/gojq-web .