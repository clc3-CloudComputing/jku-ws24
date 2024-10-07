# Pull in the layer of the base image: openjdk:8-jre-alpine
FROM openjdk:8-jre-alpine                             

# Add author label 
LABEL author=johannes.br@dynatrace.com

# Run make directory (mkdir) command
RUN mkdir /opt/app                                    

# Add a local file to the image 
COPY spring-music.jar /opt/app/app.jar                     

# Mark port for port publishing
EXPOSE 8080                                            

# Set the container startup command, i.e., execute the application 
ENTRYPOINT java -jar /opt/app/app.jar                