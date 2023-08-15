BINARY_NAME = raytracer

.PHONY: all
all: test

.PHONY: build
build:
	go build -o $(BINARY_NAME) ./cmd/raytracer

.PHONY: test
test:
	go test ./...

.PHONY: watch
watch:
	@echo "Watching for file changes..."
	@echo "Files to watch: $(SRC_FILES)"
	@while true; do \
		clear; \
		${MAKE} test --no-print-directory; \
		inotifywait -qre close_write .; \
	done
