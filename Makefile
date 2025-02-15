build:
	docker-compose -f docker-compose.yml build --no-cache

run:
	docker-compose -f docker-compose.yml up -d 

clear:
	docker-compose -f docker-compose.yml down --volumes --remove-orphans



