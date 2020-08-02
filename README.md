# golang-redis-basics
Implementación de paquete goredis para probar las operaciones basicas en REDIS.

# Como instalar redis
Primero, abra la terminal y actualice brew
$ brew update

Ahora, instale REDIS
$ brew install redis

Ahora que REDIS esta instalado hay que iniciar el servicio
$ brew services start redis

Verifiquemos que el servicio esta correctamente iniciado
$ redis-cli ping
por consola el resultado debe ser PONG, entonces todo bien a la hora de instalar redis en nuestra maquina local

# Ahora iniciemos las pruebas
Despues de clonar este repositorio comienza lo facil
Desde consola en la raiz
$ go run main.go

Deberíamos obtener este resultado

Creating a redis client connection
Inserting.... [key1, valueOfKey1]
Getting value of key1
Value is:  valueOfKey1

Como pueden observar se inicia creando una connección a nuestro REDIS localhost:6379, insertamos una llave [key1, valueOfKey1] y posteriormente la consultamos.








