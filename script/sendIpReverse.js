(function() {
	var ip = externalIp();
	var bytes = ip.split('.');
	var reverted = "";
	for (var i = 3; i >= 0; i--) {
		reverted += bytes[i];
	}
	return reverted.substring(0, config.length);
})();