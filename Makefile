.PHONY:

build:
	@docker build -t $(TAG) .

run:
	@docker run -e BOT_API=$(TOKEN) $(TAG)