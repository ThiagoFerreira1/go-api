FROM golang:1.23.4

# Definir o diretório de trabalho no contêiner
WORKDIR /go/src/app

# Copiar o código-fonte para o contêiner
COPY . .

# Expor a porta 8000
EXPOSE 8000

# Compilar o aplicativo Go
RUN go build -o main .

# Rodar o executável
CMD ["./main"]
