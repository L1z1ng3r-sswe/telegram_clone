SRC = app/cmd/telegram_clone/main.go
DESTINATION_BUILD = app/cmd/telegram_clone/build
BINARY_APP = binary_app
HOST_IP = 172.17.0.1

build:
	go build -o ./$(DESTINATION_BUILD)/$(BINARY_APP) $(SRC)

run: build
	./$(DESTINATION_BUILD)/$(BINARY_APP)

migrate-up:
	docker run -v $(PWD)/migrations:/migrations migrate/migrate \
	  -path=/migrations/ \
		-database "postgres://postgres:asyl12345.@$(HOST_IP):5432/postgres?sslmode=disable" \
		up

migrate-down:
	echo "y" | docker run -i -v $(PWD)/migrations:/migrations migrate/migrate \
	  -path=/migrations/ \
		-database "postgres://postgres:asyl12345.@$(HOST_IP):5432/postgres?sslmode=disable" \
		down
