# go-url-shortener
[![GoDoc](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

Project created from the heart, aimed at showcasing the progress of my Go programming skills and acquired knowledge.
---

## Usage

To install the project, you can use Docker or set up a local MySQL database. Below are the commands to run the project:

#### Using Docker

```bash
docker-compose up
```

#### Local Installation with MySQL Database
1. Database Configuration
```bash
# Run the MySQL server locally
# MySQL User: root
# MySQL Password: root
# Database Name: url-shortener
# Database Port: 127.0.0.1:3306
```

2. Run the project:
```bash
go run .
```

## API Documentation

After starting the project, check the [Swagger Documentation](http://localhost:8080/swagger/index.html) for information on available endpoints, parameters, and responses.

## Technology Stack

- [GIN](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [Swaggo](https://github.com/swaggo/swag)
- MySQL

## License

This project is available under the MIT license. Details can be found in the LICENSE file.

## Contributing

#### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/bstembalski/go-url-shortener/issues) to report bugs or request new features.
