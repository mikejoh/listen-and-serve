# --- Build ---
FROM alpine:latest AS build
RUN apk add --no-cache nodejs
WORKDIR /build
COPY package.json .
RUN npm install

# --- Release ---
FROM build AS release
COPY --from=build /build .
COPY . .
CMD [ "npm", "run", "start" ]
