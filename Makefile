default:
	docker build -f dev.Dockerfile -t go-hexagonal .

up: default
	docker-compose up -d --remove-orphans

logs:
	docker-compose logs -f

down:
	docker-compose down --remove-orphans

test: 
	docker exec -it gohexagonal_app_1 go test ./... -v -cover

bench: 
	docker exec -it gohexagonal_app_1 go test ./... -bench=.
