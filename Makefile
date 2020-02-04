all: run
run:
	go run server.go router.go --env=$(env)