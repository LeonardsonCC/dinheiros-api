url=http://localhost:8080

up:
	docker compose --profile all up -d

up-app:
	docker compose --profile app up -d

down:
	docker compose --profile all down

tail-api:
	docker compose logs api -f --since=5m

# Requests
ping:
	hurl --variable url=$(url) ./requests/ping.http | jq
# user
create-user:
	hurl --variable url=$(url) ./requests/users/create.http | jq
create-user-with-email:
	hurl --variable url=$(url) --variable email=test@test.com ./requests/users/create-with-email.http | jq
get-user:
	hurl --variable url=$(url) --variable email=test@test.com ./requests/users/get-by-email.http | jq

# accounts
create-account:
	hurl --variable url=$(url) ./requests/accounts/create.http | jq
get-accounts:
	hurl --variable url=$(url) ./requests/accounts/get-by-user.http | jq
update-account:
	hurl --variable url=$(url) --variable account_id=1 ./requests/accounts/update.http | jq
delete-account:
	hurl --variable url=$(url) --variable account_id=1 ./requests/accounts/delete.http | jq
