package main

const (
	index = `<!DOCTYPE html>
	<html lang="en-EN">
	
	<head>
		<link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet"
			type="text/css" />
		<meta charset="utf-8">
		<style>
			* {
				box-sizing: border-box;
			}
	
			html,
			body {
				margin: 0;
				background-color: gray;
			}
	
			/* Create two unequal columns that floats next to each other */
			.column {
				float: left;
				padding: 10px;
			}
	
			.left {
				width: 15vh;
				height: 100vh;
			}
	
			.right {
				height: 100vh;
				float: none;
			}
	
			/* Clear floats after the columns */
			.row:after {
				content: "";
				display: table;
				clear: both;
			}
	
			object {
	
				max-width: 100%;
				max-height: 100%;
				bottom: 0;
				left: 0;
				margin: auto;
				overflow: auto;
				position: fixed;
				right: 0;
				top: 0;
				-o-object-fit: contain;
				object-fit: contain;
	
				aspect-ratio: inherit;
				max-height: 100%;
				max-width: 100%;
	
			}
		</style>
		<title>{{.Name}}</title>
	</head>
	
	<body>
		<div class="row">
			<div class="column left" style="background-color:#aaa;">
				<button class="btn"><i class="fa fa-download"><a id="download" download="{{.Name}}.svg"></i> Download SVG</a></button>
				<ul>
					<li><a href="{{.URL.Path}}/../../">..<a></li>
					{{range .DirEntries}}
					<li><a href="{{.Path}}">{{.Name}}</a></li>
					{{end}}
				</ul>
			</div>
			<div class="column right" style="background-color:#bbb;">
				<pre id="cue">Generating picture...</pre>
				<hr>
				<pre id="plantuml"></pre>
				<object id="output" type="image/svg+xml"></object>
			</div>
		</div>
	</body>
	
	<script type="text/javascript">
		var url = "ws{{if eq .URL.Scheme "https"}}s{{end}}://{{.URL.Host}}{{.URL.Path}}";
		ws = new WebSocket(url);
	
		ws.onopen = function () {
			console.log("[onopen] connect ws uri.");
			var data = {
				"Action": "requireConnect"
			};
			ws.send(JSON.stringify(data));
		}
	
		ws.onmessage = function (e) {
			console.log("[onmessage] receive message.");
			var res = JSON.parse(e.data);
			document.getElementById("output").setAttribute("data", "data:image/svg+xml;utf8;base64," + res["image"]); // decodeURIComponent(escape(window.atob(res["image"]))))
			document.getElementById("download").setAttribute("href","data:image/svg+xml;utf8;base64," + res["image"]) ;
			document.getElementById("cue").innerHTML = decodeURIComponent(escape(window.atob(res["cue"])))
			document.getElementById("plantuml").innerHTML = decodeURIComponent(escape(window.atob(res["plantuml"])))
		}
	
		ws.onclose = function (e) {
			console.log("[onclose] connection closed (" + e.code + ")");
		}
	
		ws.onerror = function (e) {
			console.log("[onerror] error!");
		}
	</script>
	
	</html>
	`
)
