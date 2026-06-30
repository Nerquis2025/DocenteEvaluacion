# Deber 1: Encapsulación y Pruebas Unitarias

Proyecto en Go para la gestión de docentes, aplicando reglas de visibilidad y validación mediante pruebas unitarias internas y externas.

## Lista de Comandos Ejecutados y Resultados

Para ejecutar las pruebas en modo detallado y verificar cada caso de prueba:

```bash
go test ./docente -v

=== RUN   TestValidarEmailPrivado
--- PASS: TestValidarEmailPrivado (0.00s)
=== RUN   TestNormalizarNombrePrivado
--- PASS: TestNormalizarNombrePrivado (0.00s)
=== RUN   TestAgregarEvaluacionInterna
--- PASS: TestAgregarEvaluacionInterna (0.00s)
=== RUN   TestAccesoPublicoDocente
--- PASS: TestAccesoPublicoDocente (0.00s)
PASS
ok      DocenteEvaluacion/docente       1.702s

Porcentaje de Cobertura Alcanzado
Para comprobar el total de líneas de código probadas por los archivos de test:

go test ./docente -cover

ok      DocenteEvaluacion/docente       1.088s  coverage: 100.0% of statements

Reflexión Personal
Dificultades
Lo que más dolor de cabeza me dio fue adaptarme a cómo Go maneja lo privado y lo público. Como venía acostumbrado a usar palabras claras como private o public en otros lenguajes, cambiar el chip a que todo dependa solo de si la primera letra es mayúscula o minúscula me confundió bastante al principio y se me pasaba por alto al armar los métodos.

Aprendizajes
Me sirvió full ver el funcionamiento real de las pruebas de caja negra separando los archivos con package docente_test. Así pude comprobar directamente en la terminal que si intentas meter mano a un método privado desde afuera, el compilador te frena en seco y te rebota el código, lo que me ayudó a entender mejor cómo proteger la información.

Uso de IA
Usé la herramienta más que nada como un tutor para quitarme dudas de sintaxis cuando me saltaban errores raros en la consola y para que me dé una pista de qué líneas me faltaba probar cuando el porcentaje de cobertura no quería subir al 100%. 
