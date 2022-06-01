build:
	@$(MAKE) -C fetchprice build

clean:
	@$(MAKE) -C fetchprice clean

install:
	@$(MAKE) -C fetchprice install

uninstall:
	@$(MAKE) -C fetchprice uninstall

test: 
	@$(MAKE) -C fetchprice test