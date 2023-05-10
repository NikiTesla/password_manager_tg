PROJECTNAME=$(shell basename "$(PWD)")
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

build:
	@ echo "  >  building all for $(PROJECTNAME)..."
	@ go build -o ./.bin/$(PROJECTNAME) cmd/main.go
	@ sudo docker start vk_telegram_pg
	@ sleep 0.1
	@ ./.bin/$(PROJECTNAME)

run:
	@ echo "  >  running cmd/main.go file..."
	@ sudo docker start vk_telegram_pg
	@ sleep 0.1
	@ go run cmd/main.go

docker:
	@ sudo docker rmi -f $(PROJECTNAME)
	@ echo "  >  making docker container $(PROJECTNAME)..."
	@ sudo docker build -t $(PROJECTNAME) .
	@ sudo docker compose up