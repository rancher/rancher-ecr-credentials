build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 govendor build -a -installsuffix cgo

test:
	govendor test -v +local

clean:
	rm rancher-ecr-credentials

push:
	docker build . -t rancher/rancher-ecr-credentials
