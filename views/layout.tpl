<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="/static/bootstrap.css">
<title>{{.Title}}</title>
</head>
<body>
<div class="container">
<div class="navbar-header">
Welcome, {{.User.Name}}&nbsp;&nbsp;<a href="/logout">Log Out</a>
</div>
</div>
<hr />
<div class="sidebar">
Hello!
</div>
<div class="container">
<div class="main">
{{template "Body" $}}
</div>
</div>

</body>
</html>