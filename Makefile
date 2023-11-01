include frontend/Makefile
include backend/Makefile
start:
	cd "$(PWD)/backend" && make up
	cd "$(PWD)/frontend" && yarn dev

start-log:
	cd "$(PWD)/backend" && make up-log
	cd "$(PWD)/frontend" && yarn dev

down:
	cd "$(PWD)/backend" && docker compose down
