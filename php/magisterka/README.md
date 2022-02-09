docker build -t magisterka_php .
docker run --name magisterka_php --cpus=2 -m 2048m -p 8080:8000 magisterka_php