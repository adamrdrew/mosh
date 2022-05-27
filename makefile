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
