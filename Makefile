include frontend/Makefile
include backend/Makefile
start:
	cd "$(PWD)/backend" && docker compose build && docker compose up -d
	cd "$(PWD)/frontend" && yarn dev
