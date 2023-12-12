// Prácticas: Manejo de errores


// # Ejercicios

/*
## Ejercicio 1: Validación de entrada

Escribe un programa en Go que solicite al usuario dos números (numerador y denominador), intente realizar la división y maneje el caso en el que el
denominador sea cero. Si el denominador es cero, imprime un mensaje de error indicando que la división no es posible. En caso contrario, 
imprime el resultado de la división con dos decimales.
/*


/*
## **Ejercicio 2: Validación de entrada**

Crea una función que acepte un string y devuelva un error si el string está vacío o si es demasiado largo (por ejemplo, más de 100 caracteres).
*/

/*
## **Ejercicio 3: Conversión de tipos con manejo de errores**
Escribe una función que convierta un string a un número entero (**`int`**). La función debe devolver un error si el string no puede 
convertirse a un número. Prueba tu función con diferentes strings, incluyendo algunos que no sean números. Usa `panic` y `recover` para manejar los errores
*/

/*
## **Ejercicio 4: Manejo de errores en operaciones de archivos**

Escribe un programa que intente abrir un archivo cuyo nombre y directorio se pasa como argumento en la línea de comandos. Si el archivo no existe, 
el programa debe crearlo. Si hay algún otro error al abrir el archivo, el programa debe informar del error. No olvides cerrar el archivo
correctamente en todos los casos.
*/

/*
## Ejercicio 5: **Sistema de Calificaciones de Estudiantes**

**Objetivo**: Crear un programa en Go que gestione las calificaciones de los estudiantes utilizando mapas. y que maneje errores comunes y excepciones.
**Descripción**:

1. **Estructura del Estudiante**:
    - Define una estructura **`Estudiante`** que tenga al menos dos campos: **`Nombre`** (string) y **`ID`** (int) y `**calificaciones**` (mapa) con llave siendo la materia (string) y el valor siendo la nota (float)
2. **Funciones para Manejar Calificaciones**:
    - Escribe funciones para añadir un nuevo estudiante y sus calificaciones.
    - Implementa una función para actualizar las calificaciones de un estudiante existente.
    - Añade una función para eliminar las calificaciones de un estudiante.
    - Crea una función para calcular el promedio de calificaciones de un estudiante.
3. **Interacción con el Usuario**:
    - Permite que el usuario agregue, actualice, elimine y vea las calificaciones de los estudiantes a través de la entrada del terminal.
4. **Manejo de Errores en la Entrada de Datos**:
    - Al añadir o actualizar las calificaciones, asegúrate de que son números flotantes válidos. Si no lo son, muestra un mensaje de error.
    - Al seleccionar un ID de estudiante, verifica que sea un número entero válido.
    - Asegúrate de manejar posibles errores, como intentar actualizar las calificaciones de un estudiante que no existe.
5. **Funciones con Manejo de Errores**:
    - Modifica las funciones para añadir, actualizar, y eliminar estudiantes para que manejen situaciones como intentar actualizar un estudiante que no existe.
    - Cada función debe retornar un error si ocurre una situación excepcional.
6. **Implementación de la Interfaz de Usuario**:
    - Asegúrate de que el usuario reciba retroalimentación clara si algo va mal (por ejemplo, si intenta añadir un estudiante con un ID duplicado).

**Instrucciones Adicionales**:

- Puedes mejorar la interacción con el usuario implementando un menú simple en la terminal.
- Utiliza ciclos y declaraciones condicionales para validar la entrada del usuario.
- Practica escribir mensajes de error claros y útiles que ayuden al usuario a entender qué salió mal.
- Considera todos los casos posibles donde algo podría ir mal y cómo tu programa debería responder en esos casos.
*/