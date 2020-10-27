TAG  = timberiodev/tcp_test_server:latest

docker-build:
	docker build --tag ${TAG} .

docker-run:
	docker run --interactive --tty --rm -p 9000:9000 ${TAG} -a 0.0.0.0:9000

docker-publish: docker-build
	docker push ${TAG}
