import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router.js'
import {Axios} from "axios";

import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@mdi/font/css/materialdesignicons.css'
import 'leaflet/dist/leaflet.css'
// 引入Leaflet对象 挂载到Vue上，便于全局使用，也可以单独页面中单独引用
import L from 'leaflet'
import 'leaflet.pm'
import 'leaflet.pm/dist/leaflet.pm.css'


Vue.config.productionTip = false

// ApexCharts
import VueApexCharts from "vue-apexcharts";

Vue.use(VueApexCharts);
Vue.component("apexchart", VueApexCharts)

// LMap
import {LMap, LTileLayer, LMarker} from 'vue2-leaflet';
import 'leaflet/dist/leaflet.css';

Vue.component('l-map', LMap);
Vue.component('l-tile-layer', LTileLayer);
Vue.component('l-marker', LMarker);
Vue.L = Vue.prototype.$L = L
/* leaflet icon */
delete L.Icon.Default.prototype._getIconUrl
L.Icon.Default.mergeOptions({
    iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
    iconUrl: require('leaflet/dist/images/marker-icon.png'),
    shadowUrl: require('leaflet/dist/images/marker-shadow.png'),
})

// axios
import axios from "axios";

Vue.prototype.$axios = axios;

// Vuex
import store from './store'


new Vue({
    vuetify,
    router,
    store,
    render: h => h(App)
}).$mount('#app')

// Axios.interceptors.request.use(config => {
//         // eslint-disable-next-line no-empty
//         if (config.push !== '/')  {
//             if (localStorage.getItem('token')) {
//                 config.headers.token = localStorage.getItem('token');
//             }
//         }
//         return config;
//     },
//     error => {
//         return Promise.reject(error);
//     });
//
// Axios.interceptors.response.use(response => {
//         console.log('response：'+response.data.code)
//         if (response.data.code !== 200) {
//             this.snackbarText = "Token Expired"
//             this.snackbar = true
//             localStorage.removeItem('token');
//             router.push({name: '/'});
//         } else {
//             return response
//         }
//     },
//     error => {
//         return Promise.reject(error);
//     })


