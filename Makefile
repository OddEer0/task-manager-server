PATH_USE_CASE="./internal/app/usecase"
PATH_USE_CASE_TO="./internal/app/usecase"

mock_use_case:
	mockgen -source=$(PATH_USE_CASE)/$(src) -destination=$(PATH_USE_CASE)/$(to) -package=mocks
