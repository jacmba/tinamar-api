FROM rust:alpine

RUN mkdir /src
RUN mkdir /app

WORKDIR /src
COPY . .
RUN cargo install --path .

RUN mv target/release/tinamar-api /app

WORKDIR /app
RUN rm -rf /src

CMD ./tinamar-api