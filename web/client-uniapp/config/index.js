const getBaseUrl = () => {
	if (window && window.serverPort) {
		return `${window.location.protocol}//${window.location.hostname}:${window.serverPort}`;
	}

	if (window && window.serverPort) {
		return window.serverHost;
	}
	if (process.env.BASE_URL) {
		return process.env.BASE_URL
	}
	return ""
}

export default {
	server: "PartyAffairs",
	baseUrl: getBaseUrl(),
	rsa: {
		privateKey: "MIICXQIBAAKBgQCi8+NXa27wtcVviKOCYVltuohvqRSwWXJ8gReSqH3C0avjCaoXGznpxCXcywiupLt/EuFYd6gqYQ0jRr/qVz5qZXnE3r4VO7pdLT3GcgCfFvES/1o6WQuCQC0sq1AbZj8vF2goKAmGlSahMDM4uqnTbIsBZ9XtP7x10asuyvnfYwIDAQABAoGAWgS5FlX3F2qTAyZMjX4R1Rgyvtx83V7MbBM0XoRN1Q3XErkICXO07nq5H9vZlADUw0I+0oxnjeHHTDGNW3uOsn5iDgR+A+j5IXujy1LBmDakM9JxaQFJRkDj37hDD+L5RWDolVTcqMu2WD00SPnoLgpHd25k4OmpkdCG2zdf54ECQQDP4bxnhjVmh5N/yuOsV6obtlYbopPfVRPDpo2z+bwpyHhIfF7DtE5n6tHKZBvK3bnIzUUAZgHo2Gf+EhRh2tBRAkEAyKvUQ24B5OOEqKGlmoxlsS21r8dokhPURA49Pk+cbcE9ieTQHQOxV+unNYY1lcK1zrGz/ZE+Ma36dyHjIKnbcwJAAXjHrLP2jNNI6A63rfCw86FG2ss74Mjpauky96M9Ac65i0nWIeyU4Yd4FMbAq58GQbWyV15hhRSDJ/lYy5iRAQJBAMazzqVOQrXgzDUmixKzvuZoeqWp8NPkqqUQ7HfM4PqVBFiMp75zL6H9iD6pvFmkyd33fGCy90C+I2rftUVcrB0CQQCCID/zJIU5X7ZrPZQmIhdqG1XBlouMMbeqEaoVutoZeGBf/m3+nrkf3c0E/jYAb9GaaoRjbYfiFfw4lWKsjJ+N",
		publicKey: "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCi8+NXa27wtcVviKOCYVltuohvqRSwWXJ8gReSqH3C0avjCaoXGznpxCXcywiupLt/EuFYd6gqYQ0jRr/qVz5qZXnE3r4VO7pdLT3GcgCfFvES/1o6WQuCQC0sq1AbZj8vF2goKAmGlSahMDM4uqnTbIsBZ9XtP7x10asuyvnfYwIDAQAB"
	}
}