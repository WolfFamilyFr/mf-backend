# mf-backend

Pour le moment DB en local
- sudo docker start some-postgres

Copier le script SQL db.sql dans le conteneur
- sudo docker cp db.sql 1d4ae8512407:/

Executer le script SQL db.sql dans une DB postgres afin d'initialiser la base
- sudo docker exec -it 1d4ae8512407 bash
- psql -U postgres
- \i db.sql

Executer le backend en local
- air cmd/main.go -b 0.0.0.0