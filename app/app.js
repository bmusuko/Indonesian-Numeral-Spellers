var qs = require('querystring');
var http = require('http');
var fs = require('fs');
var url = require('url');



var server = http.createServer(function(req,res){
	var axios = require('axios');
	var q = url.parse(req.url,true)
	console.log("requests was made "+ req.url)
	if(q.pathname =="/spell/" && req.method =="GET"){ // menangani GET Request
		q = q.query.spell
	    var s = "http://localhost:8081/spell?number="+q
	    var axios = require('axios');
	    var jawab
		axios.get(s)
	  		.then(function (response) {
	    		jawab = response.data.text
	    		console.log(jawab)
	    		res.writeHead(200, {'Content-Type': 'text/html'});
	    		res.write("<html>")
	    		res.write("<style>")
	    		res.write("body{ background-color : pink;}")
	    		res.write("h2 { color: maroon; text-align: center;} ")
	    		res.write("h1 { color: maroon; text-align: center;} ")
	    		res.write("</style>")
	    		res.write("<body>")
	    		res.write("<h2>Spelling Result :</h4>")
	    		res.write("<h1>")
	    		res.write(jawab)
	    		res.write("</h1>")
	    		res.write("<a href='/'>Kembali</a>");
	    		res.write("</body>")
	    		res.write("</html>")
	  		})
	  		.catch(function (error) {
	    		console.log(error);
	  		});
	}else if (q.pathname =="/read/" && req.method =="POST"){ // menangani POST Request
		var requestBody = '';
        req.on('data', function(data) {
            requestBody += data;
        });
		req.on('end', function() {
			var formData = qs.parse(requestBody);
			console.log(formData.read)
			axios.post("http://localhost:8081/read",{
				"text" : formData.read
			}).then(function(response){
				let jawab = response.data.number
	    		res.writeHead(200, {'Content-Type': 'text/html'});
	    		res.write("<html>")
	    		res.write("<style>")
	    		res.write("body{ background-color : pink;}")
	    		res.write("h2 { color: maroon; text-align: center;} ")
	    		res.write("h1 { color: maroon; text-align: center;} ")
	    		res.write("</style>")
	    		res.write("<body>")
	    		res.write("<h2>Reading Result :</h4>")
	    		res.write("<h1>")
	    		res.write(jawab)
	    		res.write("</h1>")
	    		res.write("<a href='/'>Kembali</a>");
	    		res.write("</body>")
	    		res.write("</html>")
			})
			.catch(function (error) {
	    		console.log(error);
	  		});

		});
	} else{
		res.writeHead(200,{'Content-Type' : 'text/html'});
		var myReadStream = fs.createReadStream(__dirname+'/index.html','utf8');
		myReadStream.pipe(res);
	}

});


server.listen(3008,"127.0.0.1");
console.log("server is running in 3008");
