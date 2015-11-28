# VARIABLES
PACKAGE="github.com/SDkie/metric_collector"
BINARY_NAME="metric_collector"

print_success = echo -e "\e[1;32m$(1) $<\e[0m"
print_warning = echo -e "\e[1;33m$(1) $<\e[0m"
print_error = echo -e "\e[1;31m$(1) $<\e[0m"

# ENV Variables
export GOPATH=$(shell pwd)/../../../..
export PORT=8080
export MODE=debug
export RABBITMQ_URI=amqp://guest:guest@localhost:5672

# Default target : Do nothing
default:
	@echo "You must specify a target with this makefile"
	@echo "Usage : "
	@echo "make clean        Remove binary files"
	@echo "make install      Compile sources and build binaries"
	@echo "make run          Build application and run it !"

# Clean .o files and binary
clean:
	@echo "--> cleaning..."
	@go clean || (echo "Unable to clean project" && exit 1)
	@rm -rf $(GOPATH)/bin/$(BINARY_NAME) 2> /dev/null
	@echo "Clean OK"

# Compile sources and build binary
install: clean
	@echo "--> installing..."
	@go install $(PACKAGE) || ($(call print_error,Compilation error) && exit 1)
	@echo "Install OK"

# Run your application
run: clean install
	@echo "--> running application..."
	@$(GOPATH)/bin/$(BINARY_NAME)
