all: run
run:
	go run server.go router.go database.go schema_queries.go controllers.go models.go --env=$(env)