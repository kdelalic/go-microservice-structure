BASE = $(CURDIR)
MODULE = apigateway

# The rule that is called by our root Makefile during CI builds.
custom: buildLinux docker deploy

buildLinux:
	cd $(BASE)/services/$(MODULE) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./main-alpine

docker:
	cp $(BASE)/services/$(MODULE)/Dockerfile $(BASE)/Dockerfile.$(MODULE)
	docker build -f Dockerfile.$(MODULE) --rm -t $(MODULE):latest .

deploy:
	@chmod +x $(BASE)/scripts/docker_deploy.sh
	NAME=$(MODULE) $(BASE)/scripts/docker_deploy.sh
