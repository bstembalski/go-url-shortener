# go-url-shortener
 [![GoDoc](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://badges.mit-license.org)


---

## Usage

Aby zainstalować projekt, można skorzystać z Dockera lub postawić bazę danych MySQL lokalnie. Poniżej znajdują się komendy do uruchomienia projektu:

#### Użycie Dockera

```bash
docker-compose up
```
#### Lokalna instalacja z bazą danych MySQL
1. Konfiguracja bazy danych
```bash
# Uruchom serwer MySQL lokalnie
# Uzytkownik MySQL: root
# Haslo MySQL: root
# Nazwa bazy danych: url-shortener
# Port bazy danych: 127.0.0.1:3306
```

2. Uruchomienie projektu:
```bash
go run .
```

## Dokumentacja API

Po uruchomieniu projektu, sprawdź [Dokumentację Swaggera](http://localhost:8080/swagger/index.html) dla informacji na temat dostępnych endpointów, parametrów i odpowiedzi.


## Technology Stack

- [GIN](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [Swaggo](https://github.com/swaggo/swag)
- MySQL

## Licencja

Ten projekt jest dostępny na podstawie licencji MIT. Szczegóły można znaleźć w pliku LICENSE.

## Contributing

#### Bug Reports & Feature Requests

Prosimy korzystać z [issue tracker](https://github.com/bstembalski/go-url-shortener/issues) do zgłaszania błędów lub prośb o nowe funkcje.