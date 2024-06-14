ARG BUILD_IMAGE=node
ARG BUILD_TAG=18

FROM ${BUILD_IMAGE}:${BUILD_TAG} as builder

WORKDIR /opt/web

COPY package.json ./
COPY yarn.lock ./

ENV PATH /opt/web/node_modules/.bin:$PATH

RUN yarn install

COPY . /opt/web

RUN yarn build

FROM nginx:latest

COPY --from=builder /opt/web/dist /usr/share/nginx/html

RUN rm -v /etc/nginx/nginx.conf
ADD nginx.conf /etc/nginx/

CMD ["nginx", "-g", "daemon off;"]