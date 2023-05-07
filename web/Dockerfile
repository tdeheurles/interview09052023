# Utiliser l'image officielle Golang comme image de base
FROM golang:1.20-alpine as builder

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers go.mod et go.sum
# COPY go.mod go.sum ./
COPY go.mod ./

# Télécharger les dépendances
RUN go mod download

# Copier les fichiers source
COPY . .

# Compiler l'application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .










# Utiliser une image de base légère pour réduire la taille de l'image finale
FROM alpine:latest

# Copier l'exécutable compilé à partir de l'image précédente
COPY --from=builder /app/main /app/

# Exposer le port 8080 pour que le serveur puisse accepter les connexions
EXPOSE 8080

# Exécuter l'application
CMD ["/app/main"]
