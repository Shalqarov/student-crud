run:
	# Копирование можно убрать, поставил, чтобы время не терять создавая новую env
	cp example.env .env
	docker compose up -d
stop:
	docker compose down