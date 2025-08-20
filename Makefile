.PHONY: all build run stop clean swagger

SERVICES = admin authentication notification payment shipment

all: build run

build:
	@echo "Building all microservices..."
	@mkdir -p bin
	@for svc in $(SERVICES); do \
		if [ -f apps/$$svc/main.go ]; then \
			echo "Building $$svc..."; \
			go build -o bin/$$svc.exe apps/$$svc/main.go || { echo "Failed to build $$svc"; exit 1; }; \
		else \
			echo "Error: apps/$$svc/main.go not found"; \
			exit 1; \
		fi \
	done

run: build
	@echo "Starting all microservices..."
	@mkdir -p bin
	@for svc in $(SERVICES); do \
		if [ -f bin/$$svc.exe ]; then \
			echo "Starting $$svc..."; \
			./bin/$$svc.exe & echo $$! > bin/$$svc.pid; \
		else \
			echo "Error: Binary bin/$$svc.exe not found"; \
			exit 1; \
		fi \
	done
	@echo "All microservices started. Use 'make stop' to terminate."

stop:
	@echo "Stopping all microservices..."
	@for svc in $(SERVICES); do \
		if [ -f bin/$$svc.pid ]; then \
			taskkill /PID `cat bin/$$svc.pid` /F || true && rm -f bin/$$svc.pid; \
		fi \
	done
	@echo "All microservices stopped."

clean:
	@echo "Cleaning up binaries and Swagger docs..."
	@rm -rf bin/*
	@for svc in $(SERVICES); do \
		rm -rf apps/$$svc/docs; \
	done

swagger:
	@echo "Generating Swagger documentation..."
	@for svc in $(SERVICES); do \
		if [ -f apps/$$svc/main.go ]; then \
			echo "Generating Swagger for $$svc..."; \
			cd apps/$$svc && swag init || { echo "Failed to generate Swagger for $$svc"; exit 1; }; \
			cd ../..; \
		else \
			echo "Error: apps/$$svc/main.go not found"; \
			exit 1; \
		fi \
	done