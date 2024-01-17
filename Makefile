PATH_USE_CASE="./internal/app/usecase"
PATH_USE_CASE_TO="./internal/app/usecase"

ABS_PATH=$(CURDIR)
MAIN_PATH="./cmd/main/main.go"

mock_use_case:
	mockgen -source=$(PATH_USE_CASE)/$(src) -destination=$(PATH_USE_CASE)/$(to) -package=mocks

test:
	ABS_PATH=$(ABS_PATH) ENV=test go test ./...

run:
	ABS_PATH=$(ABS_PATH) go run $(MAIN_PATH)