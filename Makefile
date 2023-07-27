swagger-generate-api:
	oapi-codegen -config swagger/oapi-codegen-config.yaml swagger/swagger_3.yaml

swagger-validate:
	swagger validate swagger/swagger_3.yaml

db-generate:
	sqlboiler psql
