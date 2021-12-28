
dev:
	./script/dev.sh

dev-frontend:
	./script/dev.sh dev_frontend

dev-backend:
	./script/dev.sh dev_backend

image:
	./script/image.sh

clean:
	rm -fr gw
	rm -fr client/dist/*
	./script/uninstall.sh
