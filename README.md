# Proyecto 1 - Servicio Web

El objetivo del primer proyecto consiste en desarrollar un servicio web con información proveniente de una base de datos. Existe una lista de base de datos candidatas de las que podrá seleccionar cada grupo de estudiantes.

El funcionamiento del servicio web estará regido por la convención de manejo de recursos en servicios RESTful. Por tanto se deben desarrollar las operaciones para consultar, actualizar, modificar y eliminar diferentes recursos de la base de datos usando las técnicas utilizadas por dicho tipo de servicios.

### Modelo de datos

Desde la base de datos asignada se debe extraer la información que permita crear una serie de archivos individuales a partir de los registros de datos. Una estrategia es realizar una normalización de la base de datos y obtener las entidades individuales. De esta forma se obtendrían diferentes archivo CSV con información de las diferentes entidades.

### Administración de datos

Se debe crear la lógica del servicio de forma que permita consultar, actualizar, modificar y eliminar registros. Para ello se cargarán todos los registros en un arreglo de memoria y al finalizar la ejecución se escribirán nuevamente a los archivos. Note que si los datos no cambian no es necesario volverlos a escribir al archivo.

Es importante tomar en cuenta que será posible realizar la consulta y actualización de entidades que dependen de otras, por ejemplo para la base de datos de libros:

```
id,title,edition,copyright,language,pages,author,publisher
1,Operating System Concepts,9th,2012,ENGLISH,976,Abraham Silberschatz,John Wiley & Sons
2,Database System Concepts,6th,2010,ENGLISH,1376,Abraham Silberschatz,John Wiley & Sons
3,Computer Networks,5th,2010,ENGLISH,960,Andrew S. Tanenbaum,Pearson Education
4,Modern Operating Systems,4th,2014,ENGLISH,1136,Andrew S. Tanenbaum,Pearson Education
```
suponiendo que se crean tres entidades: libro, autor y editorial; se podrían consultar datos de la siguiente forma:

```
GET /autores/AbrahamSilbertschatz/libros

GET /editoriales/PearsonEducation/libros

GET /autores/AndrewSTanenbaum/libros/ComputerNetworks
```

también se podrá asociar un libro (ya creado) con su autor y su editorial:

```
POST /autores/AbrahamSilbertschatz/libros/ 
     {"id_libro":"StructuredComputerOrganization"}
	
POST /editoriales//PearsonEducation/libros/ 
    {"id_libro":"StructuredComputerOrganization"}
```

de igual forma se podrá borrar el enlace entre autor y editorial de la siguiente forma:

```
DELETE /autores/AbrahamSilbertschatz/libros/StructuredComputerOrganization

DELETE /editoriales//PearsonEducation/libros/StructuredComputerOrganization
```

### Aspectos Generales

* Se debe generar un servicio web que maneje toda esta información utilizando el software Golang. Posteriormente el sitio debe ser publicado en Heroku.
* El servicio debe incluir al menos tres entidades que se deben poder asociar mediante los URLs de la forma mostrada anteriormente.
* Se puede utilizar cualquier otra tecnología para subdividir la base de datos y generar los archivos de datos individuales que serán procesados por Golang.
* El proyecto debe ser realizado en forma individual o en parejas (bajo ningún caso se permitirá un grupo formado por tres personas)
* Cada grupo de estudiantes debe seleccionar la base de datos que utilizará en el proyecto desde la lista de bases de datos candidatas y informarlo al profesor lo antes posible. Se asignará la base de datos al primero que la solicite.
* Se deben crear una serie de scripts (en formato .bash o .sh) para probar cada funcionalidad del servicio desarrollado. Todos esos scripts se debe incluir en el repositorio del proyecto.
* Se debe crear una documentación en formato *markdown* en donde se describa el uso de cada una de las funciones del servicio web, la cual se debe agregar al repositorio.
* El proyecto se debe entregar mediante el Classroom de Github

### Bases de datos candidatas

Las bases de datos que pueden ser utilizadas son tomadas del sitio [Kaggle ](https://www.kaggle.com)y consisten de archivos CSV ó Excel.

* Space missions since the beginning of them (1957)
* US Police Shootings
* Women's International Football Results
* Goodreads-books
* TV shows on Netflix, Prime Video, Hulu and Disney+
* Game of Thrones
* Airplane Crashes Since 1908
* Volcanic Eruptions in the Holocene Period
* IMDB data from 2006 to 2016
* Data Police shootings
* US Mass Shootings


