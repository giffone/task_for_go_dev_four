build:
	docker-compose build --no-cache

run:
	docker-compose up --build

stop:
	docker-compose down --remove-orphans

clear:
	docker-compose rm -f
	docker-compose pull