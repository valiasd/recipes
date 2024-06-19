# Stage 1: Build the angular app
FROM node:22 as build

WORKDIR /app

COPY . .
RUN npm install
RUN npm run build --prod

# Stage 2: Serve the app with nginx server
FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
