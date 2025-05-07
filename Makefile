up:
	docker compose up -d

down:
	docker compose down

tail-api:
	docker compose logs api -f --since=5m
