FROM node:20.11-alpine

WORKDIR /app

COPY package*.json ./
RUN npm ci --force

# Rebuild native dependencies for the current architecture
RUN npm rebuild lightningcss --force

EXPOSE 3000

CMD ["npm", "run", "dev"]