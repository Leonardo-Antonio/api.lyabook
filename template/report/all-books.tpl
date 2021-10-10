<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href="https://unpkg.com/tailwindcss@^2/dist/tailwind.min.css" rel="stylesheet">
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=Roboto+Condensed:wght@700&family=Saira:wght@600&display=swap"
        rel="stylesheet">
</head>

<body style="padding: 0; margin: 0; box-sizing: border-box;">
    <div class="w-full" style="background-color: #F9F9FF;">
        <div class="flex justify-between">
            <div class="p-4">
                <img src="https://i.ibb.co/cbncH8H/Group-357.png" alt="">
            </div>

            <div class="p-4">
                {{ .Date }}
            </div>
        </div>

        <div>
            <h1 style="font-family: Saira; font-weight: 600; font-size: 2.5rem; text-align: center;">Reporte de libros
                y su stock</h1>
        </div>
        <div class="py-4 w-full flex justify-center">
            <span style="color: #CD7D7D; font-family: 'Roboto Condensed'; text-align: center;">Se puede visualizar el
                reporte de todos los libros, estado en la tienda y el stock en el almacen</span>
        </div>
    </div>

    <div class="py-8"></div>

    <div>
        <table class="table-fixed">
            <thead>
                <tr class="h-12" style="background-color: #5E20E4; color: #fff; font-size: 1.2rem;">
                    <th style="width: 40%;font-weight: bold;">Nombre</th>
                    <th style="width: 20%;font-weight: bold;">Editorial</th>
                    <th style="width: 10%;font-weight: bold;">Precio normal</th>
                    <th style="width: 10%;font-weight: bold;">Precio oferta</th>
                    <th style="width: 10%;font-weight: bold;">Tipo</th>
                    <th style="width: 10%;font-weight: bold;">Activo</th>
                </tr>
            </thead>
            <tbody>
                {{ range $index, $book := .Books }}
                <tr>
                    <td>{{ $book.Name }}</td>
                    <td class="text-center py-2">{{ $book.Editorial }}</td>
                    <td class="text-center py-2">S/{{ $book.PriceCurrent }}</td>
                    <td class="text-center py-2">S/{{ $book.PriceBefore }}</td>
                    <td class="text-center py-2">{{ $book.FormatBook }}</td>
                    <td class="text-center py-2">
                        <div class="flex justify-center">
                            {{ if eq $book.Active true }}
                            <div style="background-color: greenyellow; width: 1rem; height: 1rem; border-radius: 50px;">
                            </div>
                            {{ else }}
                            <div style="background-color: red; width: 1rem; height: 1rem; border-radius: 50px;">
                            </div>
                            {{ end}}
                        </div>
                    </td>
                </tr>
                {{ end }}
            </tbody>
        </table>
    </div>
</body>

</html>