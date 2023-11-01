start:
	cd "$(PWD)/backend" && make up
	cd "$(PWD)/frontend" && yarn dev

start-back:
	cd "$(PWD)/backend" && make up-log

start-front:
	cd "$(PWD)/frontend" && yarn dev

down:
	cd "$(PWD)/backend" && docker compose down
