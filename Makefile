run:
	./run.sh


migration-up:
	dbmate --url postgres://postgres:password@localhost:5432/url_shortener?sslmode=disable up