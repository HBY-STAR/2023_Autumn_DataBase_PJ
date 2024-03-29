import axios from 'axios'

const request = axios.create({
    baseURL:'/api',
    timeout: 10000,
})

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    // 在发送请求之前做些什么
    // 例如添加请求头、修改请求参数等
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

// 响应拦截器
request.interceptors.response.use(
  (response) => {
    // 对响应数据做点什么
    // 例如处理特定状态码，统一处理错误等
    return response;
  },
  (error) => {
    // 对响应错误做点什么
    // 例如处理特定状态码，统一处理错误等
    return Promise.reject(error);
  }
);

export default request

