url=http://localhost:8080

up:
	docker compose --profile all up -d

up-app:
	docker compose --profile app up -d

down:
	docker compose down

tail-api:
	docker compose logs api -f --since=5m

# Requests
ping:
	hurl --variable url=$(url) ./requests/ping.http | jq
create-user:
	hurl --variable url=$(url) ./requests/users/create.http | jq
get-user:
	hurl --variable url=$(url) --variable email=test@test.com ./requests/users/get-by-email.http | jq
