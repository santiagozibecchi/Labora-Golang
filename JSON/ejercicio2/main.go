package main

/*
Escribir un programa que levante un servidor en el puerto 9090. El programa debe mantener en memoria 
una variable entera definida a nivel “global”. Además configurar el servidor para tener dos endpoints.
POST /api/v1/inc: Incrementa la variable entera en uno
POST /api/v1/dec: Decrementa la variable entera en uno
GET /api/v1/curr: Devuelve el valor de la variable entera.
Una vez escrito este programa servidor, vas a escribir un programa que juegue como cliente y tenga 
dos gorutinas. En una vas a enviar 10000 peticiones para incrementar y otra 1000 peticiones para decrementar. 
Finalmente vas a mandar una petición para obtener el valor de la variable entera una vez que terminen las dos gorutinas.
Se pide ejecutar el programa servidor y luego :
Ejecutando una vez el programa cliente 
Ejecuta varias instancias del mismo programa cliente, es decir, ejecútalo varias veces en simultáneo.
Por último, se podrá apreciar que en el segundo caso para algún cliente el resultado que se le devuelve no es cero. ¿Esto a qué se debe? ¿ que hay que modificar para que a ambos clientes la petición “final” de obtener el valor de la variable entera devuelva un cero?

*/
func main() {
	
}