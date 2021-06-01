# API REST para la gestión de consultas

API REST para la gestión de consultas con GO.

## Requerimientos

**GO 1.16**

**PostgreSQL 12, 13**

## Instalación en local

1. Instalar la versión de GO 1.16 o posterior de acuerdo a su sistema operativo.

   [Descargar GO](https://golang.org/dl/)


2. Clonar el proyecto.


3. Crear la base de datos.


4. Descargar el ejecutable de `golang-migrate` para ejecutar las migraciones:

   [Descargar golang-migrate](https://github.com/golang-migrate/migrate/releases)


5. Guardar el archivo en la siguiente ruta:

   ```bash
   c:/go-migrate/migrate.windows-amd64.exe
   ```

   Si elige otra ruta para guardar el ejecutable debe configurar el archivo `migrate.sh`


6. Migraciones:

   Para ejecutar las migraciones abrir la consola `Git Bash` y usar el siguiente comando:

   ```bash
   export DBPWD=mypassword && . migrate.sh up
   ```
   Donde el valor de la variable de entorno `DBPWD` es la contraseña de la base de datos.


7. Configurar variables de entorno en el archivo `.env` (Ver el archivo `.env.example`).


9. Ejecutar el servidor local:

   ```bash
    go run main.go
   ```

## Esquema

Ver [Migraciones](https://github.com/luisgomez29/gestion-consultas-api/tree/main/migrations)

## Endpoints

La API REST proporciona las siguientes rutas:

### Auth

| Name | Route | Protected | Method |
| ---- | ----- | --------- | ------ |
| signup | /signup | No | POST |
| login | /login | No | POST |

### Users

| Name | Route | Protected | Method |
| ---- | ----- | --------- | ------ |
| Users list  | /users | Yes | GET |

## Author

**Luis Guillermo Gómez**

- [Github](https://github.com/luisgomez29)

```
Gracias!.
```