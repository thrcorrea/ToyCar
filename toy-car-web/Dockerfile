# Fase de Construção
FROM node:alpine as build

# Diretório de trabalho dentro do container
WORKDIR /app

# Copiando os arquivos de dependências
COPY package*.json ./

# Instalando as dependências
RUN npm install

# Copiando o restante dos arquivos do projeto
COPY . .

# Construindo o aplicativo
RUN npm run build

# Fase de Serviço
FROM nginx:alpine

# Copiando os arquivos construídos para o diretório de servir do Nginx
COPY --from=build /app/build /usr/share/nginx/html

# Expondo a porta 80
EXPOSE 80

# Iniciando o Nginx
CMD ["nginx", "-g", "daemon off;"]
