# Tahap 1: Build React/Vite
FROM node:22-alpine AS build

WORKDIR /app

# Salin package.json & lockfile untuk caching dependency
COPY package*.json ./
RUN npm ci

# Salin source code dan jalankan build
COPY . .
RUN npm run build

# Tahap 2: Runtime Web Server Nginx
FROM nginx:alpine

# Salin konfigurasi custom Nginx untuk SPA dan reverse proxy /api
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Salin hasil build dist ke direktori publik Nginx
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

