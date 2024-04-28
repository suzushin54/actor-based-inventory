.PHONY: setup
setup:
	make install
	make install-check
	go mod tidy

.PHONY: install
install:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

.PHONY: install-check
install-check:
	which buf grpcurl protoc-gen-connect-go
	# 🚀 If you see four paths, you're good to go!

.PHONY: buf
buf:
	buf lint
	buf generate

.PHONY: wire
wire: ## wire
	wire ./cmd/di
