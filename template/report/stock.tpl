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
                {{ . }}
            </div>
        </div>

        <div>
            <h1 style="font-family: Saira; font-weight: 600; font-size: 2.5rem; text-align: center;">Reporte de libros
                sin stock</h1>
        </div>
        <div class="py-4 w-full flex justify-center">
            <span style="color: #CD7D7D; font-family: 'Roboto Condensed'; text-align: center;">Del {{ . }} al {{ .
                }}</span>
        </div>
    </div>

    <div class="py-8"></div>

    <div>
        <table class="table-fixed">
            <thead>
                <tr class="h-12" style="background-color: #5E20E4; color: #fff; font-size: 1.2rem;">
                    <th class="w-1/2" style="font-weight: bold;">Title</th>
                    <th class="w-1/4" style="font-weight: bold;">Author</th>
                    <th class="w-1/4" style="font-weight: bold;">Views</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr>
                <tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr><tr>
                    <td>Intro to CSS</td>
                    <td>Adam</td>
                    <td>858</td>
                </tr>
                <tr class="bg-blue-200">
                    <td>A Long and Winding Tour of the History of UI Frameworks and Tools and the Impact on Design</td>
                    <td>Adam</td>
                    <td>112</td>
                </tr>
                <tr>
                    <td>Intro to JavaScript</td>
                    <td>Chris</td>
                    <td>1,280</td>
                </tr>
            </tbody>
        </table>
    </div>
</body>

</html>