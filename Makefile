run:
	go run main.go

generate:
	swagger validate swagger.yml
	# specify the name to ensure we regenerate correctly.
	swagger generate client -f swagger.yml -A jira_rest_api

build:
	go build ./...

clean:
	rm -rf models/ clients/
