ALL:
	go build ts.go
clean:
	rm ts
install:
	cp ts /bin/
