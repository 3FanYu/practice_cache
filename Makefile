up: 
	docker-compose up -d --force-recreate

down: 
	docker-compose down

dev:
	docker-compose exec app bash

migrate:
	docker-compose exec app migrate -path=./database/migrations/ -database "mysql://root:root@tcp(db:3306)/glossika" up
