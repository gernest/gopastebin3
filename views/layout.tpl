<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<link rel="stylesheet" href="/static/bootstrap.css" />
<link rel="stylesheet" href="/static/themes/prism-funky.css" />
<script src="/static/prism.js"></script>



<title>{{.Title}}</title>
</head>
<body>
<div class="container">
<div class="navbar-header">
Welcome, {{.User.Name}}&nbsp;&nbsp;<a href="/logout">Log Out</a>
</div>
</div>
<hr />
<table align="center">
	<tr>
		<td>
			<h3>Public Pastes</h3>
			<table align="center">
{{range .Publicpastes}}
				<tr>
					<td align="center">
					<a href="/paste/{{.Id}}">{{.Title}}</a>
					</td>
				</tr>
{{end}}
			</table>


			<h3>User Pastes</h3>
			<table align="center">
{{range .Mypastes}}
				<tr>
					<td align="center">
					<a href="/paste/{{.Id}}">{{.Title}}</a>
					</td>
				</tr>
{{end}}
			</table>
		</td>
		<td>
{{template "Body" $}}
		</td>
	</tr>
</table>

</body>
</html>