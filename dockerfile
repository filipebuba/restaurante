# Define a imagem base
FROM golang:latest

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o código fonte para o diretório de trabalho
COPY . .

# Compila o código fonte
RUN go build -o main .

# Define o comando de inicialização do container
CMD ["./main"]

# https://sharmahimanshu1911.medium.com/golang-gin-mysql-in-docker-compose-869ca2b9b875