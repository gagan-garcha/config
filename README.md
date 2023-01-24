# About

This is a cli application which reads all the json files from the files directory and merges them together.
This application can be used to read all the config files and get the perticular config value from the given key.


## Requirements

The application uses Go version 1.19

## Start application

```console
go run main.go
```

You will be asked to enter the key name for which you want to retrive the value for.
Make sure the entered key is in lower case.

Example

```console
Enter Key: database
{"host":"127.0.0.1","password":"divido","port":3306,"username":"divido"}
```

If you want to retrive a value which is nested you can use the `.` operator for it.

Example

```json
{
    "environment": "development",
    "database": {
      "host": "127.0.0.1",
      "port": 3306,
      "username": "divido",
      "password": "divido"
    },
    "cache": {
      "redis": {
        "host": "127.0.0.1",
        "port": 6379
      }
    }
  }
 ```
 If you want to retrive the value of `redis` you can enter the key `cache.redis`.

```console
Enter Key: cache.redis
{"host":"127.0.0.1","port":6379}
```

The application will be stopped if you enter empty value.


