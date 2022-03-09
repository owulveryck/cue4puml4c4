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
			ul {
				list-style-type: none;
				margin: 0;
				padding: 0;
			}	
			/* Create two unequal columns that floats next to each other */
			.column {
				float: left;
				padding: 10px;
			}
	
			.left {
				width: 7vw;
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
	
			.container {
				width: 93vw;
				height: 100%;
				left: 7vw;
				position: fixed;
				float: none;
			}
		</style>
		<title>{{.Name}}</title>
	</head>
	
	<body>
		<div class="row">
			<div id="colLeft" class="column left" style="background-color:#aaa;">
				<button class="btn"><a id="dl" download="{{.Name}}.svg" href=""><i class="fa fa-download"></i> Download SVG</a></button>
				<button class="btn" onclick="togglePUML()">Display PUML Content</button>
				<ul>
					<li><a href="../">..</a></li>
					{{range .DirEntries}}
					<li><a href="{{.Path}}">{{.Name}}</a></li>
					{{end}}
				</ul>
				Zoom: <span id="zoomValue">1</span>
			</div>
			<div id="colRight" class="column right" style="background-color:#bbb;">
				<div class="container" id="svgContainer">
				 <svg/>
				</div>
					<pre id="cue">Generating picture...</pre>
				<pre id="plantuml" style="visibility: hidden;"></pre>
			</div>
		</div>
		</div>
	</body>
	
	<script type="text/javascript">
	
		const svgContainer = document.getElementById("svgContainer");
		const right = document.getElementById("svgContainer").getBoundingClientRect();
		const left = document.getElementById("colLeft").getBoundingClientRect();
		svgImage = document.getElementsByTagName("svg")[0];
		svgImage.style = ""
		svgImage.removeAttribute("width");
		svgImage.removeAttribute("height");
		console.log(window.innerHeight)
	
	
		// Full height, including the scroll part
		const fullHeight = Math.max(
			document.body.scrollHeight,
			document.documentElement.scrollHeight,
			document.body.offsetHeight,
			document.documentElement.offsetHeight,
			document.body.clientHeight,
			document.documentElement.clientHeight
		);
		const fullWidth = Math.max(
			document.body.scrollWidth,
			document.documentElement.scrollWidth,
			document.body.offsetWidth,
			document.documentElement.offsetWidth,
			document.body.clientWidth,
			document.documentElement.clientWidth
		);
		//svgImage.setAttribute('viewBox', ` + "`" + `0 0 ${right.width} ${fullWidth}` + "`" + `);
		svgImage.setAttribute('viewBox', ` + "`" + `0 0 1200 ${fullHeight}` + "`" + `);
		console.log(fullHeight)
	
		var viewBox = { x: 0, y: 0, w: right.width, h: fullHeight };
		//svgImage.setAttribute('viewBox', ` + "`" + `${viewBox.x} ${viewBox.y} ${viewBox.w} ${viewBox.h}` + "`" + `);
		svgSize = { w: svgImage.clientWidth, h: svgImage.clientHeight };
		var isPanning = false;
		var startPoint = { x: 0, y: 0 };
		var endPoint = { x: 0, y: 0 };;
		var scale = 1;
	
	
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
			document.getElementById("svgContainer").innerHTML = decodeURIComponent(escape(window.atob(res["image"])))// setAttribute("data", "data:image/svg+xml;utf8;base64," + res["image"]); // decodeURIComponent(escape(window.atob(res["image"]))))
			document.getElementById("dl").setAttribute("href","data:image/svg+xml;utf8;base64," + res["image"]); // decodeURIComponent(escape(window.atob(res["image"]))))
			document.getElementById("cue").innerHTML = decodeURIComponent(escape(window.atob(res["cue"])))
			document.getElementById("plantuml").innerHTML = decodeURIComponent(escape(window.atob(res["plantuml"])))
			svgImage = document.getElementsByTagName("svg")[0];
			svgImage.style = ""
			svgImage.removeAttribute("width");
			svgImage.removeAttribute("height");
			svgImage.setAttribute("preserveAspectRatio","xMidYMid meet")
			console.log(right.width)
			console.log(document.getElementById("colRight").getBoundingClientRect().width)
			console.log(svgSize)
			//svgImage.setAttribute('viewBox', ` + "`" + `0 0 ${fullWidth} ${fullHeight}` + "`" + `);
			svgSize = { w: svgImage.clientWidth, h: svgImage.clientHeight };

		}
	
		ws.onclose = function (e) {
			console.log("[onclose] connection closed (" + e.code + ")");
		}
	
		ws.onerror = function (e) {
			console.log("[onerror] error!");
		}
		function togglePUML() {
			if (document.getElementById("plantuml").style.visibility == "hidden") {
				document.getElementById("plantuml").style.visibility = "visible";
			} else {
				document.getElementById("plantuml").style.visibility = "hidden";
			}
		}
	
		svgContainer.onmousewheel = function (e) {
			e.preventDefault();
			var w = viewBox.w;
			var h = viewBox.h;
			var mx = e.offsetX;//mouse x  
			var my = e.offsetY;
			var dw = w * Math.sign(e.deltaY) * 0.05;
			var dh = h * Math.sign(e.deltaY) * 0.05;
			var dx = dw * mx / svgSize.w;
			var dy = dh * my / svgSize.h;
			viewBox = { x: viewBox.x + dx, y: viewBox.y + dy, w: viewBox.w - dw, h: viewBox.h - dh };
			scale = svgSize.w / viewBox.w;
			zoomValue.innerText = ` + "`" + `${Math.round(scale * 100) / 100}` + "`" + `;
			svgImage.setAttribute('viewBox', ` + "`" + `${viewBox.x} ${viewBox.y} ${viewBox.w} ${viewBox.h}` + "`" + `);
		}
	
	
		svgContainer.onmousedown = function (e) {
			isPanning = true;
			startPoint = { x: e.x, y: e.y };
		}
	
		svgContainer.onmousemove = function (e) {
			if (isPanning) {
				endPoint = { x: e.x, y: e.y };
				var dx = (startPoint.x - endPoint.x) / scale;
				var dy = (startPoint.y - endPoint.y) / scale;
				var movedViewBox = { x: viewBox.x + dx, y: viewBox.y + dy, w: viewBox.w, h: viewBox.h };
				svgImage.setAttribute('viewBox', ` + "`" + `${movedViewBox.x} ${movedViewBox.y} ${movedViewBox.w} ${movedViewBox.h}` + "`" + `);
			}
		}
	
		svgContainer.onmouseup = function (e) {
			if (isPanning) {
				endPoint = { x: e.x, y: e.y };
				var dx = (startPoint.x - endPoint.x) / scale;
				var dy = (startPoint.y - endPoint.y) / scale;
				viewBox = { x: viewBox.x + dx, y: viewBox.y + dy, w: viewBox.w, h: viewBox.h };
				svgImage.setAttribute('viewBox', ` + "`" + `${viewBox.x} ${viewBox.y} ${viewBox.w} ${viewBox.h}` + "`" + `);
				isPanning = false;
			}
		}
	
		svgContainer.onmouseleave = function (e) {
			isPanning = false;
		}
		function togglePUML() {
			if (document.getElementById("plantuml").style.visibility == "hidden") {
				document.getElementById("plantuml").style.visibility = "visible";
				document.getElementById("svgContainer").style.visibility = "hidden";
			} else {
				document.getElementById("svgContainer").style.visibility = "visible";
				document.getElementById("plantuml").style.visibility = "hidden";
			}
		}
	</script>
	
	</html>
	`
)
