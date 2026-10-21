.PHONY: goose-create

goose-create:
	@goose -dir ./src/internal/postgres/migrations create $(filter-out $@,$(MAKECMDGOALS)) sql