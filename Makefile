build-image:
	docker build -t guifgr/finance .

run-app:
	docker-compose -f .devops/app.yml up -d

lint:
	golint ./...
	go fmt -n ./...

unique_test:
	go test ./...