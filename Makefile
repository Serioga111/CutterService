im:
	docker build -t cutter-service .
	docker run -d --name cutter-im -p 8080:8080 cutter-service ./cutter -m=im

psql:
	docker-compose up -d --build

down:
	docker stop cutter-im || true
	docker rm cutter-im || true
	docker-compose down