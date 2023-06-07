import axios from 'axios'

const instance = axios.create({})

instance.interceptors.request.use((request) => {
  request.headers.authorization = 'Bearer ' + localStorage.getItem('token')
  return request
})

instance.interceptors.response.use((response) => {
  console.log(response.headers?.["x-authorization-token"])
  if (response.headers?.["x-authorization-token"]) {
    localStorage['token'] = response.headers['x-authorization-token']
  }
  return response
})

export default instance

axios.defaults.baseURL = 'http://localhost:8888'
