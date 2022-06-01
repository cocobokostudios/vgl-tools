build:
	@$(MAKE) -C price build

clean:
	@$(MAKE) -C price clean

install:
	@$(MAKE) -C price install

uninstall:
	@$(MAKE) -C price uninstall

test: 
	@$(MAKE) -C price test