<html>
<head>
    <title>上传文件</title>
    <script src="https://cdn.jsdelivr.net/npm/dropzone@5.7.0/dist/dropzone.js"></script>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/dropzone@5.7.0/dist/dropzone.css">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<style>
	body {
	    background: #adc865;
	    height: 100%;
	    color: #333;
	    line-height: 1.4rem;
	    font-family: Roboto, "Open Sans", sans-serif;
	    font-size: 20px;
	    font-weight: 300;
	    text-rendering: optimizeLegibility;
	}
	div {
		text-align: center;
	}
	.dropzone {
	background: #89d5c9;
		    border-radius: 5px;
	border: 2px dashed rgb(0, 135, 247);
		border-image: none;
		max-width: 500px;
		margin-left: auto;
		margin-right: auto;
	}
</style>
<body>
<form enctype="multipart/form-data" action="/upload" method="post" class="dropzone" id="upload" >
<div class="fallback">
  <input type="file" name="uploadfile" multiple/>
</div>
</form>
<div>
<a href="/file/upload">查看文件</a>
</div>
<script>
  var upload = new Dropzone("#upload", { 
    url: "/upload",
    paramName: "uploadfile",
    maxFilesize: 100, 
    timeout: 190000,
    dictDefaultMessage: "拖拽文件或点击长传"
  });
</script>
</body>
</html>
