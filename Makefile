build:
	cd frontend && yarn install
	cd frontend && yarn build
	pkger -include github.com/oxodao/overflow-bot:/frontend/dist
	go build .