FROM nginx:stable-alpine

# Копирование файлов веб-клиента
COPY . /usr/share/nginx/html/

# Копирование конфигурации Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Открываем порт 80
EXPOSE 80

# Запуск Nginx в режиме переднего плана
CMD ["nginx", "-g", "daemon off;"] 