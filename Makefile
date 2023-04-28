test:
	go test -v -cover ./pkg/transaction && go test -v -cover ./pkg/bill && go test -v -cover ./pkg/budget && go test -v -cover ./pkg/goalgo test -v -cover ./pkg/db

clean:
	go clean

all:
	clean, test