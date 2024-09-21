import config from "@/config"
import {
	RefreshToken
} from "@/api/system/usercenter"

import {
	getToken,
	hasToken,
	removeToken,
	setToken
} from "@/library/auth"
import {
	useAppStore
} from "@/library/store/app"

import {
	nav
} from "@/library/nav"



const request = (vm) => {
	uni.$uv.http.interceptors.request.use(
		(config) => {
			// 可使用async await 做异步操作
			// 初始化请求拦截器时，会执行此方法，此时data为undefined，赋予默认{}
			config.data = config.data || {}
			config.header = config.header || {}
			config.header.Authorization = "Bearer " + getToken()
			// 可以在此通过vm引用vuex中的变量，具体值在vm.$store.state中
			// console.log(vm.$store.state);
			return config
		},
		// 可使用async await 做异步操作
		(config) => Promise.reject(config))
}


// 是否正在刷新的标记
let isRefresh = false;
// 重试队列，每一项将是一个待执行的函数形式
let requests = [];

const response = (vm) => {
	uni.$uv.http.interceptors.response.use((response) => {
		// 请求成功直接返回
		const data = response.data
		if (response.statusCode === 200) {
			return data.data || {}
		}

		if (response.statusCode === 401 && data.reason !== 'UnBind' && getToken() == "") {
			nav.login();
			return new Promise(() => {})
		}
		if (response.statusCode === 401) {
			// 未绑定处理
			if (data.reason === 'UnBind') {
				nav.bind();
				return new Promise(() => {})
			}

			// token刷新失败处理
			if (data.reason === 'RefreshToken') {
				return Promise.reject(data)
			}

			// token认证失败处理
			if (data.reason === 'UNAUTHORIZED') {
				console.log(isRefresh)
				if (!isRefresh) {
					isRefresh = true;
					return RefreshToken()
						.then((res) => {
							setToken(res.token);
							// 重新请求
							requests.forEach(cb => cb());
							requests = [];
							return uni.$uv.http.request(response.config);
						})
						.catch((e) => {
							removeToken()
							nav.login();
							return new Promise(() => {})
						})
						.finally(() => {
							console.log("down", isRefresh)
							isRefresh = false;
						});
				}

				return new Promise((resolve) => {
					requests.push(() => {
						resolve(uni.$uv.http.request(response.config))
					});
				});
			}

			nav.login();
			return new Promise(() => {})
		}

		console.log(data)
		// 自定义参数
		const custom = response.config?.custom
		if (data.code !== 200) { // 服务端返回的状态码不等于200，则reject()
			// 如果没有显式定义custom的toast参数为false的话，默认对报错进行toast弹出提示
			if (custom.toast !== false) {
				uni.$uv.toast(data.message)
			}
			// 如果需要catch返回，则进行reject
			if (custom?.catch) {
				return Promise.reject(data)
			} else {
				// 否则返回一个pending中的promise
				return new Promise(() => {})
			}
		}
		return data.data || {}
	}, (response) => {
		uni.$uv.toast("服务器开小差了～")
		return Promise.reject(response)
	})
}

export default () => {
	uni.$uv.http.setConfig((defaultConfig) => {
		/* defaultConfig 为默认全局配置 */
		defaultConfig.baseURL = config.baseUrl /* 根域名 */
		defaultConfig.custom = {
			toast: true,
			catch: true
		}
		defaultConfig.validateStatus = (statusCode) => { // statusCode 必存在。此处示例为全局默认配置
			return statusCode >= 200 && statusCode <= 600
		}
		return defaultConfig
	})
	request();
	response();
}