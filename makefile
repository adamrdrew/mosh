.ONESHELL :

run-dev-moshd:
	mkdir -p mosh_tmp
	mkdir -p mosh_tmp/cache
	export MOSH_CONFIG_DIR=./mosh_tmp
	export MOSH_LOG_DIR=./mosh_tmp
	export MOSH_PID_DIR=./mosh_tmp
	export MOSH_PORT=9777
	export MOSH_CACHE_DIR=./mosh_tmp/cache
	go run moshd.go

stop-dev-moshd:
	kill `cat mosh_tmp/moshd.pid`

build: clean
	mkdir bin
	go build -o bin/mosh mosh.go
	go build -o bin/moshd moshd.go

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
