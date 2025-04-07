.DEFAULT_GOAL = local

run:
	CONFIG_PATH=config/config.yaml go run cmd/main.go

local:
	CONFIG_PATH=config/local.yaml go run cmd/main.go