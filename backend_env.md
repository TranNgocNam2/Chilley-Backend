## ENV FILE
### APP_HOST=0.0.0.0
### APP_PORT=****
### CORS_DEV=....
### CORS_PROD=....

## BUILD DOCKER IMAGE
### docker build -t <image-name> .

## RUN DOCKER IMAGE
### docker run --env-file .env -p 3000:3000 <image-name>
