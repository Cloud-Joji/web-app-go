FROM node:18-alpine as builder

WORKDIR /app
COPY . .

WORKDIR /app/client
RUN npm i && npm run build

FROM golang:alpine

WORKDIR /app
COPY --from=builder /app/ .

CMD [ "go", "run", "." ]