<html>
<head>
    <title>上传文件</title>
    <script src="https://cdn.jsdelivr.net/npm/dropzone@5.7.0/dist/dropzone.js"></script>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/dropzone@5.7.0/dist/dropzone.css">
    <style>
    body {
	background: #47cf73;
	height: 100%;
	color: #333;
       line-height: 1.4rem;
       font-family: Roboto, "Open Sans", sans-serif;
       font-size: 20px;
       font-weight: 300;
       text-rendering: optimizeLegibility;
    }

h1 { 
	text-align: center; 
}

.dropzone {
background: #f2f2f2;
	    border-radius: 5px;
border: 2px dashed rgb(0, 135, 247);
	border-image: none;
	max-width: 500px;
	margin-left: auto;
	margin-right: auto;
}
</style>
</head>
<body>

<form class="dropzone needsclick" id="demo-upload" enctype="multipart/form-data" action="/upload" method="post">
      <div class="dz-message needsclick">    
        拖拽文件或者点击上传
      </div>
</form>

<script>
var dropzone = new Dropzone('#demo-upload', {
previewTemplate: document.querySelector('preview-template').innerHTML,
parallelUploads: 2,
thumbnailHeight: 120,
thumbnailWidth: 120,
maxFilesize: 1,
filesizeBase: 1000,
thumbnail: function(file, dataUrl) {
if (file.previewElement) {
file.previewElement.classList.remove("dz-file-preview");
var images = file.previewElement.querySelectorAll("[data-dz-thumbnail]");
for (var i = 0; i < images.length; i++) {
var thumbnailElement = images[i];
thumbnailElement.alt = file.name;
thumbnailElement.src = dataUrl;
}
setTimeout(function() { file.previewElement.classList.add("dz-image-preview"); }, 1);
}
}

});


// Now fake the file upload, since GitHub does not handle file uploads
// and returns a 404
var minSteps = 6,
    maxSteps = 60,
    timeBetweenSteps = 100,
    bytesPerStep = 100000;

dropzone.uploadFiles = function(files) {
	var self = this;

	for (var i = 0; i < files.length; i++) {

		var file = files[i];
		totalSteps = Math.round(Math.min(maxSteps, Math.max(minSteps, file.size / bytesPerStep)));

		for (var step = 0; step < totalSteps; step++) {
			var duration = timeBetweenSteps * (step + 1);
			setTimeout(function(file, totalSteps, step) {
					return function() {
					file.upload = {
progress: 100 * (step + 1) / totalSteps,
total: file.size,
bytesSent: (step + 1) * file.size / totalSteps
};

self.emit('uploadprogress', file, file.upload.progress, file.upload.bytesSent);
if (file.upload.progress == 100) {
file.status = Dropzone.SUCCESS;
self.emit("success", file, 'success', null);
self.emit("complete", file);
self.processQueue();
//document.getElementsByClassName("dz-success-mark").style.opacity = "1";
}
};
}(file, totalSteps, step), duration);
}
}
}
</script>
</body>
</html>
