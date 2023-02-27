run:
	cp example.env .env
	docker compose up -d
stop:
	docker compose down