FROM golang:1.17 as build

# Скопировать исходные файлы с хоста
COPY . /src
# Назначить рабочим каталог с исходным кодом
WORKDIR /src

# Собрать двоичный файл!
RUN go test
RUN CGO_ENABLED=0 GOOS=linux go build -o server


# Этап 2: Сборка образа со службой хранилища пар ключ/значение
#
# Использовать образ "scratch", не содержащий распространяемых файлов
FROM scratch
# Скопировать двоичный файл из контейнера build
COPY --from=build /src/server .
# Если предполагается использовать TLS, скопировать файлы .pem
#COPY --from=build /src/*.pem .
# Сообщить фреймворку Docker, что служба будет использовать порт 8080
EXPOSE 8081
# Команда, которая должна быть выполнена при запуске контейнера
CMD ["/server"]