dev:
	@docker-compose down && \
		docker-compose \
			-f docker-compose.yml \
			-f docker-compose.dev.yml \
			up -d --remove-orphans --build \
			&& docker-compose logs

web:
	@docker stop webserver && \
		docker-compose \
			-f docker-compose.yml \
			-f docker-compose.dev.yml \
			build server && \
		docker start webserver

heroku:
	#heroku container:login &&
	heroku container:push --app cw3guide web && \
	heroku container:release --app cw3guide web && \
	heroku logs --app cw3guide

logh:
	heroku logs --tail --app cw3guide