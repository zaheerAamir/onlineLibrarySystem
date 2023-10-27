FROM nginx:latest

# Create a directory to store wait-for-it script
WORKDIR /usr/src/app

# Download the wait-for-it script
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /usr/src/app/wait-for-it.sh

COPY nginx/swagger-ui /usr/src/app/swagger-ui/

# Make the script executable
RUN chmod +x /usr/src/app/wait-for-it.sh