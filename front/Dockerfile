# syntax=docker/dockerfile:1.4

# Stage 1: Build Stage
FROM node:18 AS build

WORKDIR /app

# 依存関係をインストール
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

# ソースコードをマウントしてビルド
COPY . .
RUN yarn build

# Stage 2: Production Image
FROM node:18-alpine

WORKDIR /app

# ビルド成果物をコピー
COPY --from=build /app/.next /app/.next
COPY --from=build /app/public /app/public
COPY --from=build /app/package.json /app/yarn.lock ./

# 依存関係をインストール（productionのみ）
RUN yarn install --production --frozen-lockfile

# アプリケーションを提供
CMD ["yarn", "start"]

