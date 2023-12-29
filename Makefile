gen:
	goctl api go -api spider.api -dir .

doc:
	goctl api plugin -p goctl-swag="init" -api spider.api -dir swagger/spider

format:
	goctl api format --dir ./spider.api -declare

run:
	go run spider.go -f ./etc/spider-api.yaml

all:
	make gen && make doc && make format