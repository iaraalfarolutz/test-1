# Test FIF Integraciones TLV
Se requiere una funcion en GO que lea una cadena de caracteres la cual contiene multiples campos en el formato TLV y genere un map[string]string con los campos TLV encontrados en la cadena.

## Formato de los campos TLV
Cada campo TLV esta compuesto por 3 partes:

 - **Tipo**: El tipo tiene un largo de 3 caracteres donde el primer caracter indica el tipo de dato  (A: Alfanumérico y N: Numérico) y dos caracteres para indicar el numero de campo Ejemplo: "01" o "15".
 - **Largo**: 2 caracteres que indican el largo del valor, este campo es importante puesto que indica cuantos caracteres leer a continuación.
 - **Valor**: Este es el valor del campo, el cual corresponde al valor del campo, su largo esta determinado por la porción **Largo**.

Ejemplo:

> Para "A0511AB398765UJ1N230200" Los campos son:
> - 05 de tipo Alfanumérico de largo 11 y valor AB398765UJ1
> - 23 de tipo Numérico de largo 2 y valor 00

# Compilar y ejecutar
- cd falabella_test1
- go install
- cd ..
- go build main.go
- ./main

# Test
- cd falabella_test1
- go test -coverprofile=coverage.out
- go tool cover -html=coverage.out (consultar coverage)

# CI/CD
- Hacer fork del proyecto.
- Ingresar a travis-ci.com e ingresar con Github.
- Aceptar la autorización de Travis.
- En el dashboard de Travis, seleccionar tu foto de perfil arriba en la derecha.
- Clickear el boton verde "Activate" y seleccionar el repo.
- Correr el job.
