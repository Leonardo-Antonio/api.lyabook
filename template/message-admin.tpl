{{ define "message-admin" }}

<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link
        href="https://fonts.googleapis.com/css2?family=Baloo+Chettan+2:wght@500;600;700&family=Roboto:wght@300&display=swap"
        rel="stylesheet">
</head>

<body style="background-color: darkgrey; padding-top: 2rem; padding-bottom: 2rem;">
    <table width="800" align="center"
        style="background-image: url('https://i.ibb.co/RPKHmWV/Group-350.png'); background-repeat: no-repeat; background-size: cover; border-radius: 15px; background-color: #fff;">
        <tbody>
            <tr>
                <td align="center">
                    <table width="60%">
                        <tbody>
                            <tr>
                                <td align="center">
                                    <p><strong
                                            style="color: #5E20E4; font-size: 40px; font-family: 'Baloo Chettan 2';">!Nuevo mensaje¡</strong></p>
                                    <p style="font-size: 30px; font-family: 'Baloo Chettan 2';">Hola, <strong>{{ .Name }}</strong></p>

                                    <p style="font-size: 20px; font-family: Roboto;">
                                        {{ .Message }}
                                    </p>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </td>
            </tr>

            <tr style="height: 1rem;"></tr>

            {{ template "footer" }}
        </tbody>
    </table>
</body>

</html>

{{ end }}