Установка конфигурации `cp example.config.yml config.yml`

Запуск сервера `make run`

* `-config_dir_path` `GO_EXAMPLE_CONFIG_DIR_PATH` - путь к директории на файл `config.yml`

```
# Запрос на пагинацию
http://localhost:8080/posts?page=4&limit=1

# Запрос на сортировку
http://localhost:8080/posts?sort[text]=desc&sort[id]=asc
```
