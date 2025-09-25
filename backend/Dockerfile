# --- Etapa 1: Builder ---
# Usamos una imagen de Go para compilar nuestra aplicación.
FROM golang:1.25-alpine AS builder

# Establecemos el directorio de trabajo dentro del contenedor.
WORKDIR /app

# Copiamos los archivos de dependencias primero para aprovechar la caché de Docker.
COPY go.mod go.sum ./
RUN go mod download

# Copiamos todo el resto del código fuente del backend.
COPY . .

# Compilamos la aplicación.
# Es MUY IMPORTANTE especificar la ruta a tu main.go.
# Las flags CGO_ENABLED=0 y GOOS=linux son buenas prácticas para crear un binario
# estático compatible con la imagen ligera de Alpine.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/api/main.go

# --- Etapa 2: Final ---
# Usamos una imagen mínima para mantener el tamaño final pequeño y seguro.
FROM alpine:latest

WORKDIR /app

# Copiamos ÚNICAMENTE el binario compilado desde la etapa 'builder'.
COPY --from=builder /app/main .

# Exponemos el puerto en el que correrá nuestra aplicación.
EXPOSE 8080

# El comando para ejecutar la aplicación cuando el contenedor inicie.
CMD ["./main"]