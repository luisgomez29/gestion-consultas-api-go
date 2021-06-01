# Migraciones

En la carpeta `migrations` es el entorno de migración. En el entorno de migración esta la carpeta `versions` que
contiene los scripts de versiones individuales.

Se usa [golang-migrate](https://github.com/golang-migrate/migrate) para hacer la migraciones a la base de datos
PostgreSQL.

Los archivos de migraciones tienen el formato:

```
{version}_{title}.up.sql
{version}_{title}.down.sql
```

Para mas información [Migrations](https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md).

# Esquema

El esquema inicial de la base de datos esta definido en el archivo `000001_initial.up.sql` para crearlo y
`00001_initial.down.sql` para eliminarlo.

Los campos `created_at` y `updated_at` tienen el valor por defecto `CURRENT_TIMESTAMP`. Esto se debe aplicar al momento
de crear nuevas tablas.

# Agregando nuevas migraciones

Se sigue el formato de nombre de archivo de la migración para editar el esquema de la base de datos, por ejemplo cambiar
el nombre al campo `username` de la tabla `users`:

```
000002_rename_username_field.up.sql
000002_rename_username_field.down.sql
```
