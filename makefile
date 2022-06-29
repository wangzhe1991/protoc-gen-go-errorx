export GOBIN  := $(CURDIR)/bin

.PHONY: init
# init dependcy
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

.PHONY: build
# build protoc-gen-go-xerrors
build:
	go build -o ./bin/protoc-gen-go-errorx
	cp -rf ./bin/protoc-gen-go-errorx /Users/wangzhe/.g/go/bin

.PHONY: gen
# generate error test code
gen:
	@cd ./gerr && protoc -I . \
		--proto_path=../third_party \
		--go_out=paths=source_relative:. \
		errors.proto

.PHONY: test
# generate error test code
test:
	@cd ./test && protoc -I . -I ../gerr \
 		--proto_path=../third_party \
		--go_out=paths=source_relative:. \
        --go-errorx_out=paths=source_relative:. \
		test.proto && \
	go test ./...

# sdk的git版本号
TagName = v1.0.2
gitag:
	git add . && git commit -m "$(TagName)"
	git push
	git tag -a $(TagName) -m "$(TagName)" # 创建带标签的Tag
	git push origin $(TagName)  # 推送Tag到远程
