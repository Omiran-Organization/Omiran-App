
FROM node:13.12.0 as build

WORKDIR /usr/src/app/
ENV TZ America/New_York
ENV PATH /usr/src/app/node_modules/.bin:$PATH
COPY package.json . /usr/src/app/
RUN npm install
COPY . ./
EXPOSE 3000

# DO THIS IN PROD
# RUN npm run build
# CMD npm run start
CMD npm run dev

