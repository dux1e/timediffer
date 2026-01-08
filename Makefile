BINARY_NAME=timediff

.PHONY: build-linux
build-linux:
	@GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) .

.PHONY: build-windows
build-windows:
	@GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe .

.PHONY: install-linux
install-linux: build-linux
	@echo "installing timediff to /usr/local/bin"
	@sudo cp $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@sudo chmod +x /usr/local/bin/$(BINARY_NAME)
	@echo "$(BINARY_NAME) was installed sucessfully"

.PHONY: uninstall-linux
uninstall-linux:
	@echo "uninstalling $(BINARY_NAME) from /usr/local/bin"
	@sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "$(BINARY_NAME) was uninstalled successfully"
