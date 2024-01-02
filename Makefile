GO_BIN := $(shell go env GOPATH)/bin
GO_VER := $(shell go version | grep -Eo '[.]+[0-9]+[.]' | grep -Eo '[^.][0-9]')

M = $(shell printf "\033[33;1m▶︎\033[0m")
MOCKERY = $(GO_BIN)/mockery

.PHONY: test
test: cc ; $(info $(M) tests…)
	@go test -v ./...

.PHONY: mocks
mocks: clean.mocks ; $(info $(M) generate mocks…)
	@ # interfaces
	$(MOCKERY) --dir=./domains/interfaces --name=UserRepository --filename=mock_user_repository_interfaces.go --output=./mocks/interfaces --outpkg=mocks
	$(MOCKERY) --dir=./domains/interfaces --name=BandRepository --filename=mock_band_repository_interfaces.go --output=./mocks/interfaces --outpkg=mocks
	$(MOCKERY) --dir=./domains/interfaces --name=PostRepository --filename=mock_post_repository_interfaces.go --output=./mocks/interfaces --outpkg=mocks

.PHONY: clean.mocks
clean.mocks: ; $(info $(M) remove mocks…)
	@rm -rf ./mocks

.PHONY: cc
cc: ; $(info $(M) clean go cache…)
	@go clean -cache
