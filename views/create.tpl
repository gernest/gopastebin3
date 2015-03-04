{{define "Body"}}
<p>{{.Paste.Title}}</p>
<pre><code class="language-{{.Prism}}">{{.Paste.Content}}</code></pre>
<p>{{.Paste.Language}}</p>
<a href="/test">New Paste</a>
{{end}}