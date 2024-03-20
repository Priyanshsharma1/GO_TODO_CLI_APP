# Define variables
APP = ./cmd/todo
RUN = go run $(APP)
RUN_LIST = $(RUN) -list
RUN_ADD = $(RUN) -add
RUN_DELETE = $(RUN) -delete
RUN_COMPLETE = $(RUN) -complete

# Define phony targets
.PHONY: run list add delete complete

# Default target
default: run

# Run the application
run:
	@$(RUN)

# List all items
list:
	@$(RUN_LIST)

# Add a new item
add:
	@$(RUN_ADD)

# Delete an item by index
delete:
	@if [ -z "$(index)" ]; then \
		echo "Error: Missing 'index' parameter. Usage: make delete index=<index>"; \
	else \
		$(RUN_DELETE) $(index); \
	fi

# Mark an item as completed by index
complete:
	@if [ -z "$(index)" ]; then \
		echo "Error: Missing 'index' parameter. Usage: make complete index=<index>"; \
	else \
		$(RUN_COMPLETE) $(index); \
	fi