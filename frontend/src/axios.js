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
        if (config.push !== '/')  {
            let token = localStorage.getItem('token')
            if (token !== null && token !== '') {
                config.headers.Authorization = "Bearer " + token;
            }
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    });

axios.interceptors.response.use(response => {
        return response
    }, async error => {
        console.log(error)
        let { response } = error
        let error_msg = ''
        if (response) {
            switch (response.status) {
                case 401:
                    error_msg = 'Unauthorized'
                    localStorage.removeItem('username')
                    localStorage.removeItem('token')
                    router.push('/');
                    break
                default:
                    break
            }
            if(error.includes("timeout")) {
                error_msg = 'Request Timeout'
            }
        } else {
            //服务器连结果都没有返回
            error_msg = "Server Not Responding"
            this.$message.bottom().error()
        }
        if (error_msg !== '') {
            this.$message.bottom().error()
        }
        return Promise.reject(error_msg)

    })

// axios.defaults.headers.common['Authorization'] = "Bearer " + localStorage.getItem('token');


export default axios;