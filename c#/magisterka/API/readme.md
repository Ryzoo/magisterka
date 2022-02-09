docker build -t magisterka_c .
docker run --name magisterka_c --cpus=2 -m 2048m -p 8080:80 magisterka_c
