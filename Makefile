run:
	docker compose up --build

success:
	curl -iv localhost/success

bad_request:
	curl -iv localhost/bad_request

not_found:
	curl -iv localhost/not_found

internal_server_error:
	curl -iv localhost/internal_server_error