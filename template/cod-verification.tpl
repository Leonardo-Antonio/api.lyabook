<!DOCTYPE html>
<html lang="en">
<!--  Name - Description - Slug - PriceCurrent - PriceBefore  -->

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>

<body bgcolor="#dddddf">
    <table width="100%" border="0" align="center" cellpadding="0" cellspacing="0">

        <!-- START HEADER/BANNER -->

        <tbody>
            <tr>
                <td align="center" bgcolor="#dddddf">
                    <table style="padding: 1rem;" class="col-600" width="600" border="0" align="center" cellpadding="0"
                        cellspacing="0">
                        <tbody>
                            <tr>
                                <td align="center" valign="top" bgcolor="#fff"
                                    style="box-shadow: 0px 4px 20px rgba(255, 11, 89, 0.15);height: 100%;background-size:cover; background-position:top; border-radius: 20px;"
                                    height=600>
                                    <table class="col-600" width="600" height="auto" border="0" align="center"
                                        cellpadding="0" cellspacing="0">

                                        <tbody>

                                            <tr>
                                                <td align="center" style="line-height: 0px;padding: 0 2rem;">
                                                    <img style="display:block;  line-height:0px; font-size:0px; border:0px;padding: 0;"
                                                        src="https://i.ibb.co/jw0D3Br/Group-198.png" width="100%"
                                                        height="100%" alt="logo">
                                                </td>
                                            </tr>



                                            <tr>
                                                <td align="center"
                                                    style="font-family: 'Roboto', sans-serif; font-size:17px; color:#000000; line-height:20px; font-weight: 300; letter-spacing: 0px;">
                                                    <p style="width: 70%; padding-top: 2rem;padding-bottom: 1.5rem;">
                                                        Hola, <span style="color: blue;">{{ .Name }}</span>.

                                                        Necesitamos su confirmación para otorgar acceso a la página que
                                                        está intentando ver.

                                                        Regrese a la página y use el token a continuación para iniciar
                                                        sesión:
                                                    </p>
                                                </td>
                                            </tr>

                                            <tr align="center">
                                                <td>
                                                    <table style="width: 100%;">
                                                        <tbody>
                                                            <tr align="center">
                                                                <td style="width: 50%;">
                                                                    <div
                                                                        style="background-color: #F9F9FF; width: 200px; padding: 1.5rem 0; border-radius: 15px; font-size: 1.5rem;">
                                                                        <strong
                                                                            style="font-family: Saira, sans-serif; font-weight: bold;">{{
                                                                            .VerificationCode }}</strong>
                                                                    </div>

                                                                    <div style="padding-top: 1.5rem; text-align: start; width: 200px;">
                                                                        <span>
                                                                            <strong style="font-family: sans-serif;">
                                                                                Gracias, <br>
                                                                                Equipo LyaBook
                                                                            </strong>
                                                                        </span>
                                                                    </div>
                                                                </td>
                                                                <td style="width: 50%;">
                                                                    <img src="https://i.postimg.cc/wxNSMksG/undraw-Fingerprint-re-uf3f-1-1.png"
                                                                        alt="">
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>

                                            <tr align="center">
                                                <td>
                                                    <div style="padding-top: 2rem;">
                                                        <div
                                                            style="border-radius: 0 0 15px 15px;background-color: #021639; padding: 1rem 0;">
                                                            <span>
                                                                <strong style="color: #fff; font-family: sans-serif;">En
                                                                    LyaBook, lee y aprende
                                                                    más</strong>
                                                            </span>
                                                        </div>
                                                    </div>
                                                </td>
                                            </tr>

                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </td>
            </tr>
        </tbody>
    </table>
</body>

</html>