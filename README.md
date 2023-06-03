
# Desafío GO Bases
### Programa de Análisis de Ventas de Pasajes de una Aerolínea


## Autoras

- [EFFM00](https://github.com/EFFM00)
- [rainvare](https://github.com/rainvare)


## Documentación
El Programa de Análisis de Ventas de Pasajes de una Aerolínea es una herramienta diseñada para procesar y analizar los datos de ventas de pasajes en un archivo CSV proporcionado por la aerolínea. Este programa permite calcular diferentes datos estadísticos relacionados con las tendencias de compra de los pasajeros, como el número total de personas que viajan a un país determinado, la cantidad de personas que viajan en diferentes períodos del día y el promedio de personas que viajan a un país en un día determinado.
## Requerimiento 1: Función para calcular el número total de personas que viajan a un país determinado

La función GetTotalTickets(destination string) se utiliza para calcular cuántas personas viajan a un país específico. Toma como entrada el nombre del país de destino y devuelve el número total de pasajes reservados hacia ese país. En caso de que ocurra algún error durante el cálculo, la función también devuelve un error.

```bash
func (a Airline) GetTotalTickets(destination string) (int, error) {}
```
## Requerimiento 2: Función para calcular el número de personas que viajan en diferentes períodos del día

El programa también proporciona una o varias funciones para calcular la cantidad de personas que viajan en diferentes períodos del día. La función GetCountByPeriod(time string) se utiliza para este propósito y toma como entrada un período de tiempo en formato de texto. Los períodos de tiempo disponibles son madrugada (0 a 6), mañana (7 a 12), tarde (13 a 19) y noche (20 a 23). Esta función devuelve el número de pasajeros que viajan durante el período de tiempo especificado. Si se produce algún error durante el cálculo, la función también devuelve un error.

```bash
func (a Airline) GetCountByPeriod(time string) (int, error) {}
```
## Requerimiento 3: Cálculo del promedio de personas que viajan a un país en un día determinado

Además, el programa permite calcular el promedio de personas que viajan a un país específico en un día determinado. Esta información es útil para analizar las tendencias diarias de compra hacia un país en particular. El cálculo del promedio se realiza dividiendo el número total de pasajes vendidos hacia ese país en un día entre el número total de días considerados. El resultado es el promedio diario de pasajeros que viajan a ese país.

```bash
func (a Airline) AverageDestination(destination string) (float64, error) {}
```
## Ejecución del programa

1. Abre una terminal o consola.

2. Utiliza el siguiente comando para navegar al directorio donde se encuentra el archivo del programa Go:

Reemplaza "/ruta/al/directorio" con la ruta real del directorio donde se encuentra el archivo.

```bash
cd /ruta/al/directorio
```

3. Asegúrate de tener Go instalado en tu sistema. Puedes verificarlo ejecutando el siguiente comando en la consola:

```bash
go version
```

4. Ejecuta el siguiente comando para compilar y ejecutar el programa Go:

```bash
go run main.go
```

Go compilará y ejecutará automáticamente el programa, mostrando la salida en la consola.
## Demo

Salida en consola

![demo](https://imgur.com/a/EGa1GhB)
