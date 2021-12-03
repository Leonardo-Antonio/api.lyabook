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

                <div>
                    <h1 style="font-family: Saira; font-weight: 600;font-size: 2rem;line-height: 2.5;">Libreria
                        <span style="color: #5E20E4;">“LyaBook”</span>
                    </h1>

                    <div>
                        <div>
                            <span
                                style="font-family: Roboto;font-weight: normal;color: rgba(127, 127, 127, 1);font-weight: normal;">Señor(a):
                            </span> <span style="color: #CD7D7D; font-family: 'Roboto Condensed';">
                                {{ range . }}
                                    {{ range .Client }}
                                        {{ .Name }} {{ .LastName }}
                                    {{ end }}
                                {{ end }}
                            </span>
                        </div>

                        <div>
                            <span
                                style="font-family: Roboto;font-weight: normal;color: rgba(127, 127, 127, 1);font-weight: normal;">DNI:
                            </span> <span style="color: #CD7D7D; font-family: 'Roboto Condensed';">
                                {{ range . }}
                                    {{ range .Client }}
                                        {{ if eq .Dni "" }}
                                            ---
                                        {{ else }}
                                            {{ .Dni }}
                                        {{ end }}
                                    {{ end }}
                                {{ end }}
                            </span>
                        </div>

                        <div>
                            <span
                                style="font-family: Roboto;font-weight: normal;color: rgba(127, 127, 127, 1);font-weight: normal;">Email:
                            </span> <span style="color: #CD7D7D; font-family: 'Roboto Condensed';">
                                {{ range . }}
                                    {{ range .Client }}
                                        {{ if eq .Email "" }}
                                            ---
                                        {{ else }}
                                            {{ .Email }}
                                        {{ end }}
                                    {{ end }}
                                {{ end }}
                            </span>
                        </div>
                    </div>
                </div>
            </div>

            <div class="p-4">
                <div class="w-full h-full flex items-center">
                    <div>
                        <div style="text-align: center;">
                            <h1 style="font-family: Saira; font-weight: 600;font-size: 1.5rem;line-height: 1.5;">R.U.C.
                                000000000000</h1>
                            <h1 style="font-family: Saira; font-weight: 600;font-size: 1.5rem;line-height: 1.5;">BOLETA
                                DE
                                VENTA
                            </h1>
                            <h1 style="font-family: Saira; font-weight: 600;font-size: 1.5rem;line-height: 1.5;">N°
                                {{ range . }}
                                    {{ .IdPayment }}
                                {{ end }}
                            </h1>
                        </div>

                        <div>
                            <div>
                                <span
                                    style="font-family: Roboto;font-weight: normal;color: rgba(127, 127, 127, 1);font-weight: normal;">Fecha
                                    de emisión:
                                </span> <span style="color: #CD7D7D; font-family: 'Roboto Condensed';">
                                    {{ range . }}
                                        {{ .CreateAtString }}
                                    {{ end }}
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="py-8"></div>

    <div>
        <table class="table-fixed">
            <thead>
                <tr class="h-12" style="background-color: #5E20E4; color: #fff; font-size: 1.2rem;">
                    <th style="width: 10%;font-weight: bold;">N°</th>
                    <th style="width: 40%;font-weight: bold;">Descripción</th>
                    <th style="width: 20%;font-weight: bold;">Tipo</th>
                    <th style="width: 10%;font-weight: bold;">Precio unitario</th>
                    <th style="width: 10%;font-weight: bold;">Cantidad</th>
                    <th style="width: 10%;font-weight: bold;">Importe</th>
                </tr>
            </thead>
            <tbody>
                {{ range . }}
                    {{ range $index, $book := .Products}}
                <tr>
                    <td>{{ $index }}</td>
                    <td class="text-center py-2">{{$book.Title}}</td>
                    <td class="text-center py-2">
                        {{ if eq $book.Description "d" }}
                            Digital
                        {{ else }}
                            {{ if eq $book.Description "f" }}
                                Fisico
                            {{ else }}
                                Fisico y Digital
                            {{ end }}
                        {{ end }}
                    </td>
                    <td class="text-center py-2">S/. {{$book.PriceUnit}}</td>
                    <td class="text-center py-2">{{$book.Quantity}}</td>
                    <td class="text-center py-2">S/. {{$book.Importe}}</td>
                </tr>
                    {{ end }}
                {{ end }}
            </tbody>
        </table>
        <h1>
            {{ range . }}
                <strong>Total a pagar</strong>:  S/. {{ .TotalPagar }}
            {{ end }}
        </h1>
    </div>
</body>

</html>