package templ

const NoteTemp = `<!doctype html>
<html>
<head>
    <title>notepad</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">
    <style type="text/css">
            html, body {
                margin: 0px;
                width: 100%;
                height: 100%;
            }
            body {
                background-color: #444444;
                caret-color: #E91E63;
                overflow:hidden;
            }
            #save_fab {
                z-index: 99;
                display: block;
                position: fixed;
                width: 56px;
                height: 56px;
                bottom: 24px;
                right: 24px;
                background-color: #E91E63;
                border-radius: 50%;
                box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 1px 5px 0 rgba(0, 0, 0, 0.12), 0 3px 1px -2px rgba(0, 0, 0, 0.2);
            }
            #save_fab:hover {
                cursor: pointer;
            }
            #save_fab:focus {
                outline:none !important;
            }
            #save_fab:active {
                background-color: #b40c45;
                outline:none !important;
            }
            #save_icon {
	        border-style:none;
                display: inline-block;
                width: 56px;
                height: 56px;
                background: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAAxklEQVR4Ae3TAQYCQRTG8UFUABvRBSIBqHt0gugK0REmC1vtXQJoQGOsvdP0YbBWnjGzvSzv8QN89m8Z1T85Oe99ARVYaDMY2KYE3DM/nBeBgQvjJ5QJ8iI6Q60Sjv4T/AF0BGdAYP4bAOMIyCYBEsAQ0ICGQ3CFhivgDbsvuz1YjoAjsT1xBCyJ7YojYEJspxwBa2K74QjQxLbkeoZnmHU2c7gM9QxdZMQLHlCDidzYmIAa2h+pYgIWIcIN+GEHNyhU7+TkPurjyU/FgJkDAAAAAElFTkSuQmCC) center no-repeat;
            }
	    .container {
		    position: absolute;
		    top: 20px;
		    right: 20px;
		    bottom: 20px;
		    left: 20px;
	    }
	    #content {
		    font-size: 16px;
		    margin: 0;
		    padding: 20px;
		    overflow-y: auto;
		    color: #fff;
		    resize: none;
		    width: 100%;
		    height: 100%;
		    background-color: #303030;
		    min-height: 100%;
		    -webkit-box-sizing: border-box;
		    -moz-box-sizing: border-box;
		    box-sizing: border-box;
		    border: 1px #ddd solid;
		    outline: none;
	    }
    </style>
</head>
<body>
<div>
    <form action="/save" method="post" id="save" target="n_frame">
	    <div id="save_fab">
		    <input type="submit" id="save_icon" value=""/>
	    </div>
    </form>
    <div class="container" form="save">
	    <textarea id="content" oninput="formSubmit()" name="note" form="save" autocomplete="off" spellcheck="false" autocapitalize="off">{{.}}</textarea>
    </div>
    <iframe id="id_iframe" name="n_frame" style="display:none;"></iframe>
</div>
   <script type="text/javascript">
   	function formSubmit(){
		document.getElementById("save").submit()
   	}
   </script>
</body>
</html>`
