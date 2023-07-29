test:
	go test -v -cover ./pkg/transaction ./pkg/bill ./pkg/budget ./pkg/goal ./db

clean:
	go clean

all:
	clean, test