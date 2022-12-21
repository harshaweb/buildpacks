
	FROM node:16
	WORKDIR ./projects/news-backend
	COPY package*.json ./
	COPY . ./app
	EXPOSE 8000
	RUN npm install
	CMD [ "node", "index.js" ]
	