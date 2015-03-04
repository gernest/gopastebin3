{{define "Body"}}
<form action="/new" method="post">
<p><b>Title</b>
<input type="text" name="title" required /></p>
<p><b>Content</b><br />
<textarea name="content" required></textarea></p>
<b>Languages</b><br />
<select name="language" required>
	{{range $k, $v := .Languages}}
	<option value={{$k}}>{{$k}}</option>
	{{end}}
</select>
<p><input type="radio" name="ispublic" value="true" checked>Public</input>
   <input type="radio" name="ispublic" value="false">Private</input></p>
<p><input type="submit" value="Create Pastebin" /></p>
</form>
{{end}}