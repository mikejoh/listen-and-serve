FROM alpine:latest
RUN apk add --no-cache \
    python3
COPY app/ app/
WORKDIR app/
RUN pip3 install flask
ENTRYPOINT ["python3"]
CMD ["app.py"]
