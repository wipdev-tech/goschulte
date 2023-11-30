tw:
	tailwindcss -i ./static/input.css -o ./static/styles.css -w -m

fmt:
	go fmt ./...

lint: fmt
	golangci-lint run
