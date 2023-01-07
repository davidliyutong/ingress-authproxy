import axios from 'axios';
import router from './router.js';

// axios 配置
axios.defaults.timeout = 60 * 1000;
axios.defaults.baseURL = '/';

// http request 拦截器
// axios.interceptors.request.use(
//     config => {
//         if (token) {
//             // 判断是否存在token，如果存在的话，则每个http header都加上token
//             config.headers.Authorization = `Bearer ${token}`; // 根据实际情况自行修改
//         }
//         return config;
//     },
//     err => {
//         return Promise.reject(err);
//     }
// );


axios.interceptors.request.use(config => {
        // eslint-disable-next-line no-empty
        if (config.push !== '/')  {
            let token = localStorage.getItem('token')
            if (token) {
                config.headers.Authorization = "Bearer " + token;
            }
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    });

axios.interceptors.response.use(response => {
        console.log('response：'+response.data.code)
        if (response.data.code === 401) {
            localStorage.removeItem('username');
            localStorage.removeItem('token');
            router.push('/');
        } else {
            return response
        }
        return response
    },
    error => {
        return Promise.reject(error);
    })

// axios.defaults.headers.common['Authorization'] = "Bearer " + localStorage.getItem('token');


export default axios;