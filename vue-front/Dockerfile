FROM node:22-alpine3.20 AS build-stage
# 作者信息
LABEL authors="xing.xiaolin@foxmail.com"

# 设置工作目录
WORKDIR /app

# 复制所有文件到工作目录
COPY . .

# 安装依赖
RUN npm install --registry=https://registry.npm.taobao.org

# 构建生产环境下到Vue项目
RUN npm run build

FROM nginx:stable-alpine

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=build-stage /app/dist /usr/share/nginx/html

EXPOSE 80

# 启动Nginx服务
CMD ["nginx", "-g", "daemon off;"]


