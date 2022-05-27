.ONESHELL :
build: clean
	mkdir bin
	go build -o bin/mosh mosh.go

clean:
	go clean
	rm -rf bin

deps:
	go get

install:
	cp bin/* /usr/local/bin/
	mkdir -p /etc/mosh
	chmod 777 /etc/mosh
	mkdir -p /var/log/mosh
	chmod 777 /etc/mosh
	touch /var/log/mosh/moshd.log
	chmod 666 /var/log/mosh/moshd.log

test:
	mkdir -p mosh_test
	export MOSH_CONFIG_DIR=./mosh_test
	export MOSH_LOG_DIR=./mosh_test
	export MOSH_PID_DIR=./mosh_test
	export MOSH_PORT=9888
	export MOSH_CACHE_DIR=./mosh_test
	go clean -testcache
	go test ./...
	rm -rf ./mosh_test
