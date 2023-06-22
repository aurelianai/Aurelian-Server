ARG env

FROM node:alpine as base
WORKDIR /aels
COPY . /aels
RUN npm i

FROM base as dev
ENTRYPOINT ["npm", "run", "dev"]


FROM base as prod
RUN npm run build
RUN npm prune --production
WORKDIR /aels/build
ENTRYPOINT ["node", "index.js"]


FROM ${env} as final