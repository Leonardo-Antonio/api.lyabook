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
                                                <td align="center" style="line-height: 0px;">
                                                    <img style="display:block;  line-height:0px; font-size:0px; border:0px;padding: 0;"
                                                        src="https://i.ibb.co/kxx1Jcr/Group-211.png" width="100%"
                                                        height="100%" alt="logo">
                                                </td>
                                            </tr>



                                            <tr>
                                                <td align="center"
                                                    style="font-family: 'Roboto', sans-serif; font-size:17px; color:#000000; line-height:1.5; font-weight: 300; letter-spacing: 0px;">
                                                    <p style="width: 70%; padding-top: 2rem;padding-bottom: 1.5rem;">
                                                        Hola, <br>

                                                        Te traemos un descuento que puede interesarte <br>

                                                        Debido a su exitó y popularidad
                                                        quemos que nadie se pieda de ese maravilloso libro <br>

                                                        Asi que el equipo de LyaBook
                                                        decidio bajar el precio para que nadie se pierda de esta
                                                        historia
                                                    </p>
                                                </td>
                                            </tr>

                                            <tr align="center">
                                                <td bgcolor="#F9F9FF">
                                                    <table style="width: 100%;">
                                                        <tbody>
                                                            <tr align="center">
                                                                <td style="width: 50%; padding: 1rem 0;">
                                                                    <div style="width: 45%;">
                                                                        <div>
                                                                            <img src={{ index .ImagesSrc 0 }}
                                                                                alt="" width="100" height="150">
                                                                        </div>
    
                                                                        <div>
                                                                            
                                                                            <div>
                                                                                <p style="font-family: sans-serif;">
                                                                                    {{ .Description }}
                                                                                </p>
                                                                                <div style="padding: .5rem 0;">
                                                                                    <span style="font-family: sans-serif;">S/{{ .PriceCurrent }}</span>
                                                                                    <span style="text-decoration-line: line-through;font-family: sans-serif; font-size: .8rem;">S/{{ .PriceBefore }}</span>
                                                                                </div>
                                                                                <div>
                                                                                    <a href="#" style="color: #fff;font-family: sans-serif;font-weight: bold;width: 100px;text-decoration: none;">
                                                                                        <div style="background: #5E20E4;border-radius: 50px; width: 80%; padding: .4rem 0;">
                                                                                            Ver más
                                                                                        </div>
                                                                                    </a>
                                                                                </div>
                                                                            </div>
                                                                        </div>
                                                                    </div>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>

                                            <tr align="center">
                                                <td>
                                                    <div>
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