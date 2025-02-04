FROM node:20 AS build-frontend
WORKDIR /build

COPY ./client/package.json .
COPY ./client/pnpm-lock.yaml .

RUN npm install -g pnpm
RUN pnpm i --frozen-lockfile

COPY ./client .
RUN pnpm build

FROM golang:1.23 AS build

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=build-frontend /build/dist ./client/dist

RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -o ./bin/encurtago ./cmd/api/main.go

FROM alpine:3.14

COPY --from=build /build/bin/encurtago /usr/bin/encurtago

EXPOSE 8080

CMD [ "/usr/bin/encurtago" ]




