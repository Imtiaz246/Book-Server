RUN: BUILD
	./main --port=3000

# env credentials
ADMIN_USERNAME := admin
ADMIN_PASSWORD := password
JWT_SECRET_KEY := abcd

BUILD:
	go mod download
	echo "ADMIN_USERNAME=$(ADMIN_USERNAME)\nADMIN_PASSWORD=$(ADMIN_PASSWORD)\nJWT_SECRET_KEY=$(JWT_SECRET_KEY)\n" > .env
	go build -o main