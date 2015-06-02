(function() {
	storage.count = storage.count || 0;
	try
	{
		var ip = request.RemoteAddr;
	
		if (ip.indexOf(":") > -1) {
			ip = ip.split(':')[0];
		}
		
		if (ip === "127.0.0.1") {
			ip = externalIp()
		}
		
		var bytes = ip.split('.'),
			reverted = "";
		
		for (var i = 3; i >= 0; i--) {
			reverted += bytes[i];
		}
		
		console.log(JSON.stringify(request));
		
		var response = { ip: reverted.substring(0, config.length), count: storage.count++, randomError: getRandomErrorCode()};
		
		return JSON.stringify(response);
	} catch (e) {
		httpStatusCode = 500
	}
	
})();