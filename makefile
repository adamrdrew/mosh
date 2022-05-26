.ONESHELL :

run-dev-daemon:
	mkdir -p mosh_tmp
	export MOSH_CONFIG_DIR=./mosh_tmp
	export MOSH_LOG_DIR=./mosh_tmp
	go run moshd.go 

run-dev:
	mkdir -p mosh_tmp
	export MOSH_CONFIG_DIR=./mosh_tmp
	export MOSH_LOG_DIR=./mosh_tmp
	go run mosh.go $(args)

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
	export MOSH_CONFIG_DIR=/etc/mosh
	export MOSH_LOG_DIR=/var/log/mosh
	mkdir -p /etc/mosh
	chmod 777 /etc/mosh
	mkdir -p /var/log/mosh
	chmod 777 /etc/mosh
	touch /var/log/mosh/moshd.log
	chmod 666 /var/log/mosh/moshd.log
	rm /etc/profile.d/mosh.sh
	touch /etc/profile.d/mosh.sh
	echo "export MOSH_CONFIG_DIR=/etc/mosh" > /etc/profile.d/mosh.sh
	echo "export MOSH_LOG_DIR=/var/log/mosh" > /etc/profile.d/mosh.sh