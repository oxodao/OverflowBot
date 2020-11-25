FROM node:lts-alpine as build-frontend
WORKDIR /app
COPY frontend/ ./
RUN yarn install
RUN yarn build

FROM golang:alpine as build-backend

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY . .
COPY --from=build-frontend /app/dist /frontend/dist

RUN go mod download
RUN go get github.com/markbates/pkger/cmd/pkger
RUN pkger -include github.com/oxodao/overflow-bot:/frontend/dist
RUN apk add build-base
RUN go build -o overflow-bot

FROM alpine
WORKDIR /app
COPY --from=build-backend /app/overflow-bot /app/overflow-bot
ENTRYPOINT ["/app/overflow-bot"]
