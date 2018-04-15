# build environment
FROM node:9.6.1 as builder
LABEL author="{{ cookiecutter.author }} <{{ cookiecutter.email }}"

RUN mkdir /app
ENV PATH /app/node_modules/.bin:$PATH
COPY package.json /app/package.json
RUN npm install --silent
COPY . /app
RUN npm run build

# production environment
FROM nginx:1.13.9
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]