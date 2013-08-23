<!DOCTYPE html>

<html>
  	<head>
    	<title>Beego</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  	</head>
	
	<style type="text/css">
		body {
			margin: 0px;
			font-family: "Helvetica Neue",Helvetica,Arial,sans-serif;
			font-size: 14px;
			line-height: 20px;
			color: rgb(51, 51, 51);
			background-color: rgb(255, 255, 255);
		}

		.hero-unit {
			padding: 60px;
			margin-bottom: 30px;
			border-radius: 6px 6px 6px 6px;
		}

		.container {
			width: 940px;
			margin-right: auto;
			margin-left: auto;
		}

		.row {
			margin-left: -20px;
		}

		h1 {
			margin: 10px 0px;
			font-family: inherit;
			font-weight: bold;
			text-rendering: optimizelegibility;
		}

		.hero-unit h1 {
			margin-bottom: 0px;
			font-size: 60px;
			line-height: 1;
			letter-spacing: -1px;
			color: inherit;
		}

		.description {
		    padding-top: 5px;
		    padding-left: 5px;
		    font-size: 18px;
		    font-weight: 200;
		    line-height: 30px;
		    color: inherit;
		}

		p {
		    margin: 0px 0px 10px;
		}
	</style>
  	
  	<body>
  		<header class="hero-unit" style="background-color:#A9F16C">
			<div class="container">
			<div class="row">
			  <div class="hero-text">
			    <h1>Welcome to Beego!</h1>
			    <p class="description">
			    	Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
			    <br />
			    	Official website: <a href="http://{{.Website}}">{{.Website}}</a>
			    <br />
			    	Contact me: {{.Email}}</a>
			    </p>
			  </div>

			  <div>
			  	{{if .TrueCond}}
			  	true condition.
			  	{{end}}
			  </div>

			  <div>
			  	{{if .FalseCond}}
			  	{{else}}
			  	false condition.
			  	{{end}}
			  </div>

			  <div>
			  	with output:
			  	{{with .User}}
				Name:{{.Name}}; Age:{{.Age}}; Sex:{{.Sex}}
			  	{{end}}
			  </div>

			  <div>
			  	range ouput:
			  	{{range .Nums}}
			  	{{.}}
			  	{{end}};

			  	{{range .Users}}
				Name:{{.Name}}; Age:{{.Age}}; Sex:{{.Sex}}
			  	{{end}}
			  </div>

			  <div>
			  	{{$tplVar := .TplVar}}
			  	The template variable is: {{$tplVar}}.
			  </div>

			  <div>
			  	The result of tempalte function is: {{htmlquote .Html}}
			  </div>

			  <div>
			  	pipeline: {{.Pipe | htmlquote}}
			  </div>

			  <div>
			  	{{template "nested"}}
			  </div>
			</div>
			</div>
		</header>
	</body>
</html>


{{define "nested"}}
Nested template test
{{end}}