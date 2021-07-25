docker_build:
	docker build -t manganato-cli .

docker_run:
	docker run -it -v ~/Desktop/manganato-cli:/root/Desktop/manganato-cli manganato-cli