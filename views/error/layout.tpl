<!doctype html>
<html lang="zh-hans">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{.title}}-{{.code}}</title>
    <style>
        body {
            margin: 0;
        }
        .error-page {
            width: 100vw;
            height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .content {
            position: relative;
            width: 1200px;
        }
        status_code_img.img {
            /*width: 1200px;*/
        }
        .summary {
            position: absolute;
            left: 0;
            bottom: 0;
            color: #18479D;
            /*width: 280px;*/
            /*text-align: center;*/
        }
        .summary .code {
            font-size: 120px;
            margin-bottom: 15px;
        }
    </style>
</head>
<body>
    <div class="error-page">
        <div class="content">
            <img class="status_code_img" src="/static/img/status_code/{{.code}}.svg">
            <div class="summary">
                <p class="code">{{.code}}</p>
                <p>{{.message}}</p>
            </div>
        </div>
    </div>
</body>
</html>
