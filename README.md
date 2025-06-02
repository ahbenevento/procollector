# Busca y recopila proyectos

Desarrollado en **Go** para buscar y recopilar proyectos en distintas carpetas repositorios de código.

No solo imprime en pantalla los proyectos encontrados sino que puede crear un archivo CSV o incluso JSON con la lista completa.

El formato JSON utilizado es compatible con la extensión **Project Manager** ([vscode-project-manager](https://github.com/alefragnani/vscode-project-manager)) de **Visual Studio Code**.

De forma predeterminada la información de los proyectos debe estar alojada en un archivo INI con alguno de los siguientes nombres:

- `.project`
- `project.ini`

## Algunos ejemplos

### Buscar proyectos y etiquetarlos con **Go** si se encuentran dentro de una determinada subcarpeta

Guarda los proyectos con formato CSV.

```console
procollector -csv projects.csv -t "Go=/dev/go/" ~/dev
```

### Buscar proyectos y etiquetarlos según distintas subcarpetas

Guarda los proyectos con formato JSON.

```console
procollector -json projects.json -t "local=/dev/" -t "net=/repo/web/" ~/dev /mnt/repo/web
```

## Más información

Ejecute simplemente el comando (`procollector`) para conocer la lista completa de parámetros.

## Formato del archivo INI

```ini
name = Nombre del proyecto
tag = etiqueta1, etiqueta2
; disabled = true
```
