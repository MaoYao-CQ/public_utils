<!DOCTYPE html>
<html>
<head>
    <title>{{.website.Title}}</title>
    <meta name="keywords" content='{{.website.Keywords}}'>
    <meta name="description" content='{{.website.Description}}'>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta property="og:locale" content="zh_CN">
    <meta property="og:title" content='{{.website.Title}}'>
    <meta property="og:site_name" content='{{.website.Name}}'>
    <meta property="og:description" content='{{.website.Description}}'>
    <meta property="og:type" content="website">
    <meta property="og:image" content='{{.website.PublicUrl}}/static/img/logo_utils.png'>
    <meta property="og:url" content='{{.website.PublicUrl}}'>

    <link rel="shortcut icon" href="/static/img/favicon.ico" type="image/x-icon"/>
    <link rel="stylesheet" href="/static/css/bootstrap.css">
    <script src="/static/js/jquery-3.6.0.min.js"></script>
    <script src="/static/js/bootstrap.js"></script>
    <script src="/static/js/bootstrap.bundle.min.js"></script>
</head>
<style>
    @media (max-width: 1200px) {
        body{
            font-size: 1.1rem;
        }
        .w-50{
            width: 100% !important;
        }
        .container{
            min-width: 100% !important;
        }
    }
</style>
<body class="h-100 body">

<nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top">
    <div class="container">
        <a class="navbar-brand" href="/">UTILS | 工具喵</a>
        <button type="button" id="collect" class="btn btn-secondary my-2 my-md-0">收藏</button>
    </div>
</nav>
<div class="d-block p-4"></div>