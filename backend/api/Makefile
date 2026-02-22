.PHONY: clear-swag swag-init

clear-swag:
	@echo "Clearing swag files..."
	del /Q docs\* 2>null || exit 0

swag-init: clear-swag
	@echo "Initializing swag files..."
	swag init -g main.go