import cryptoJs from 'crypto-js';
import base64 from "./sdk/base64.js"
import rsa from './sdk/rsa.js';

export default {
	sha256: (data) => {
		const wordArray = cryptoJs.lib.WordArray.create(data);
		return cryptoJs.SHA256(wordArray).toString();
	},
	md5: (data) => {
		const wordArray = cryptoJs.lib.WordArray.create(data);
		return cryptoJs.MD5(wordArray).toString();
	},
	base64: () => {
		return base64
	},
	rsa: () => {
		return rsa
	}
}