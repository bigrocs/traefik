.PHONY: git
git:
	git add .
	git commit -m"自动提交 git 代码"
	git push
.PHONY: tag
tag:
	git push --tags
.PHONY: user-api
user-api:
	go run user-api/main.go
.PHONY: build-user-api
build-user-api:
	docker build -f user-api/Dockerfile  -t user-api .

.PHONY: run
run:
	docker-compose up -d