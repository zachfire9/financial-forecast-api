# financial-forecast-api

heroku create financial-forecast-api --buildpack heroku/go

The following env vars will need to be set in production environment:
DATABASE_CONNECTION - Database connection string
DATABASE_NAME - Name of database being used
TZ - Timezone being used