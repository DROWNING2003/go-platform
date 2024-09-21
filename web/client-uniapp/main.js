import App from './App'
import uvUI from '@/uni_modules/uv-ui-tools'
import mounteGlobal from '@/library/global'
import setting from '@/library/setting';
import "@/library/auth/interceptor"
import interceptor from '@/library/request/interceptor'
import * as Pinia from 'pinia';

import "@/library/auth";

import {
	createSSRApp
} from 'vue'

export function createApp() {
	const app = createSSRApp(App)
	app.use(Pinia.createPinia());

	// 挂载全局对象
	mounteGlobal(app)

	// 引入请求封装
	interceptor(app)

	app.use(uvUI, {
		mpShare: true
	})

	setting();

	return {
		app,
		Pinia,
	}
}