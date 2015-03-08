<html>
<head>
<meta charset="utf-8">

<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="/static/bootstrap.min.css" />
<link rel="stylesheet" href="/static/themes/prism-funky.css" />
<script src="/static/prism.js"></script>
<script src="/static/bootstrap.min.js"></script>


<title>{{.Title}}</title>
</head>
<body>
<div class="container">
<div class="navbar-header">
Welcome, {{.User.Name}}&nbsp;&nbsp;<a href="/logout">Log Out</a>
</div>
</div>
<hr />
<div class="container-fluid">

    <div class="main row">
        <div class="content col-xs-8">
{{template "Body" $}}
        </div>
        <div class="sidebar col-xs-4">
			<h3 align="center">Public Pastes</h3>
			<table align="center">
{{range .Publicpastes}}
				<tr>
					<td align="center">
						<a href="/paste/{{.Id}}">{{.Title}}</a>
					</td>
				</tr>
{{end}}
			</table>
			<h3 align="center">My Pastes</h3>
			<table align="center">
{{range .Mypastes}}
				<tr>
					<td align="center">
						<a href="/paste/{{.Id}}">{{.Title}}</a>
					</td>
				</tr>
{{end}}
			</table>					
							
        </div>
    </div>
</div>





</body>
</html>