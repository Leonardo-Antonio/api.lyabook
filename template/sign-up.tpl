{{ define "sign-up" }}

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

    <table width="800" align="center" style="background-color: #fff; border-radius: 20px;">
        <tbody>
            <tr>
                <td align="center">
                    <img src="https://i.ibb.co/fHDb5Fg/logo-1.png" alt="" width="110" style="padding: 1rem 0;">
                </td>
            </tr>
            <tr>
                <td align="center">
                    <table width="60%">
                        <tbody>
                            <tr>
                                <td align="center">
                                    <p><strong
                                            style="color: #5E20E4; font-size: 40px; font-family: 'Baloo Chettan 2';">Bienvenido
                                            a Lyabook</strong></p>
                                    <p style="font-size: 20px; font-family: 'Baloo Chettan 2';">Descubre la cantidad de
                                        libros de diversas categorias que tenemos para ofrecerles.</p>
                                    <p style="font-size: 18px; color: #7F7F7F;font-family: 'Baloo Chettan 2';">
                                        Recuerda que tenemos una sección para que usted pueda compartir sus historias,
                                        poemas, novelas u otros textos.
                                    </p>
                                    <div style="height: .8rem;"></div>
                                    <a href="https://lyabook.hostman.site/" style="text-decoration: none;">
                                        <span
                                            style="background-color: #5E20E4; color: #fff;font-family: 'Baloo Chettan 2'; border-radius: 50px; height: 4rem;padding: .6rem;">Visitanos</span>
                                    </a>
                                    <div style="height: .8rem;"></div>
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