<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  
  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    } 
  </style>
</head>

<body>
 
    <h1 class="logo">测试</h1>
	  <a href=/appser/>首页</a>
 
 数据提交:<br/>
 <form action="/appser/?act=save"  method="post"  >
 id <input type="text" name="id"><br>
title: <input type="text" name="title"><br>
modname <input type="text" name="modname"><br>
cateid <input type="text" name="cateid"><br>
cateid2 <input type="text" name="cateid2"><br>
tags <input type="text" name="tags"><br>
<input type="submit">
 </form>
 数据取回（无内容，只有标题）:<br/>
  <form action="/appser/?act=load"  method="post"  >
 load data:<br>
modname <input type="text" name="modname"><br>
 cateid <input type="text" name="cateid"><br>
  cateid2 <input type="text" name="cateid2"><br>
<input type="submit">
 </form>

  数据取回（一条，有内容）,用于有标题时，用一个id取回正文。:<br/>
    <form action="/appser/?act=loadone"  method="post"  >
 load data:
id <input type="text" name="id"><br>
 
<input type="submit">
 </form>
</body>
</html>
