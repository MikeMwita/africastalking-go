
# Install commands
install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install -v github.com/go-critic/go-critic/cmd/gocritic@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

# Linting commands
lint:
	@echo "Running golangci-lint..."
	golangci-lint run || exit 1

	@echo "Running constructor-check..."
	constructor-check ./...

	@echo "Running go-consistent..."
	go-consistent ./...

	@echo "Running gosec..."
	gosec -fmt=golint -quiet ./...

	@echo "Running gocritic..."
	gocritic check ./...

	@echo "Running staticcheck..."
	staticcheck -checks all,-ST1000,-U1000,-tests=false ./...

	@echo "Running nilaway..."
	nilaway ./... || nilaway -exclude-pkgs "adomain.com","github.com/MikeMwita/africastalking-go/mocks" ./...
