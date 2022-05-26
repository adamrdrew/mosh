build: clean
	mkdir bin
	go build -o bin/mosh mosh.go
	go build -o bin/moshd moshd.go

clean:
	go clean
	rm -rf bin

run-moshd:
	go run moshd.go

deps:
	go get

install: 
	go install mosh.go
	go install moshd.go
