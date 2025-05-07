up:
	docker compose --profile all up -d

up-app:
	docker compose --profile app up -d

down:
	docker compose down

tail-api:
	docker compose logs api -f --since=5m
