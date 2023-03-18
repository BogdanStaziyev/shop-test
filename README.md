## Technical Assignment Completion

**I have completed a technical assignment in which I developed a CRUD entity for a customer in an online store, using
clean architecture and adhering to SOLID principles. I used basic authorization, which means that only the administrator
can perform operations on the entities. The login and password can be found in the env file or in Docker, but they are
not hardcoded, so they can be changed at any time.**

- Folder "cmd" contains the main entry point of the application.
- Folder "config" contains environment variables necessary for the application to function.
- Folder "internal" contains the "app" folder which has the file that starts all the necessary dependencies for the
  application to function. Migrations are also located here.
    - Folder "controller" contains everything for receiving data from the internet, including middleware for basic
      authorization, handlers, routers, and entities for requests and responses.
    - Folder "database" contains the file for working with the database.
    - All entities are located in the "entity" folder.
    - The services are located in the "service" folder.
- In the "migrations" folder, there are SQL schemas of entities
- The 'pkg' folder contains files with external dependencies for implementing clean architecture. They are replaced with
  interfaces and structures that implement them, allowing for easy substitution at any time

<img alt="struct" height="700" src="https://docs.google.com/uc?id=15O8wi5hE0pXKNALCQ_eG1ruqR7Qn8in4"/>

## Testing and Linting

**I also wrote tests for the service and handler, added Docker and Make files for ease of testing and linting checks.
Additionally, I added a Postman specification for convenience in testing.**

## Database Schema

**I designed a database for the main entities of the project: salesman, customer, products, and orders. In doing so, I
implemented various types of relationships: one-to-one, one-to-many, and many-to-many.**

![Database schema](https://docs.google.com/uc?id=1Gg0l0gVbG1rCs46ROo5zhxPDUgEU2M0d)

# Техническое задание для кандидатов:

1. разработка хттп сервера (задание обязательно к выполнению)

Представьте, что вы получили новый проект интернет-магазина и вам нужно заложить архитектуру для его разработки и
поддержки. В качестве тестовго задания полностью спроектируйте базу данных, а так же сделайте CRUD одной (любой)
сущности. HTTP сервер должен быть написан на GoLang, масимально просто, без использования фреймворков.

Техническое Задание:

- разработать HTTP API с базовой авторизацией, которое будет позволять выполнять CRUD операции над сущностями.
  Пользователь будет один (администратор, который и будет создавать эти сущности)
- формат ответа: JSON
- описание сущностей и полей (если вы считаете, что какого-то поля не хватает, вы можете смело его добавить):
    - продавец (имя, телефон)
    - товар (название, описание, цена, продавец)
    - покупатель (имя, телефон)
    - заказ (покупатель, несколько товаров)


2. оптимизация функции конкатенации. (задание со звездочкой, можно не делать, если не знаете)

Опмтимизируйте скорость выполнения функции. Кол-во значений во входящем параметре (len(str)) >= 30.
Напишите бенчмарк тест на эту функцию и на её оптимизированную версию.

```
func concat(str []string) string  {
    result := ""
    for _, v := range str {
        result += v
    }
    return result
}
```

Выполненное тестовое задание разместите на гитхабе.
Доступ к проекту предоставьте на аккаунт: https://github.com/Kirill-Shkodkin