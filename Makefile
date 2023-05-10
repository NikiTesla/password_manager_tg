PROJECTNAME=$(shell basename "$(PWD)")
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

build:
	@ echo "  >  building all for $(PROJECTNAME)..."
	@ go build -o ./.bin/$(PROJECTNAME) cmd/main.go
	@ sudo docker start vk_pg
	@ sleep 0.1
	@ ./.bin/$(PROJECTNAME)

run:
	@ echo "  >  running cmd/main.go file..."
	@ sudo docker start vk_pg
	@ sleep 0.1
	@ go run cmd/main.go

docker:
	@ sudo docker rmi -f $(PROJECTNAME)
	@ echo "  >  making docker container $(PROJECTNAME)..."
	@ sudo docker build -t $(PROJECTNAME) .
	@ sudo docker compose up

# usage make migration-up ARGS="[version]" 
migration-up:
	@ echo "  >  making migrations"
	@ sudo docker start vk_pg
	@ sleep 0.1
	@ cat schemas/$(ARGS)_init.up.sql | sudo docker exec -i vk_pg  psql -U postgres -d telegram

# usage make migration-down ARGS="[version]" 
migration-down:
	@ echo "  >  making migrations"
	@ sudo docker start vk_pg
	@ sleep 0.1
	@ cat schemas/$(ARGS)_init.down.sql | sudo docker exec -i vk_pg  psql -U postgres -d telegram