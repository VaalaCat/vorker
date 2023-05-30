import axios from 'axios'

const instance = axios.create({})

instance.interceptors.request.use((request) => {
  request.headers.authorization = 'Bearer ' + localStorage.getItem('token')
  return request
})

axios.interceptors.response.use((response) => {
  if (response.headers?.['X-Authorization-Token'] && !localStorage['token']) {
    localStorage['token'] = response.headers['X-Authorization-Token']
  }
  return response
})

export default instance

axios.defaults.baseURL = 'http://localhost:8888'
