(function() {
	storage.count = storage.count || 0;
	var ip = externalIp();
	var bytes = ip.split('.');
	var reverted = "";
	for (var i = 3; i >= 0; i--) {
		reverted += bytes[i];
	}
	
	console.log(JSON.stringify(request));
	
	var response = { ip: reverted.substring(0, config.length), count: storage.count++};
	
	return JSON.stringify(response);
})();