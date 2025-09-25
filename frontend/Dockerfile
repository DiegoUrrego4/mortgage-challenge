# frontend/Dockerfile

# --- Etapa 1: Build ---
# Usamos una imagen de Node para instalar dependencias y construir la app
FROM node:23.6-alpine AS builder
WORKDIR /app

# Copiamos los archivos de manifiesto y el lockfile
COPY package.json pnpm-lock.yaml ./

# Instalamos pnpm y las dependencias
RUN npm install -g pnpm
RUN pnpm install

# Copiamos el resto del código fuente
COPY . .

# Construimos la aplicación para producción
RUN pnpm build

# --- Etapa 2: Serve ---
# Usamos un servidor web Nginx súper ligero para servir los archivos estáticos
FROM nginx:1.25-alpine

# Copiamos los archivos construidos desde la etapa 'builder'
COPY --from=builder /app/dist /usr/share/nginx/html

# Copiamos una configuración personalizada de Nginx (la crearemos a continuación)
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Exponemos el puerto 80, que es el que Nginx usa por defecto
EXPOSE 80

# Comando para iniciar Nginx
CMD ["nginx", "-g", "daemon off;"]