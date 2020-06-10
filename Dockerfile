FROM scratch

WORKDIR /app
ADD ./service /app/service

CMD [ "/app/service" ]