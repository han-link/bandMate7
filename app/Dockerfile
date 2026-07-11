FROM node:24-alpine AS build

WORKDIR /app

COPY . ./

RUN npm ci && \
    npx expo export --platform web

FROM nginx:1.30.3

COPY --from=build /app/dist /usr/share/nginx/html

